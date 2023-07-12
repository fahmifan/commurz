package service

import (
	"context"
	"fmt"

	"github.com/fahmifan/commurz/pkg/internal/pkguser"
)

type CreateUserRequest struct {
	Email string `json:"email"`
}

func (service *Service) CreateUser(ctx context.Context, req CreateUserRequest) (user pkguser.User, err error) {
	userRepo := pkguser.UserRepository{}

	user = pkguser.NewUser(req.Email)
	user, err = userRepo.CreateUser(ctx, service.db, user)
	if err != nil {
		return pkguser.User{}, fmt.Errorf("[CreateUser] SaveUser: %w", err)
	}

	// use complete user fields
	user, err = userRepo.FindUserByID(ctx, service.db, user.ID)
	if err != nil {
		return pkguser.User{}, fmt.Errorf("[CreateUser] FindUserByID: %w", err)
	}

	return user, nil
}
