package service

import (
	"context"
	"database/sql"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/auth"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/commurz/pkg/logs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
	"github.com/fahmifan/ulids"
)

func (service *Service) ListBackofficeProducts(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ListBackofficeProductsRequest],
) (*connect.Response[commurzpbv1.ListBackofficeProductsResponse], error) {
	if err := service.can(ctx, auth.Manage, auth.Product); err != nil {
		return nil, err
	}

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

func (service *Service) FindProductByID(
	ctx context.Context,
	req *connect.Request[commurzpbv1.FindByIDRequest],
) (*connect.Response[commurzpbv1.Product], error) {
	// TODO: add authz

	id, err := ulids.Parse(req.Msg.GetId())
	if err != nil {
		logs.ErrCtx(ctx, err, "FindProductByID: Parse")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	productReader := pkgproduct.ProductReader{}
	product, err := productReader.FindProductByID(ctx, service.DB, id)
	if err != nil {
		logs.ErrCtx(ctx, err, "FindProductByID: FindProductByID")
		return nil, connect.NewError(connect.CodeInternal, ErrInternal)
	}

	res := &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
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
