// Package service is the API entry point,
// it orchestrates between business logic, 3rd party services, and persistence storage.
package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/pb/commurz/v1/commurzv1connect"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
)

type Config struct {
	DB *sql.DB
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
