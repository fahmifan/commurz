// Package pkgproduct to manage product and stock
package pkgproduct

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/oklog/ulid/v2"
	"github.com/samber/lo"
)

var (
	ErrInsufficientStock = errors.New("insufficient stock")
)

type Price int64

type Product struct {
	ID    ulids.ULID
	Name  string
	Price Price

	Stocks []ProductStock
}

func productFromSqlc(p sqlcs.Product, index int) Product {
	return Product{
		ID:    ulids.ULID{ULID: ulid.MustParse(p.ID)},
		Name:  p.Name,
		Price: Price(p.Price),
	}
}

func CreateProduct(name string, price Price) Product {
	return Product{
		ID:    ulids.New(),
		Name:  name,
		Price: price,
	}
}

func (product Product) AddStock(stockIn int64, createdAt time.Time) (Product, ProductStock) {
	stock := ProductStock{
		ID:        ulids.New(),
		ProductID: product.ID,
		StockIn:   stockIn,
		CreatedAt: createdAt,
	}

	product.Stocks = append(product.Stocks, stock)
	return product, stock
}

func (product Product) ReduceStock(stockOut int64, createdAt time.Time) (Product, ProductStock, error) {
	if product.currentStock() < stockOut {
		return Product{}, ProductStock{}, ErrInsufficientStock
	}

	stock := ProductStock{
		ID:        ulids.New(),
		ProductID: product.ID,
		StockOut:  stockOut,
		CreatedAt: createdAt,
	}
	product.Stocks = append(product.Stocks, stock)

	return product, stock, nil
}

func (product Product) HaveStock(qty int64) bool {
	return product.currentStock() >= qty
}

func (product Product) currentStock() int64 {
	return product.totalStockIn() - product.totalStockOut()
}

func (product Product) totalStockIn() (total int64) {
	for _, stock := range product.Stocks {
		total += stock.StockIn
	}

	return total
}

func (product Product) totalStockOut() (total int64) {
	for _, stock := range product.Stocks {
		total += stock.StockOut
	}

	return total
}

// ProductStock is a log of product stock in/out
type ProductStock struct {
	ID        ulids.ULID
	ProductID ulids.ULID
	StockIn   int64
	StockOut  int64
	CreatedAt time.Time
}

func productStockFromSqlc(from sqlcs.ProductStock, index int) ProductStock {
	return ProductStock{
		ID:        ulids.ULID{ULID: ulid.MustParse(from.ID)},
		ProductID: ulids.ULID{ULID: ulid.MustParse(from.ProductID)},
		StockIn:   from.StockIn,
		StockOut:  from.StockOut,
		CreatedAt: from.CreatedAt,
	}
}

func mustParseULID(s string) ulids.ULID {
	return ulids.ULID{ULID: ulid.MustParse(s)}
}

func stringULIDs(ids []ulids.ULID) []string {
	return lo.Map(ids, func(id ulids.ULID, index int) string {
		return id.String()
	})
}

func prettyJSON(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
