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

func (service *StoreFrontQuery) FindAllProductListing(
	ctx context.Context,
	req *connect.Request[commurzv1.FindAllProductListingRequest],
) (res *connect.Response[commurzv1.FindAllProductListingResponse], err error) {
	productListingReader := storefront.ProductListingReader{}
	productListings, count, err := productListingReader.FindAllProducts(ctx, service.DB, sqlcs.FindAllProductListingParams{
		SetName:    req.Msg.Name != "",
		Name:       core.NullString(req.Msg.Name),
		PageOffset: core.PageOffset(req.Msg.Pagination.Page, req.Msg.Pagination.Size),
		PageLimit:  core.PageLimit(req.Msg.Pagination.Size),
	})

	res = &connect.Response[commurzv1.FindAllProductListingResponse]{
		Msg: &commurzv1.FindAllProductListingResponse{
			Products: ListFromProductListingsPkg(productListings),
			Count:    int32(count),
		},
	}

	return
}
