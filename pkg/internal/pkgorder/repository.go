package pkgorder

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/commurz/pkg/preloads"
	"github.com/fahmifan/ulids"
	"github.com/samber/lo"
)

func init() {
	sq.StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

type CartRepository struct{}

func (repo CartRepository) FindCartByUserID(ctx context.Context, tx sqlcs.DBTX, userID ulids.ULID) (Cart, error) {
	queries := sqlcs.New(tx)

	xcart, err := queries.FindCartByUserID(ctx, userID.String())
	if err != nil {
		return Cart{}, fmt.Errorf("[FindCartByUserID] FindCartByUserID: %w", err)
	}

	cart := cartFromSqlc(xcart)

	user, err := pkguser.UserRepository{}.FindUserByID(ctx, tx, userID)
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

func (CartRepository) FindCartItemsByIDs(ctx context.Context, tx sqlcs.DBTX, cartIDs []ulids.ULID) ([]CartItem, error) {
	query := sqlcs.New(tx)

	productRepo := pkgproduct.ProductRepository{}

	xitems, err := query.FindAllCartItemsByCartIDs(ctx, stringULIDs(cartIDs))
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
			return productRepo.FindProductsByIDs(ctx, tx, productIDs)
		},
	}.Preload()
	if err != nil {
		return nil, fmt.Errorf("[FindCartItemsByIDs] preload products: %w", err)
	}

	return items, nil
}

func (CartRepository) SaveCart(ctx context.Context, tx sqlcs.DBTX, cart Cart) (Cart, error) {
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

func (CartRepository) SaveCartItem(ctx context.Context, tx sqlcs.DBTX, cartItem CartItem) (CartItem, error) {
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

func (CartRepository) DeleteCart(ctx context.Context, tx sqlcs.DBTX, cart Cart) error {
	query := sqlcs.New(tx)

	err := query.DeleteCart(ctx, cart.ID.String())
	if err != nil {
		return fmt.Errorf("[DeleteCart] DeleteCart: %w", err)
	}

	return nil
}

type OrderRepository struct{}

func (OrderRepository) CreateOrder(ctx context.Context, tx sqlcs.DBTX, order Order) (Order, error) {
	query := sqlcs.New(tx)

	xorder, err := query.SaveOrder(ctx, sqlcs.SaveOrderParams{
		ID:     order.ID.String(),
		UserID: order.UserID.String(),
		Number: string(order.Number),
	})
	if err != nil {
		return Order{}, fmt.Errorf("[CreateOrder] CreateOrder: %w", err)
	}

	return orderFromSqlc(xorder), nil
}
