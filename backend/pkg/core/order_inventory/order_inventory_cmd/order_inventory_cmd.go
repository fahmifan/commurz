package order_inventory_cmd

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/fahmifan/commurz/pkg/core"
	"github.com/fahmifan/commurz/pkg/core/auth"
	"github.com/fahmifan/commurz/pkg/core/order_inventory"
	"github.com/fahmifan/commurz/pkg/logs"
	"github.com/fahmifan/commurz/pkg/parseutil"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/pkgmoney"
	"github.com/fahmifan/commurz/pkg/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/google/uuid"
)

type OrderInventoryCmd struct {
	*core.Ctx
}

func (service *OrderInventoryCmd) UpdateProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.UpdateProductStockRequest],
) (res *connect.Response[commurzpbv1.Empty], err error) {
	if err := service.CanUser(ctx, auth.Manage, auth.Product); err != nil {
		return nil, err
	}

	productRepo := order_inventory.ProductReader{}
	productWriter := order_inventory.ProductWriter{}
	product := order_inventory.Product{}

	err = core.Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		productID, err := parseutil.ParseULID(req.Msg.GetProductId())
		if err != nil {
			return fmt.Errorf("[AddProductStock] ParseULID: %w", err)
		}

		product, err = productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[AddProductStock] FindProductByID: %w", err)
		}

		if !product.SameVersion(req.Msg.GetVersion()) {
			return fmt.Errorf("[AddProductStock] VersionMismatch")
		}

		var stockIn, stockOut order_inventory.ProductStock

		product, stockIn, err = product.AddStock(req.Msg.GetStockIn(), time.Now())
		if err != nil {
			return fmt.Errorf("[AddProductStock] AddStock: %w", err)
		}

		product, stockOut, err = product.ReduceStock(req.Msg.GetStockOut(), time.Now())
		if err != nil {
			return fmt.Errorf("[AddProductStock] ReduceStock: %w", err)
		}

		_, err = productWriter.BulkSaveProductStocks(ctx, tx, []order_inventory.ProductStock{stockIn, stockOut})
		if err != nil {
			return fmt.Errorf("[AddProductStock] SaveProductStock: %w", err)
		}

		_, err = productWriter.BumpProductVersion(ctx, tx, product)
		if err != nil {
			return fmt.Errorf("[AddProductStock] BumpProductVersion: %w", err)
		}

		return nil
	})

	if err != nil {
		logs.ErrCtx(ctx, err, "UpdateProductStock: Transaction")
		return nil, core.ErrInternal
	}

	res = &connect.Response[commurzpbv1.Empty]{}

	return
}

func (service *OrderInventoryCmd) CreateProduct(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CreateProductRequest],
) (res *connect.Response[commurzpbv1.Empty], err error) {
	productRepo := order_inventory.ProductReader{}
	productWriter := order_inventory.ProductWriter{}

	product, err := order_inventory.CreateProduct(req.Msg.Name, pkgmoney.New(req.Msg.Price))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	product, err = productWriter.SaveProduct(ctx, service.DB, product)
	if err != nil {
		logs.ErrCtx(ctx, err, "[CreateProduct] SaveProduct")
		return nil, core.ErrInternal
	}

	product, err = productRepo.FindProductByID(ctx, service.DB, product.ID)
	if err != nil {
		logs.ErrCtx(ctx, err, "[CreateProduct] FindProductByID")
		return nil, core.ErrInternal
	}

	res = &connect.Response[commurzpbv1.Empty]{}

	return res, nil
}

func (service *OrderInventoryCmd) AddProductToCart(
	ctx context.Context,
	req *connect.Request[commurzpbv1.AddProductToCartRequest],
) (*connect.Response[commurzpbv1.Empty], error) {
	userID, err := uuid.Parse(req.Msg.GetUserId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	productID, err := parseutil.ParseULID(req.Msg.ProductId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	cartWriter := order_inventory.CartWriter{}
	cart, err := service.getOrCreateCart(ctx, service.DB, userID)
	if err != nil {
		logs.ErrCtx(ctx, err, "[AddProductToCart] getOrCreateCart")
		return nil, connect.NewError(connect.CodeInternal, core.ErrInternal)
	}

	err = core.Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		productRepo := order_inventory.ProductReader{}

		product, err := productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[AddItemToCart] FindProductByID: %w", err)
		}

		var cartItem order_inventory.CartItem

		cart, cartItem, err = cart.AddItem(product, req.Msg.GetQuantity())
		if err != nil {
			return fmt.Errorf("[AddItemToCart] AddItem: %w", err)
		}

		_, err = cartWriter.SaveCartItem(ctx, tx, cartItem)
		if err != nil {
			return fmt.Errorf("[AddItemToCart] SaveCartItem: %w", err)
		}

		return nil
	})
	if err != nil {
		logs.ErrCtx(ctx, err, "[AddProductToCart] Transaction")
		return nil, core.ErrInternal
	}

	res := &connect.Response[commurzpbv1.Empty]{}

	return res, nil
}

