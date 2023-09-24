package user_profile_query

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/core"
	"github.com/fahmifan/commurz/pkg/core/auth"
	"github.com/fahmifan/commurz/pkg/core/user_profile"
	"github.com/fahmifan/commurz/pkg/logs"
	commurzv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/google/uuid"
)

type UserQuery struct {
	*core.Ctx
}

func (service *UserQuery) FindUserByToken(
	ctx context.Context,
	req *connect.Request[commurzv1.Empty],
) (res *connect.Response[commurzv1.User], err error) {
	user, ok := auth.UserFromCtx(ctx)
	if !ok {
		return &connect.Response[commurzv1.User]{}, nil
	}

	res = &connect.Response[commurzv1.User]{
		Msg: &commurzv1.User{
			Id:    user.GUID,
			Email: user.Email,
		},
	}

	return res, nil
}

func (service *UserQuery) ListUsers(
	ctx context.Context,
	req *connect.Request[commurzv1.ListUsersRequest],
) (*connect.Response[commurzv1.ListUsersResponse], error) {
	userRepo := user_profile.UserReader{}
	users, err := userRepo.FindAllUsers(ctx, service.DB)
	if err != nil {
		return nil, fmt.Errorf("[ListUsers] FindAllUsers: %w", err)
	}

	res := &connect.Response[commurzv1.ListUsersResponse]{
		Msg: &commurzv1.ListUsersResponse{
			Users: ListFromUserPkg(users),
		},
	}

	return res, nil
}

func (service *UserQuery) FindUserByID(
	ctx context.Context,
	req *connect.Request[commurzv1.FindByIDRequest],
) (*connect.Response[commurzv1.User], error) {
	userRepo := user_profile.UserReader{}

	id, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	user, err := userRepo.FindUserByID(ctx, service.DB, id)
	if err != nil {
		logs.ErrCtx(ctx, err, "[FindUserByID] FindUserByID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return &connect.Response[commurzv1.User]{
		Msg: FromUserPkg(user),
	}, nil
}
