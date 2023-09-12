package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
)

func (service *Service) ListAppProducts(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ListAppProductsRequest],
) (res *connect.Response[commurzpbv1.ListAppProductsResponse], err error) {
	productListingReader := pkgproduct.ProductListingReader{}
	productListings, count, err := productListingReader.FindAllProducts(ctx, service.DB, sqlcs.FindAllProductsForAppParams{
		SetName:    req.Msg.Name != "",
		Name:       NullString(req.Msg.Name),
		PageOffset: PageOffset(req.Msg.Pagination.Page, req.Msg.Pagination.Size),
		PageLimit:  PageLimit(req.Msg.Pagination.Size),
	})

	res = &connect.Response[commurzpbv1.ListAppProductsResponse]{
		Msg: &commurzpbv1.ListAppProductsResponse{
			Products: protoserde.ListFromProductListingsPkg(productListings),
			Count:    int32(count),
		},
	}
	return
}
