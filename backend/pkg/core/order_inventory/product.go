// Package pkgproduct to manage product and stock
package order_inventory

import (
	"fmt"
	"time"

	"github.com/fahmifan/commurz/pkg/core/pkgmoney"
	"github.com/fahmifan/commurz/pkg/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/oklog/ulid/v2"
)

type Product struct {
	ID      ulids.ULID
	Name    string
	Price   pkgmoney.Money
	Version int64

	Stocks []ProductStock
}

func productFromSqlc(p sqlcs.Product, index int) Product {
	return Product{
		ID:      ulids.ULID{ULID: ulid.MustParse(p.ID)},
		Name:    p.Name,
		Price:   pkgmoney.New(p.Price),
		Version: p.Version,
	}
}

var minPrice = pkgmoney.New(10)

func CreateProduct(name string, productPrice pkgmoney.Money) (Product, error) {
	minNameLen := 3
	if len(name) < minNameLen {
		return Product{}, fmt.Errorf("min name length is %d characters", minNameLen)
	}

	if productPrice.Value() < minPrice.Value() {
		return Product{}, fmt.Errorf("min price is %s", minPrice.String())
	}

	product := Product{
		ID:    ulids.New(),
		Name:  name,
		Price: productPrice,
	}

	return product, nil
}

func (product Product) SameVersion(version int64) bool {
	return product.Version == int64(version)
}

func (product Product) AddStock(stockIn int64, createdAt time.Time) (Product, ProductStock, error) {
	if stockIn < 0 {
		return Product{}, ProductStock{}, fmt.Errorf("stock in must be positive")
	}

	stock := ProductStock{
		ID:        ulids.New(),
		ProductID: product.ID,
		StockIn:   stockIn,
		CreatedAt: createdAt,
	}

	product.Stocks = append(product.Stocks, stock)
	return product, stock, nil
}

func (product Product) ReduceStock(stockOut int64, createdAt time.Time) (Product, ProductStock, error) {
	if stockOut < 0 {
		return Product{}, ProductStock{}, fmt.Errorf("stock out must be positive")
	}

	if product.CurrentStock() < stockOut {
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

// HaveStock check if product have enough stock for the qty
func (product Product) HaveStock(qty int64) bool {
	return product.CurrentStock() >= qty
}

func (product Product) CurrentStock() int64 {
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
