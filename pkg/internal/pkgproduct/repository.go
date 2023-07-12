package pkgproduct

import (
	"context"
	"fmt"

	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/commurz/pkg/preloads"
	"github.com/fahmifan/ulids"
	"github.com/samber/lo"
)

type ProductRepository struct{}

func (ProductRepository) FindProductByID(ctx context.Context, tx sqlcs.DBTX, id ulids.ULID) (Product, error) {
	queries := sqlcs.New(tx)
	sqlcProduct, err := queries.FindProductByID(ctx, id.String())
	if err != nil {
		return Product{}, fmt.Errorf("[FindProductByID] FindProductByID: %w", err)
	}

	productStocks, err := ProductRepository{}.FindAllProductStocksByIDs(ctx, tx, []ulids.ULID{id})
	if err != nil {
		return Product{}, fmt.Errorf("[FindProductByID] FindAllProductStocksByIDs: %w", err)
	}

	product := productFromSqlc(sqlcProduct, 0)
	product.Stocks = productStocks

	return product, nil
}

func (repo ProductRepository) FindProductsByIDs(ctx context.Context, tx sqlcs.DBTX, productIDs []ulids.ULID) ([]Product, error) {
	query := sqlcs.New(tx)

	xproducts, err := query.FindAllProductsByIDs(ctx, stringULIDs(productIDs))
	if err != nil {
		return nil, fmt.Errorf("[FindProductsByIDs] FindAllProductsByIDs: %w", err)
	}

	products, err := preloads.PreloadMany[Product, ProductStock, ulids.ULID]{
		Targets:   lo.Map(xproducts, productFromSqlc),
		RefItem:   func(item ProductStock) ulids.ULID { return item.ProductID },
		RefTarget: func(target Product) ulids.ULID { return target.ID },
		SetItem:   func(target *Product, items []ProductStock) { target.Stocks = items },
		FetchItems: func() ([]ProductStock, error) {
			return repo.FindAllProductStocksByIDs(ctx, tx, productIDs)
		},
	}.Preload()
	if err != nil {
		return nil, fmt.Errorf("[FindProductsByIDs] FindAllProductStocksByIDs: %w", err)
	}

	return products, nil
}

func (ProductRepository) FindAllProductStocksByIDs(ctx context.Context, tx sqlcs.DBTX, productIDs []ulids.ULID) ([]ProductStock, error) {
	queries := sqlcs.New(tx)

	xstocks, err := queries.FindAllProductStocksByIDs(ctx, stringULIDs(productIDs))
	if err != nil {
		return nil, fmt.Errorf("[FindAllProductStocksByIDs] FindAllProductStocksByIDs: %w", err)
	}

	return lo.Map(xstocks, productStockFromSqlc), nil
}

func (repo ProductRepository) SaveProduct(ctx context.Context, tx sqlcs.DBTX, product Product) (Product, error) {
	queries := sqlcs.New(tx)

	xproduct, err := queries.SaveProduct(ctx, sqlcs.SaveProductParams{
		ID:    product.ID.String(),
		Name:  product.Name,
		Price: int64(product.Price),
	})
	if err != nil {
		return Product{}, fmt.Errorf("[Save] SaveProduct: %w", err)
	}

	return productFromSqlc(xproduct, 0), nil
}

func (repo ProductRepository) UpdateProduct(ctx context.Context, tx sqlcs.DBTX, product Product) (Product, error) {
	query := sqlcs.New(tx)

	xproduct, err := query.UpdateProduct(ctx, sqlcs.UpdateProductParams{
		ID:    product.ID.String(),
		Name:  product.Name,
		Price: int64(product.Price),
	})
	if err != nil {
		return Product{}, fmt.Errorf("[UpdateProduct] UpdateProduct: %w", err)
	}

	return productFromSqlc(xproduct, 0), nil
}

func (repo ProductRepository) SaveProductStock(ctx context.Context, tx sqlcs.DBTX, stock ProductStock) (ProductStock, error) {
	query := sqlcs.New(tx)

	xstock, err := query.CreateProductStock(ctx, sqlcs.CreateProductStockParams{
		ID:        stock.ID.String(),
		ProductID: stock.ProductID.String(),
		StockIn:   int64(stock.StockIn),
		StockOut:  int64(stock.StockOut),
	})
	if err != nil {
		return ProductStock{}, fmt.Errorf("[AddProductStock] CreateProductStock: %w", err)
	}

	return productStockFromSqlc(xstock, 0), nil
}
