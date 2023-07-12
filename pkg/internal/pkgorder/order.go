// Package pkgorder to manage cart and order
package pkgorder

import (
	"errors"

	products "github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/oklog/ulid/v2"
	"github.com/samber/lo"
)

var (
	ErrCartIsFull        = errors.New("cart is full")
	ErrInvalidQuantity   = errors.New("invalid quantity")
	ErrOutOfStock        = errors.New("out of stock")
	ErrInsufficientStock = errors.New("insufficient stock")
)

type Price int64

type Cart struct {
	ID     ulids.ULID
	UserID ulids.ULID

	User  pkguser.User
	Items []CartItem
}

func cartFromSqlc(xcart sqlcs.Cart) Cart {
	return Cart{
		ID:     mustParseULID(xcart.ID),
		UserID: mustParseULID(xcart.UserID),
	}
}

func NewCart(userID ulids.ULID) Cart {
	return Cart{
		ID:     ulids.New(),
		UserID: userID,
	}
}

func (cart Cart) CheckoutAll(orderNumber OrderNumber) (Order, error) {
	if !cart.isAllItemsHaveStocks() {
		return Order{}, ErrOutOfStock
	}

	order := Order{
		ID:     ulids.New(),
		Number: orderNumber,
		Items:  cart.getOrderItems(),
	}

	return order, nil
}

func (cart Cart) isAllItemsHaveStocks() bool {
	return lo.EveryBy(cart.Items, func(item CartItem) bool {
		return item.Product.HaveStock(item.Quantity)
	})
}

// CheckoutByProducts will checkout only the given products
func (cart Cart) CheckoutByProducts(products []products.Product, orderNumber OrderNumber) Order {
	// check products stock

	items := lo.Filter(
		cart.getOrderItems(),
		cart.filterOrderItemsByProduct(products),
	)

	return Order{
		ID:     ulids.New(),
		Number: orderNumber,
		Items:  items,
	}
}

func (cart Cart) filterOrderItemsByProduct(allProducts []products.Product) func(item OrderItem, index int) bool {
	mapProduct := make(map[ulids.ULID]products.Product, len(allProducts))
	for _, product := range allProducts {
		mapProduct[product.ID] = product
	}

	return func(item OrderItem, index int) bool {
		_, ok := mapProduct[item.Product.ID]
		return ok
	}
}

func (cart Cart) AddItem(product products.Product, qty int64) (Cart, CartItem, error) {
	const maxCartItems = 99

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

func (c Cart) LastItem() CartItem {
	if len(c.Items) == 0 {
		return CartItem{}
	}

	return c.Items[len(c.Items)-1]
}

func (c Cart) RemoveItem(id ulids.ULID) Cart {
	if len(c.Items) == 0 {
		return c
	}

	items := make([]CartItem, 0, len(c.Items)-1)
	for _, item := range c.Items {
		if item.ID == id {
			continue
		}

		items = append(items, item)
	}

	c.Items = items

	return c
}

func (cart Cart) getOrderItems() []OrderItem {
	orderID := ulids.New()

	items := make([]OrderItem, len(cart.Items))
	for i := range cart.Items {
		cartItem := cart.Items[i]

		items[i] = OrderItem{
			ID:         ulids.New(),
			OrderID:    orderID,
			CartItemID: cartItem.ID,
			Product:    cartItem.Product,
			Price:      Price(cartItem.ProductPrice),
			Quantity:   cartItem.Quantity,
		}
	}

	return items
}

type CartItem struct {
	ID        ulids.ULID `json:"id" db:"id"`
	CartID    ulids.ULID `json:"cart_id" db:"cart_id"`
	ProductID ulids.ULID `json:"product_id" db:"product_id"`
	Quantity  int64      `json:"quantity" db:"quantity"`
	// ProductPrice is price per product that will be used when checkout
	ProductPrice products.Price `json:"product_price" db:"product_price"`

	Product products.Product `json:"product" db:"-"`
}

func cartItemFromSqlc(xcartItem sqlcs.CartItem, idx int) CartItem {
	return CartItem{
		ID:           mustParseULID(xcartItem.ID),
		CartID:       mustParseULID(xcartItem.CartID),
		ProductID:    mustParseULID(xcartItem.ProductID),
		Quantity:     xcartItem.Quantity,
		ProductPrice: products.Price(xcartItem.Price),
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
		ID:     mustParseULID(xorder.ID),
		UserID: mustParseULID(xorder.UserID),
		Number: OrderNumber(xorder.Number),
	}
}

type OrderItem struct {
	ID         ulids.ULID
	OrderID    ulids.ULID
	CartItemID ulids.ULID
	Product    products.Product
	Price      Price
	Quantity   int64
}

func mustParseULID(s string) ulids.ULID {
	return ulids.ULID{ULID: ulid.MustParse(s)}
}

func stringULIDs(ids []ulids.ULID) []string {
	return lo.Map(ids, func(id ulids.ULID, index int) string {
		return id.String()
	})
}