func (service *OrderInventoryCmd) getOrCreateCart(ctx context.Context, tx sqlcs.DBTX, userID uuid.UUID) (cart order_inventory.Cart, err error) {
	cartReader := order_inventory.CartReader{}
	cartWriter := order_inventory.CartWriter{}

	cart, err = cartReader.FindCartByUserID(ctx, tx, userID)

	if err != nil && !core.IsNotFoundErr(err) {
		return order_inventory.Cart{}, fmt.Errorf("[getOrCreateCart] FindCartByUserID: %w", err)
	}

	if err == nil {
		return cart, nil
	}

	cart = order_inventory.NewCart(userID)
	_, err = cartWriter.SaveCart(ctx, tx, cart)
	if err != nil && !core.IsNotFoundErr(err) {
		return order_inventory.Cart{}, fmt.Errorf("[getOrCreateCart] SaveCart: %w", err)
	}

	// refresh cart data
	cart, err = cartReader.FindCartByUserID(ctx, tx, userID)
	if err != nil {
		return order_inventory.Cart{}, fmt.Errorf("[getOrCreateCart] FindCartByUserID: %w", err)
	}

	return cart, nil
}

func (service *OrderInventoryCmd) CheckoutAll(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CheckoutAllRequest],
) (*connect.Response[commurzpbv1.Empty], error) {
	_, ok := auth.UserFromCtx(ctx)
	if !ok {
		return nil, core.ErrUnauthenticated
	}

	userID, err := uuid.Parse(req.Msg.UserId)
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] parse userID: %w", err)
	}

	cartReader := order_inventory.CartReader{}
	cartWriter := order_inventory.CartWriter{}
	orderReader := order_inventory.OrderReader{}
	orderWriter := order_inventory.OrderWriter{}
	productWriter := order_inventory.ProductWriter{}

	orderNumber := order_inventory.OrderNumber(ulids.New().String())
	now := time.Now()
	order := order_inventory.Order{}
	var checkOutStocks []order_inventory.ProductStock

	err = core.Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		cart, err := cartReader.FindCartByUserID(ctx, tx, userID)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] FindCartByUserID: %w", err)
		}

		cart, order, checkOutStocks, err = cart.CheckoutAll(orderNumber, now)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] CheckoutAll: %w", err)
		}

		_, err = orderWriter.CreateOrder(ctx, tx, order)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] CreateOrder: %w", err)
		}

		_, err = orderWriter.BulkSaveOrderItems(ctx, tx, order.Items)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] BulkSaveOrderItems: %w", err)
		}

		// TODO: we can improve this later
		//
		// bump product version while saving the reduced product stocks to avoid race condition.
		// product stocks should have rolledback if the product version is failed to bumped.
		{
			_, err = productWriter.BulkSaveProductStocks(ctx, tx, checkOutStocks)
			if err != nil {
				return fmt.Errorf("[CheckoutAll] BulkSaveProductStocks: %w", err)
			}
			_, err = productWriter.BulkBumpProductVersion(ctx, tx, order.Products())
			if err != nil {
				return fmt.Errorf("[CheckoutAll] BulkBumpProductVersion: %w", err)
			}
		}

		err = cartWriter.DeleteCart(ctx, tx, cart)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] DeleteCart: %w", err)
		}

		// TODO: create and send invoice::pending_payment to user

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] Transaction: %w", err)
	}

	order, err = orderReader.FindOrderByID(ctx, service.DB, order.ID)
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] FindOrderByID: %w", err)
	}

	res := &connect.Response[commurzpbv1.Empty]{}

	return res, nil
}
