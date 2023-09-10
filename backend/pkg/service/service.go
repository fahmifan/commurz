// Package service is the API entry point,
// it orchestrates between business logic, 3rd party services, and persistence storage.
package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/logs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/pb/commurz/v1/commurzv1connect"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
	"github.com/fahmifan/flycasbin/acl"
	"github.com/google/uuid"
)

type Config struct {
	DB  *sql.DB
	ACL *acl.ACL
}

var _ commurzv1connect.CommurzServiceHandler = &Service{}

type Service struct {
	*Config
}

func NewService(config *Config) *Service {
	return &Service{config}
}

func (service *Service) ListUsers(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ListUsersRequest],
) (*connect.Response[commurzpbv1.ListUsersResponse], error) {
	userRepo := pkguser.UserReader{}
	users, err := userRepo.FindAllUsers(ctx, service.DB)
	if err != nil {
		return nil, fmt.Errorf("[ListUsers] FindAllUsers: %w", err)
	}

	res := &connect.Response[commurzpbv1.ListUsersResponse]{
		Msg: &commurzpbv1.ListUsersResponse{
			Users: protoserde.ListFromUserPkg(users),
		},
	}

	return res, nil
}

func (service *Service) FindUserByID(
	ctx context.Context,
	req *connect.Request[commurzpbv1.FindByIDRequest],
) (*connect.Response[commurzpbv1.User], error) {
	userRepo := pkguser.UserReader{}

	id, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	user, err := userRepo.FindUserByID(ctx, service.DB, id)
	if err != nil {
		logs.ErrCtx(ctx, err, "[FindUserByID] FindUserByID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return &connect.Response[commurzpbv1.User]{
		Msg: protoserde.FromUserPkg(user),
	}, nil
}

type Perm struct {
	Action   acl.Action
	Resource acl.Resource
}

func (service *Service) InternalHasAccess(
	ctx context.Context,
	userID uuid.UUID,
	perms []Perm,
) error {
	userRepo := pkguser.UserReader{}

	user, err := userRepo.FindUserByID(ctx, service.DB, userID)
	if err != nil {
		return logs.ErrWrapCtx(ctx, err, "[InternalHasAccess] FindUserByID")
	}

	fmt.Println("user", user)

	for _, perm := range perms {
		err = service.ACL.Can(user.Role, perm.Action, perm.Resource)
		if err != nil {
			return logs.ErrWrapCtx(ctx, err, "[InternalHasAccess] service.ACL.Can")
		}
	}

	return err
}
