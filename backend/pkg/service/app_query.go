package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/auth"
	"github.com/fahmifan/commurz/pkg/core/pkgorder"
	"github.com/fahmifan/commurz/pkg/core/pkgproduct"
	"github.com/fahmifan/commurz/pkg/logs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
	"github.com/fahmifan/commurz/pkg/sqlcs"
	"github.com/google/uuid"
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

func (service *Service) FindCartByUserToken(
	ctx context.Context,
	req *connect.Request[commurzpbv1.Empty],
) (res *connect.Response[commurzpbv1.Cart], err error) {
	user, ok := auth.UserFromCtx(ctx)
	if !ok {
		return &connect.Response[commurzpbv1.Cart]{}, nil
	}

	userID, err := uuid.Parse(user.GUID)
	if err != nil {
		return nil, ErrInternal
	}

	cartReader := pkgorder.CartReader{}
	cart, err := cartReader.FindCartByUserID(ctx, service.DB, userID)
	if isNotFoundErr(err) {
		return nil, ErrNotFound
	}
	if err != nil {
		logs.ErrCtx(ctx, err, "FindCartByUserToken: FindCartByUserID")
		return nil, ErrInternal
	}

	res = &connect.Response[commurzpbv1.Cart]{
		Msg: protoserde.FromCartPkg(cart),
	}

	return res, nil
}

func (service *Service) FindUserByToken(
	ctx context.Context,
	req *connect.Request[commurzpbv1.Empty],
) (res *connect.Response[commurzpbv1.User], err error) {
	user, ok := auth.UserFromCtx(ctx)
	if !ok {
		return &connect.Response[commurzpbv1.User]{}, nil
	}

	res = &connect.Response[commurzpbv1.User]{
		Msg: &commurzpbv1.User{
			Id:    user.GUID,
			Email: user.Email,
		},
	}

	return res, nil
}
