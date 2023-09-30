package order_inventory

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/fahmifan/commurz/pkg/core"
	"github.com/fahmifan/commurz/pkg/parseutil"
	"github.com/fahmifan/commurz/pkg/preloads"
	"github.com/fahmifan/commurz/pkg/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/samber/lo"
)

type Count = int64

type ProductBackofficeReader struct{}

func (ProductBackofficeReader) FindAllProducts(
	ctx context.Context,
	tx sqlcs.DBTX,
	arg sqlcs.FindAllProductsForBackofficeParams,
) ([]Product, Count, error) {
	query := sqlcs.New(tx)

	xproducts, err := query.FindAllProductsForBackoffice(ctx, arg)
	if err != nil {
		return nil, 0, fmt.Errorf("[FindAllProducts] FindAllProducts: %w", err)
	}

	products := lo.Map(xproducts, productFromSqlc)
	productIDs := lo.Map(products, func(product Product, _ int) ulids.ULID { return product.ID })

	products, err = preloads.PreloadMany[Product, ProductStock, ulids.ULID]{
		Targets:    products,
		RefItem:    func(item ProductStock) ulids.ULID { return item.ProductID },
		RefTarget:  func(target Product) ulids.ULID { return target.ID },
		SetItem:    func(target *Product, items []ProductStock) { target.Stocks = items },
		FetchItems: func() ([]ProductStock, error) { return ProductReader{}.FindAllProductStocksByIDs(ctx, tx, productIDs) },
	}.Preload()
	if err != nil {
		return nil, 0, fmt.Errorf("[FindAllProducts] FindAllProductStocksByIDs: %w", err)
	}

	count, err := query.CountAllProductsForBackoffice(ctx, sqlcs.CountAllProductsForBackofficeParams{
		Name:    arg.Name,
		SetName: arg.SetName,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("[FindAllProducts] CountAllProductsForBackoffice: %w", err)
	}

	return products, count, nil
}

type ProductReader struct{}

func (ProductReader) FindProductByID(ctx context.Context, tx sqlcs.DBTX, id ulids.ULID) (Product, error) {
	queries := sqlcs.New(tx)
	sqlcProduct, err := queries.FindProductByID(ctx, id.String())
	if err != nil {
		return Product{}, fmt.Errorf("[FindProductByID] FindProductByID: %w", err)
	}

	productStocks, err := ProductReader{}.FindAllProductStocksByIDs(ctx, tx, []ulids.ULID{id})
	if err != nil {
		return Product{}, fmt.Errorf("[FindProductByID] FindAllProductStocksByIDs: %w", err)
	}

	product := productFromSqlc(sqlcProduct, 0)
	product.Stocks = productStocks

	return product, nil
}

func (ProductReader) FindProductByIDForUpdate(ctx context.Context, tx sqlcs.DBTX, id ulids.ULID) (Product, error) {
	queries := sqlcs.New(tx)
	sqlcProduct, err := queries.FindAllProductByIDFroUpdate(ctx, id.String())
	if err != nil {
		return Product{}, fmt.Errorf("[FindProductByID] FindProductByID: %w", err)
	}

	productStocks, err := ProductReader{}.FindAllProductStocksByIDs(ctx, tx, []ulids.ULID{id})
	if err != nil {
		return Product{}, fmt.Errorf("[FindProductByID] FindAllProductStocksByIDs: %w", err)
	}

	product := productFromSqlc(sqlcProduct, 0)
	product.Stocks = productStocks

	return product, nil
}

func (repo ProductReader) FindProductsByIDs(ctx context.Context, tx sqlcs.DBTX, productIDs []ulids.ULID) ([]Product, error) {
	query := sqlcs.New(tx)

	xproducts, err := query.FindAllProductsByIDs(ctx, parseutil.StringULIDs(productIDs))
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

func (repo ProductReader) FindAllProductsByIDslockProductStock(ctx context.Context, tx sqlcs.DBTX, productIDs []ulids.ULID) ([]Product, error) {
	query := sqlcs.New(tx)

	xproducts, err := query.FindAllProductsByIDs(ctx, parseutil.StringULIDs(productIDs))
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

func (ProductReader) FindAllProductStocksByIDs(ctx context.Context, tx sqlcs.DBTX, productIDs []ulids.ULID) ([]ProductStock, error) {
	if len(productIDs) == 0 {
		return nil, nil
	}

	queries := sqlcs.New(tx)

	xstocks, err := queries.FindAllProductStocksByIDs(ctx, parseutil.StringULIDs(productIDs))
	if err != nil {
		return nil, fmt.Errorf("[FindAllProductStocksByIDs] FindAllProductStocksByIDs: %w", err)
	}

	return lo.Map(xstocks, productStockFromSqlc), nil
}

type ProductWriter struct{}

func (repo ProductWriter) SaveProduct(ctx context.Context, tx sqlcs.DBTX, product Product) (Product, error) {
	queries := sqlcs.New(tx)

	xproduct, err := queries.SaveProduct(ctx, sqlcs.SaveProductParams{
		ID:    product.ID.String(),
		Name:  product.Name,
		Price: product.Price.Value(),
	})
	if err != nil {
		return Product{}, fmt.Errorf("[Save] SaveProduct: %w", err)
	}

	return productFromSqlc(xproduct, 0), nil
}

func (repo ProductWriter) UpdateProduct(ctx context.Context, tx sqlcs.DBTX, product Product) (Product, error) {
	query := sqlcs.New(tx)

	xproduct, err := query.UpdateProduct(ctx, sqlcs.UpdateProductParams{
		ID:             product.ID.String(),
		Name:           product.Name,
		Price:          product.Price.Value(),
		LatestStock:    product.CurrentStock(),
		CurrentVersion: product.Version,
	})
	if err != nil {
		return Product{}, fmt.Errorf("[UpdateProduct] UpdateProduct: %w", err)
	}

	return productFromSqlc(xproduct, 0), nil
}

func (repo ProductWriter) SaveProductStock(ctx context.Context, tx sqlcs.DBTX, stock ProductStock) (ProductStock, error) {
	query := sqlcs.New(tx)

	xstock, err := query.CreateProductStock(ctx, sqlcs.CreateProductStockParams{
		ID:        stock.ID.String(),
		ProductID: stock.ProductID.String(),
		StockIn:   int64(stock.StockIn),
		StockOut:  int64(stock.StockOut),
	})
	if err != nil {
		return ProductStock{}, fmt.Errorf("[SaveProductStock] CreateProductStock: %w", err)
	}

	return productStockFromSqlc(xstock, 0), nil
}

func (repo ProductWriter) BumpProductVersion(ctx context.Context, tx sqlcs.DBTX, product Product) (Product, error) {
	query := sqlcs.New(tx)

	updated, err := query.BumpProductVersion(ctx, sqlcs.BumpProductVersionParams{
		ID:             product.ID.String(),
		CurrentVersion: product.Version,
	})
	if err != nil {
		return Product{}, fmt.Errorf("[BumpProductVersion] BumpProductVersion: %w", err)
	}

	product.Version = updated.Version

	return product, nil
}

func (repo ProductWriter) BulkBumpProductVersion(ctx context.Context, tx sqlcs.DBTX, products []Product) ([]Product, error) {
	updatedProducts := make([]Product, len(products))
	for i := range products {
		updatedProduct, err := repo.BumpProductVersion(ctx, tx, products[i])
		if err != nil {
			return nil, fmt.Errorf("[BulkBumpProductVersion] BumpProductVersion: %w", err)
		}

		updatedProducts[i] = updatedProduct
	}

	return updatedProducts, nil
}

func (repo ProductWriter) BulkSaveProductLatestStock(ctx context.Context, tx sqlcs.DBTX, products []Product) ([]Product, error) {
	vals := make([]string, len(products))
	for i, prod := range products {
		vals[i] = fmt.Sprintf("('%s', %d, %d)",
			prod.ID.String(),
			prod.CurrentStock(),
			prod.Version,
		)
	}

	tmpVal := strings.Join(vals, ",")

	query := fmt.Sprintf(`--sql
		UPDATE products
		SET 
			latest_stock = tmp_val.tmp_latest_stock,
			version = version + 1
		FROM (VALUES
			%s 
		) AS tmp_val (tmp_id, tmp_latest_stock, tmp_version)
		WHERE 
			products.id = tmp_val.tmp_id
			AND products.version = tmp_val.tmp_version
		RETURNING %s
		`, tmpVal, productFields)

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("[BulkSaveProductLatestStock] QueryContext: %w", err)
	}
	defer rows.Close()

	xproducts, err := scanProducts(rows)
	if err != nil {
		return nil, fmt.Errorf("[BulkSaveProductLatestStock] scanProducts: %w", err)
	}

	products = lo.Map(xproducts, productFromSqlc)
	return products, nil
}

func (repo ProductWriter) BulkSaveProductStocks(ctx context.Context, tx sqlcs.DBTX, stocks []ProductStock) ([]ProductStock, error) {
	instBuilder := sq.Insert("product_stocks").
		Columns("id", "product_id", "stock_in", "stock_out")

	for _, stock := range stocks {
		instBuilder = instBuilder.Values(
			stock.ID.String(),
			stock.ProductID.String(),
			stock.StockIn,
			stock.StockOut,
		)
	}

	instBuilder = instBuilder.Suffix("RETURNING id, product_id, stock_in, stock_out, created_at")

	query, args, err := instBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("[bulkSaveProductStocks] ToSql: %w", err)
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("[bulkSaveProductStocks] QueryContext: %w", err)
	}
	defer rows.Close()

	xstock, err := scanProductStocks(rows)
	if err != nil {
		return nil, fmt.Errorf("[bulkSaveProductStocks] scanProductStocks: %w", err)
	}

	insertedStocks := lo.Map(xstock, productStockFromSqlc)

	return insertedStocks, nil
}

type ProductStockLocker struct{}

func (ProductStockLocker) LockProductStocks(ctx context.Context, tx sqlcs.DBTX, productIDs []ulids.ULID) error {
	query := sqlcs.New(tx)

	_, err := query.LockProductStock(ctx, parseutil.StringULIDs(productIDs))
	if err != nil && !core.IsDBNotFoundErr(err) {
		return fmt.Errorf("[LockProductStocks] LockProductStocks: %w", err)
	}

	return nil
}

func (ProductStockLocker) CreateProductStockLock(ctx context.Context, tx sqlcs.DBTX, productID ulids.ULID) error {
	query := sqlcs.New(tx)

	_, err := query.CreateProductStockLock(ctx, productID.String())
	if err != nil {
		return fmt.Errorf("[CreateProductStockLock] CreateProductStockLock: %w", err)
	}

	return nil
}

func scanProductStocks(rows *sql.Rows) (products []sqlcs.ProductStock, err error) {
	for rows.Next() {
		var v sqlcs.ProductStock
		if err := rows.Scan(
			&v.ID,
			&v.ProductID,
			&v.StockIn,
			&v.StockOut,
			&v.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, v)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return
}

const productFields = `id, name, price, latest_stock, version`

func scanProducts(rows *sql.Rows) (products []sqlcs.Product, err error) {
	for rows.Next() {
		var v sqlcs.Product
		if err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.Price,
			&v.LatestStock,
			&v.Version,
		); err != nil {
			return nil, err
		}
		products = append(products, v)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return
}
