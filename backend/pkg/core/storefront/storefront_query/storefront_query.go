package storefront_query

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fahmifan/commurz/pkg/core"
	"github.com/fahmifan/commurz/pkg/core/storefront"
	commurzv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/sqlcs"
)

type StoreFrontQuery struct {
	*core.Ctx
}

func (service *StoreFrontQuery) ListAppProducts(
	ctx context.Context,
	req *connect.Request[commurzv1.ListAppProductsRequest],
) (res *connect.Response[commurzv1.ListAppProductsResponse], err error) {
	productListingReader := storefront.ProductListingReader{}
	productListings, count, err := productListingReader.FindAllProducts(ctx, service.DB, sqlcs.FindAllProductsForAppParams{
		SetName:    req.Msg.Name != "",
		Name:       core.NullString(req.Msg.Name),
		PageOffset: core.PageOffset(req.Msg.Pagination.Page, req.Msg.Pagination.Size),
		PageLimit:  core.PageLimit(req.Msg.Pagination.Size),
	})

	res = &connect.Response[commurzv1.ListAppProductsResponse]{
		Msg: &commurzv1.ListAppProductsResponse{
			Products: ListFromProductListingsPkg(productListings),
			Count:    int32(count),
		},
	}
	return
}
