// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package sqlcs

import (
	"time"
)

type Cart struct {
	ID     string
	UserID string
}

type CartItem struct {
	ID        string
	CartID    string
	ProductID string
	Quantity  int64
	Price     int64
}

type Order struct {
	ID     string
	UserID string
	Number string
}

type OrderItem struct {
	ID        string
	OrderID   string
	ProductID string
	Quantity  int64
	Price     int64
}

type Product struct {
	ID      string
	Name    string
	Price   int64
	Version int64
}

type ProductStock struct {
	ID        string
	ProductID string
	StockIn   int64
	StockOut  int64
	CreatedAt time.Time
}

type User struct {
	ID    string
	Email string
}