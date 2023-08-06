package pkgorder

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/commurz/pkg/preloads"
	"github.com/fahmifan/ulids"
	"github.com/samber/lo"
)

func init() {
	sq.StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

type CartReader struct{}

func (repo CartReader) FindCartByUserID(ctx context.Context, tx sqlcs.DBTX, userID ulids.ULID) (Cart, error) {
	queries := sqlcs.New(tx)

	xcart, err := queries.FindCartByUserID(ctx, userID.String())
	if err != nil {
		return Cart{}, fmt.Errorf("[FindCartByUserID] FindCartByUserID: %w", err)
	}

	cart := cartFromSqlc(xcart)

	user, err := pkguser.UserReader{}.FindUserByID(ctx, tx, userID)
	if err != nil {
		return Cart{}, fmt.Errorf("[FindCartByUserID] FindUserByID: %w", err)
	}

	cartItems, err := repo.FindCartItemsByIDs(ctx, tx, []ulids.ULID{cart.ID})
	if err != nil {
		return Cart{}, fmt.Errorf("[FindCartByUserID] FindCartItemsByIDs: %w", err)
	}

	cart.Items = cartItems
	cart.User = user

	return cart, nil
}

func (CartReader) FindCartItemsByIDs(ctx context.Context, tx sqlcs.DBTX, cartIDs []ulids.ULID) ([]CartItem, error) {
	query := sqlcs.New(tx)

	productRader := pkgproduct.ProductReader{}

	xitems, err := query.FindAllCartItemsByCartIDs(ctx, pkgutil.StringULIDs(cartIDs))
	if err != nil {
		return nil, fmt.Errorf("[FindCartItemsByIDs] FindAllCartItemsByCartIDs: %w", err)
	}

	items := lo.Map(xitems, cartItemFromSqlc)
	productIDs := lo.Map(items, func(item CartItem, index int) ulids.ULID { return item.ProductID })
	items, err = preloads.Preload[CartItem, pkgproduct.Product, ulids.ULID]{
		Targets:   items,
		RefItem:   func(item pkgproduct.Product) ulids.ULID { return item.ID },
		RefTarget: func(target CartItem) ulids.ULID { return target.ProductID },
		SetItem:   func(item *CartItem, target pkgproduct.Product) { item.Product = target },
		FetchItems: func() ([]pkgproduct.Product, error) {
			return productRader.FindProductsByIDs(ctx, tx, productIDs)
		},
	}.Preload()
	if err != nil {
		return nil, fmt.Errorf("[FindCartItemsByIDs] preload products: %w", err)
	}

	return items, nil
}

type CartWriter struct{}

func (CartWriter) SaveCart(ctx context.Context, tx sqlcs.DBTX, cart Cart) (Cart, error) {
	query := sqlcs.New(tx)

	xcart, err := query.CreateCart(ctx, sqlcs.CreateCartParams{
		ID:     cart.ID.String(),
		UserID: cart.UserID.String(),
	})
	if err != nil {
		return Cart{}, fmt.Errorf("[SaveCart] query: %w", err)
	}

	return cartFromSqlc(xcart), nil
}

func (CartWriter) SaveCartItem(ctx context.Context, tx sqlcs.DBTX, cartItem CartItem) (CartItem, error) {
	query := sqlcs.New(tx)

	xcartItem, err := query.SaveCartItem(ctx, sqlcs.SaveCartItemParams{
		CartID:    cartItem.CartID.String(),
		ID:        cartItem.ID.String(),
		Price:     cartItem.ProductPrice.IDR(),
		ProductID: cartItem.ProductID.String(),
		Quantity:  cartItem.Quantity,
	})
	if err != nil {
		return CartItem{}, fmt.Errorf("[SaveCartItem] SaveCartItem: %w", err)
	}

	newCartItem := cartItemFromSqlc(xcartItem, 0)

	// set additional data not returning from save query
	newCartItem.Product = cartItem.Product

	return newCartItem, nil
}

func (CartWriter) DeleteCart(ctx context.Context, tx sqlcs.DBTX, cart Cart) error {
	query := sqlcs.New(tx)

	err := query.DeleteCart(ctx, cart.ID.String())
	if err != nil {
		return fmt.Errorf("[DeleteCart] DeleteCart: %w", err)
	}

	return nil
}

type OrderReader struct{}

func (repo OrderReader) FindOrderItemsByOrderID(ctx context.Context, tx sqlcs.DBTX, orderID ulids.ULID) ([]OrderItem, error) {
	query := sqlcs.New(tx)
	productRepo := pkgproduct.ProductReader{}

	xitems, err := query.FindOrderItemsByOrderID(ctx, orderID.String())
	if err != nil {
		return nil, fmt.Errorf("[FindOrderItemsByOrderID] FindOrderItemsByOrderID: %w", err)
	}

	items := lo.Map(xitems, orderItemFromSqlc)
	productIDs := lo.Map(items, func(item OrderItem, index int) ulids.ULID { return item.ProductID })

	items, err = preloads.Preload[OrderItem, pkgproduct.Product, ulids.ULID]{
		Targets:   items,
		RefTarget: func(target OrderItem) ulids.ULID { return target.ProductID },
		RefItem:   func(item pkgproduct.Product) ulids.ULID { return item.ID },
		SetItem:   func(item *OrderItem, target pkgproduct.Product) { item.Product = target },
		FetchItems: func() ([]pkgproduct.Product, error) {
			return productRepo.FindProductsByIDs(ctx, tx, productIDs)
		},
	}.Preload()
	if err != nil {
		return nil, fmt.Errorf("[FindOrderItemsByOrderID] preload products: %w", err)
	}

	return items, nil
}

func (repo OrderReader) FindOrderByID(ctx context.Context, tx sqlcs.DBTX, id ulids.ULID) (Order, error) {
	query := sqlcs.New(tx)

	xorder, err := query.FindOrderByID(ctx, id.String())
	if err != nil {
		return Order{}, fmt.Errorf("[FindByID] FindOrderByID: %w", err)
	}

	order := orderFromSqlc(xorder)
	order.Items, err = repo.FindOrderItemsByOrderID(ctx, tx, order.ID)
	if err != nil {
		return Order{}, fmt.Errorf("[FindByID] preload products: %w", err)
	}

	return order, nil
}

type OrderWriter struct{}

func (OrderWriter) CreateOrder(ctx context.Context, tx sqlcs.DBTX, order Order) (Order, error) {
	query := sqlcs.New(tx)

	_, err := query.SaveOrder(ctx, sqlcs.SaveOrderParams{
		ID:     order.ID.String(),
		UserID: order.UserID.String(),
		Number: string(order.Number),
	})
	if err != nil {
		return Order{}, fmt.Errorf("[CreateOrder] CreateOrder: %w", err)
	}

	return order, nil
}

func (OrderWriter) CreateOrderItem(ctx context.Context, tx sqlcs.DBTX, orderItem OrderItem) (OrderItem, error) {
	query := sqlcs.New(tx)

	xorderItem, err := query.SaveOrderItem(ctx, sqlcs.SaveOrderItemParams{
		ID:        orderItem.ID.String(),
		OrderID:   orderItem.OrderID.String(),
		Price:     orderItem.Product.Price.Value(),
		ProductID: orderItem.Product.ID.String(),
		Quantity:  orderItem.Quantity,
	})
	if err != nil {
		return OrderItem{}, fmt.Errorf("[SaveOrderItem] SaveOrderItem: %w", err)
	}

	newOrderItem := orderItemFromSqlc(xorderItem, 0)

	// set additional data not returning from save query
	newOrderItem.Product = orderItem.Product

	return newOrderItem, nil
}

func (repo OrderWriter) BulkSaveOrderItems(ctx context.Context, tx sqlcs.DBTX, orderItems []OrderItem) ([]OrderItem, error) {
	for i := range orderItems {
		item, err := repo.CreateOrderItem(ctx, tx, orderItems[i])
		if err != nil {
			return nil, fmt.Errorf("[SaveOrderItems] SaveOrderItem: %w", err)
		}

		orderItems[i] = item
	}

	return orderItems, nil
}