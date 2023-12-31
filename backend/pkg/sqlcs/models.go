// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlcs

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID     string
	UserID uuid.UUID
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
	UserID uuid.UUID
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
	ID          string
	Name        string
	Price       int64
	Version     int64
	LatestStock int64
}

type ProductStock struct {
	ID        string
	ProductID string
	StockIn   int64
	StockOut  int64
	CreatedAt time.Time
}

type User struct {
	ID           uuid.UUID
	Email        string
	Name         string
	PasswordHash string
	VerifyToken  string
	Status       string
	LastLoginAt  sql.NullTime
	Archived     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Role         string
}
