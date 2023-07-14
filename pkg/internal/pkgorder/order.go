// Package pkgorder to manage cart and order
package pkgorder

import (
	"errors"
	"fmt"
	"time"

	"github.com/fahmifan/commurz/pkg/internal/pkgprice"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	products "github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/samber/lo"
)

var (
	ErrNotFound          = errors.New("not found")
	ErrCartIsFull        = errors.New("cart is full")
	ErrInvalidQuantity   = errors.New("invalid quantity")
	ErrOutOfStock        = errors.New("out of stock")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrTooManyItems      = errors.New("too many items")
)

type Cart struct {
	ID     ulids.ULID
	UserID ulids.ULID

	User  pkguser.User
	Items []CartItem
}

func cartFromSqlc(xcart sqlcs.Cart) Cart {
	return Cart{
		ID:     pkgutil.WeakParseULID(xcart.ID),
		UserID: pkgutil.WeakParseULID(xcart.UserID),
	}
}

func NewCart(userID ulids.ULID) Cart {
	return Cart{
		ID:     ulids.New(),
		UserID: userID,
	}
}

const maxCheckedOutItems = 10

func (cart Cart) CheckoutAll(newOrderNumber OrderNumber, now time.Time) (_ Cart, order Order, checkedoutStocks []pkgproduct.ProductStock, err error) {
	if len(cart.Items) > maxCheckedOutItems {
		return Cart{}, Order{}, nil, fmt.Errorf("[CheckoutAll] too many items: %w", ErrTooManyItems)
	}

	if !cart.isAllItemsHaveStocks() {
		return Cart{}, Order{}, nil, fmt.Errorf("[CheckoutAll] out of stock: %w", ErrOutOfStock)
	}

	cart, items, checkedoutStocks, err := cart.makeOrderItems(now)
	if err != nil {
		return Cart{}, Order{}, nil, fmt.Errorf("[CheckoutAll] make order items: %w", err)
	}

	order = Order{
		ID:     ulids.New(),
		UserID: cart.UserID,
		Number: newOrderNumber,
		Items:  items,
	}

	return cart, order, checkedoutStocks, nil
}

func (cart Cart) isAllItemsHaveStocks() bool {
	return lo.EveryBy(cart.Items, func(item CartItem) bool {
		return item.Product.HaveStock(item.Quantity)
	})
}

const maxCartItems = 99

func (cart Cart) AddItem(product products.Product, qty int64) (Cart, CartItem, error) {

	if len(cart.Items) >= maxCartItems {
		return Cart{}, CartItem{}, ErrCartIsFull
	}

	if qty <= 0 {
		return Cart{}, CartItem{}, ErrInvalidQuantity
	}

	if !product.HaveStock(qty) {
		return Cart{}, CartItem{}, ErrOutOfStock
	}

	cartItem := CartItem{
		ID:           ulids.New(),
		CartID:       cart.ID,
		ProductID:    product.ID,
		Quantity:     qty,
		ProductPrice: product.Price,
		Product:      product,
	}
	cart.Items = append(cart.Items, cartItem)

	return cart, cartItem, nil
}

func (cart Cart) RemoveItem(id ulids.ULID) (Cart, CartItem, error) {
	if len(cart.Items) == 0 {
		return cart, CartItem{}, ErrNotFound
	}

	removedItem, found := lo.Find(cart.Items, func(item CartItem) bool {
		return item.ID == id
	})
	if !found {
		return cart, CartItem{}, ErrNotFound
	}

	cart.Items = lo.Filter(cart.Items, func(item CartItem, _ int) bool {
		return item.ID != id
	})

	return cart, removedItem, nil
}

func (cart Cart) makeOrderItems(now time.Time) (_ Cart, items []OrderItem, checkoutStocks []pkgproduct.ProductStock, err error) {
	orderID := ulids.New()

	checkoutStocks = make([]pkgproduct.ProductStock, len(cart.Items))
	for i := range cart.Items {
		products, reduceStock, err := cart.Items[i].Product.ReduceStock(cart.Items[i].Quantity, now)
		if err != nil {
			return Cart{}, nil, nil, fmt.Errorf("[makeOrderItems] reduce product stock: %w", err)
		}

		checkoutStocks[i] = reduceStock
		cart.Items[i].Product = products
	}

	items = make([]OrderItem, len(cart.Items))
	for i := range cart.Items {
		cartItem := cart.Items[i]

		items[i] = OrderItem{
			ID:       ulids.New(),
			OrderID:  orderID,
			Product:  cartItem.Product,
			Price:    cartItem.ProductPrice,
			Quantity: cartItem.Quantity,
		}
	}

	return cart, items, checkoutStocks, nil
}

type CartItem struct {
	ID        ulids.ULID `json:"id" db:"id"`
	CartID    ulids.ULID `json:"cart_id" db:"cart_id"`
	ProductID ulids.ULID `json:"product_id" db:"product_id"`
	Quantity  int64      `json:"quantity" db:"quantity"`
	// ProductPrice is price per product that will be used when checkout
	ProductPrice pkgprice.Price `json:"product_price" db:"product_price"`

	Product products.Product `json:"product" db:"-"`
}

func cartItemFromSqlc(xcartItem sqlcs.CartItem, idx int) CartItem {
	return CartItem{
		ID:           pkgutil.WeakParseULID(xcartItem.ID),
		CartID:       pkgutil.WeakParseULID(xcartItem.CartID),
		ProductID:    pkgutil.WeakParseULID(xcartItem.ProductID),
		Quantity:     xcartItem.Quantity,
		ProductPrice: pkgprice.New(xcartItem.Price),
	}
}

type OrderNumber string

type Order struct {
	ID     ulids.ULID
	UserID ulids.ULID
	Number OrderNumber

	Items []OrderItem
}

func orderFromSqlc(xorder sqlcs.Order) Order {
	return Order{
		ID:     pkgutil.WeakParseULID(xorder.ID),
		UserID: pkgutil.WeakParseULID(xorder.UserID),
		Number: OrderNumber(xorder.Number),
	}
}

func (order Order) TotalPrice() pkgprice.Price {
	var totalPrice pkgprice.Price
	for _, item := range order.Items {
		totalPrice = item.Price.Times(item.Quantity)
	}

	return totalPrice
}

func (order Order) Products() []products.Product {
	products := lo.Map(order.Items, func(item OrderItem, index int) products.Product {
		return item.Product
	})

	return lo.UniqBy(products, func(product pkgproduct.Product) ulids.ULID {
		return product.ID
	})
}

type OrderItem struct {
	ID        ulids.ULID
	OrderID   ulids.ULID
	ProductID ulids.ULID
	Price     pkgprice.Price
	Quantity  int64

	Product products.Product
}

func orderItemFromSqlc(xorderItem sqlcs.OrderItem, idx int) OrderItem {
	return OrderItem{
		ID:        pkgutil.WeakParseULID(xorderItem.ID),
		OrderID:   pkgutil.WeakParseULID(xorderItem.OrderID),
		Price:     pkgprice.New(xorderItem.Price),
		Quantity:  xorderItem.Quantity,
		ProductID: pkgutil.WeakParseULID(xorderItem.ProductID),
	}
}
