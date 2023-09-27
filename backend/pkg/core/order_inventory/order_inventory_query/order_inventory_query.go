package order_inventory_query

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fahmifan/commurz/pkg/core"
	"github.com/fahmifan/commurz/pkg/core/auth"
	"github.com/fahmifan/commurz/pkg/core/order_inventory"
	"github.com/fahmifan/commurz/pkg/logs"
	commurzv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/google/uuid"
)

type OrderInventoryQuery struct {
	*core.Ctx
}

func NewListProductArg(arg *commurzv1.ListBackofficeProductsRequest) sqlcs.FindAllProductsForBackofficeParams {
	return sqlcs.FindAllProductsForBackofficeParams{
		Name:       core.NullString(arg.Name),
		SetName:    arg.Name != "",
		PageLimit:  core.PageLimit(arg.Pagination.Size),
		PageOffset: core.PageOffset(arg.Pagination.Page, arg.Pagination.Size),
	}
}

func (service *OrderInventoryQuery) ListBackofficeProducts(
	ctx context.Context,
	req *connect.Request[commurzv1.ListBackofficeProductsRequest],
) (*connect.Response[commurzv1.ListBackofficeProductsResponse], error) {
	if err := service.CanUser(ctx, auth.Manage, auth.Product); err != nil {
		return nil, err
	}

	productReader := order_inventory.ProductBackofficeReader{}

	arg := NewListProductArg(req.Msg)

	products, count, err := productReader.FindAllProducts(ctx, service.DB, arg)
	if err != nil {
		logs.ErrCtx(ctx, err, "ListProducts: FindAllProducts")
		return nil, connect.NewError(connect.CodeInternal, core.ErrInternal)
	}

	res := &connect.Response[commurzv1.ListBackofficeProductsResponse]{
		Msg: &commurzv1.ListBackofficeProductsResponse{
			Products: ListFromProductPkg(products),
			Count:    int32(count),
		},
	}

	return res, nil
}

func (service *OrderInventoryQuery) FindCartByUserToken(
	ctx context.Context,
	req *connect.Request[commurzv1.Empty],
) (res *connect.Response[commurzv1.Cart], err error) {
	user, ok := auth.UserFromCtx(ctx)
	if !ok {
		return &connect.Response[commurzv1.Cart]{}, nil
	}

	userID, err := uuid.Parse(user.GUID)
	if err != nil {
		return nil, core.ErrInternal
	}

	cartReader := order_inventory.CartReader{}
	cart, err := cartReader.FindCartByUserID(ctx, service.DB, userID)
	if core.IsNotFoundErr(err) {
		return nil, core.ErrNotFound
	}
	if err != nil {
		logs.ErrCtx(ctx, err, "FindCartByUserToken: FindCartByUserID")
		return nil, core.ErrInternal
	}

	res = &connect.Response[commurzv1.Cart]{
		Msg: FromCartPkg(cart),
	}

	return res, nil
}

func (service *OrderInventoryQuery) FindProductByID(
	ctx context.Context,
	req *connect.Request[commurzv1.FindByIDRequest],
) (*connect.Response[commurzv1.Product], error) {
	// TODO: add authz

	id, err := ulids.Parse(req.Msg.GetId())
	if err != nil {
		logs.ErrCtx(ctx, err, "FindProductByID: Parse")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	productReader := order_inventory.ProductReader{}
	product, err := productReader.FindProductByID(ctx, service.DB, id)
	if err != nil {
		logs.ErrCtx(ctx, err, "FindProductByID: FindProductByID")
		return nil, core.ErrInternal
	}

	res := &connect.Response[commurzv1.Product]{
		Msg: FromProductPkg(product),
	}

	return res, nil
}
