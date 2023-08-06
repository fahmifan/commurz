package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkgorder"
	"github.com/fahmifan/commurz/pkg/internal/pkgprice"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
	"github.com/fahmifan/ulids"
)

func (service *Service) CheckoutAll(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CheckoutAllRequest],
) (*connect.Response[commurzpbv1.Order], error) {
	userID, err := pkgutil.ParseULID(req.Msg.UserId)
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] parse userID: %w", err)
	}

	cartReader := pkgorder.CartReader{}
	cartWriter := pkgorder.CartWriter{}
	orderReader := pkgorder.OrderReader{}
	orderWriter := pkgorder.OrderWriter{}
	productWriter := pkgproduct.ProductWriter{}

	orderNumber := pkgorder.OrderNumber(ulids.New().String())
	now := time.Now()
	order := pkgorder.Order{}
	var checkOutStocks []pkgproduct.ProductStock

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
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

	res := &connect.Response[commurzpbv1.Order]{
		Msg: protoserde.FromOrderPkg(order),
	}

	return res, nil
}

func (service *Service) AddProductToCart(
	ctx context.Context,
	req *connect.Request[commurzpbv1.AddProductToCartRequest],
) (*connect.Response[commurzpbv1.Cart], error) {
	userID, err := pkgutil.ParseULID(req.Msg.GetUserId())
	if err != nil {
		return nil, fmt.Errorf("[AddItemToCart] parse userID: %w", err)
	}

	productID, err := pkgutil.ParseULID(req.Msg.ProductId)
	if err != nil {
		return nil, fmt.Errorf("[AddItemToCart] parse productID: %w", err)
	}

	cartWriter := pkgorder.CartWriter{}
	cart, err := service.getOrCreateCart(ctx, service.DB, userID)
	if err != nil {
		return nil, fmt.Errorf("[AddItemToCart] getOrCreateCart: %w", err)
	}

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		productRepo := pkgproduct.ProductReader{}

		product, err := productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[AddItemToCart] FindProductByID: %w", err)
		}

		var cartItem pkgorder.CartItem

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
		return nil, fmt.Errorf("[AddItemToCart] Transaction: %w", err)
	}

	res := &connect.Response[commurzpbv1.Cart]{
		Msg: protoserde.FromCartPkg(cart),
	}

	return res, nil
}

func (service *Service) getOrCreateCart(ctx context.Context, tx sqlcs.DBTX, userID ulids.ULID) (cart pkgorder.Cart, err error) {
	cartReader := pkgorder.CartReader{}
	cartWriter := pkgorder.CartWriter{}

	cart, err = cartReader.FindCartByUserID(ctx, tx, userID)

	if err != nil && !isNotFoundErr(err) {
		return pkgorder.Cart{}, fmt.Errorf("[getOrCreateCart] FindCartByUserID: %w", err)
	}

	if err == nil {
		return cart, nil
	}

	cart = pkgorder.NewCart(userID)
	_, err = cartWriter.SaveCart(ctx, tx, cart)
	if err != nil && !isNotFoundErr(err) {
		return pkgorder.Cart{}, fmt.Errorf("[getOrCreateCart] SaveCart: %w", err)
	}

	// refresh cart data
	cart, err = cartReader.FindCartByUserID(ctx, tx, userID)
	if err != nil {
		return pkgorder.Cart{}, fmt.Errorf("[getOrCreateCart] FindCartByUserID: %w", err)
	}

	return cart, nil
}

func (service *Service) CreateUser(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CreateUserRequest],
) (res *connect.Response[commurzpbv1.User], err error) {
	userReader := pkguser.UserReader{}
	userWriter := pkguser.UserWriter{}

	user, err := pkguser.NewUser(req.Msg.Email)
	if err != nil {
		return res, fmt.Errorf("[CreateUser] NewUser: %w", err)
	}

	_, err = userWriter.CreateUser(ctx, service.DB, user)
	if err != nil {
		return res, fmt.Errorf("[CreateUser] SaveUser: %w", err)
	}

	user, err = userReader.FindUserByID(ctx, service.DB, user.ID)
	if err != nil {
		return res, fmt.Errorf("[CreateUser] FindUserByID: %w", err)
	}

	res = &connect.Response[commurzpbv1.User]{
		Msg: protoserde.FromUserPkg(user),
	}

	return res, nil
}

func (service *Service) CreateProduct(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CreateProductRequest],
) (res *connect.Response[commurzpbv1.Product], err error) {
	productRepo := pkgproduct.ProductReader{}
	productWriter := pkgproduct.ProductWriter{}

	product, err := pkgproduct.CreateProduct(req.Msg.Name, pkgprice.New(req.Msg.Price))
	if err != nil {
		return res, fmt.Errorf("[CreateProduct] CreateProduct: %w", err)
	}

	_, err = productWriter.SaveProduct(ctx, service.DB, product)
	if err != nil {
		return res, fmt.Errorf("[CreateProduct] CreateProduct: %w", err)
	}

	product, err = productRepo.FindProductByID(ctx, service.DB, product.ID)
	if err != nil {
		return res, fmt.Errorf("[CreateProduct] FindProductByID: %w", err)
	}

	res = &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return res, nil
}

func (service *Service) AddProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ChangeProductStockRequest],
) (res *connect.Response[commurzpbv1.Product], err error) {
	productRepo := pkgproduct.ProductReader{}
	productWriter := pkgproduct.ProductWriter{}
	product := pkgproduct.Product{}

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		productID, err := pkgutil.ParseULID(req.Msg.GetProductId())
		if err != nil {
			return fmt.Errorf("[AddProductStock] ParseULID: %w", err)
		}

		product, err = productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[AddProductStock] FindProductByID: %w", err)
		}

		var stock pkgproduct.ProductStock
		product, stock = product.AddStock(req.Msg.GetStockQuantity(), time.Now())

		_, err = productWriter.SaveProductStock(ctx, tx, stock)
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
		return res, fmt.Errorf("[AddProductStock] Transaction: %w", err)
	}

	res = &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return
}

func (service *Service) ReduceProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ChangeProductStockRequest],
) (*connect.Response[commurzpbv1.Product], error) {
	productReader := pkgproduct.ProductReader{}
	productWriter := pkgproduct.ProductWriter{}

	product := pkgproduct.Product{}
	productID, err := pkgutil.ParseULID(req.Msg.ProductId)
	if err != nil {
		return nil, fmt.Errorf("[ReduceProductStock] ParseULID: %w", err)
	}

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		product, err = productReader.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] FindProductByID: %w", err)
		}

		var stock pkgproduct.ProductStock
		product, stock, err = product.ReduceStock(req.Msg.GetStockQuantity(), time.Now())
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] ReduceStock: %w", err)
		}

		_, err = productWriter.SaveProductStock(ctx, tx, stock)
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] SaveProductStock: %w", err)
		}

		_, err = productWriter.BumpProductVersion(ctx, tx, product)
		if err != nil {
			return fmt.Errorf("[AddProductStock] BumpProductVersion: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[AddProductStock] Transaction: %w", err)
	}

	product, err = productReader.FindProductByID(ctx, service.DB, productID)
	if err != nil {
		return nil, fmt.Errorf("[ReduceProductStock] FindProductByID: %w", err)
	}

	res := &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return res, nil
}
