package storefront

import (
	"context"
	"database/sql"

	"github.com/fahmifan/commurz/pkg/pkgmoney"
	"github.com/fahmifan/commurz/pkg/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/oklog/ulid/v2"
	"github.com/samber/lo"
)

type ProductListing struct {
	ID      ulids.ULID
	Name    string
	Price   pkgmoney.Money
	Version int64
	// LatestStock is a denormalized field calculated from the ProductStocks.
	// It uses eventual consistentcy model.
	LatestStock int64
}

type ProductListingReader struct{}

func (app ProductListingReader) FindAllProducts(ctx context.Context, db *sql.DB, arg sqlcs.FindAllProductListingParams) (products []ProductListing, count int64, err error) {
	query := sqlcs.New(db)
	xproducts, err := query.FindAllProductListing(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	products = lo.Map(xproducts, func(xproduct sqlcs.Product, _ int) ProductListing {
		return ProductListing{
			ID:          ulids.ULID{ULID: ulid.MustParse(xproduct.ID)},
			Name:        xproduct.Name,
			Price:       pkgmoney.New(xproduct.Price),
			Version:     xproduct.Version,
			LatestStock: xproduct.LatestStock,
		}
	})

	return
}
