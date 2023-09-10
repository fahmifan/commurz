package service

import (
	"context"
	"database/sql"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/commurz/pkg/logs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
)

func (service *Service) ListBackofficeProducts(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ListBackofficeProductsRequest],
) (*connect.Response[commurzpbv1.ListBackofficeProductsResponse], error) {
	productReader := pkgproduct.ProductBackofficeReader{}

	arg := NewListProductArg(req.Msg)

	products, count, err := productReader.FindAllProducts(ctx, service.DB, arg)
	if err != nil {
		logs.ErrCtx(ctx, err, "ListProducts: FindAllProducts")
		return nil, connect.NewError(connect.CodeInternal, ErrInternal)
	}

	res := &connect.Response[commurzpbv1.ListBackofficeProductsResponse]{
		Msg: &commurzpbv1.ListBackofficeProductsResponse{
			Products: protoserde.ListFromProductPkg(products),
			Count:    int32(count),
		},
	}

	return res, nil
}

func NewListProductArg(arg *commurzpbv1.ListBackofficeProductsRequest) sqlcs.FindAllProductsForBackofficeParams {
	return sqlcs.FindAllProductsForBackofficeParams{
		Name:       NullString(arg.Name),
		SetName:    arg.Name != "",
		PageLimit:  PageLimit(arg.Pagination.Size),
		PageOffset: PageOffset(arg.Pagination.Page, arg.Pagination.Size),
	}
}

func PageLimit(limit int32) int32 {
	if limit <= 0 {
		limit = 10
	}
	return int32(limit)
}

func PageOffset(page, size int32) int32 {
	if size <= 0 {
		size = 10
	}
	return int32(PageLimit(page)-1) * int32(size)
}

func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}
