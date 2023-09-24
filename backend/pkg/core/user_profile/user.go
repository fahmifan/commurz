package user_profile

import (
	"context"
	"errors"
	"fmt"

	"github.com/fahmifan/commurz/pkg/sqlcs"
	"github.com/fahmifan/flycasbin/acl"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type User struct {
	ID    uuid.UUID
	Email string
	Role  acl.Role
}

var (
	ErrHashingPassword = errors.New("error hashing password")
)

func userFromSqlc(xuser sqlcs.User) User {
	return User{
		ID:    xuser.ID,
		Email: xuser.Email,
		Role:  acl.Role(xuser.Role),
	}
}

type UserReader struct{}

func (UserReader) FindUserByID(ctx context.Context, tx sqlcs.DBTX, id uuid.UUID) (User, error) {
	queries := sqlcs.New(tx)
	xuser, err := queries.FindUserByID(ctx, id)
	if err != nil {
		return User{}, fmt.Errorf("[FindUserByID] FindUserByID: %w", err)
	}

	return userFromSqlc(xuser), nil
}

func (UserReader) FindAllUsers(ctx context.Context, tx sqlcs.DBTX) ([]User, error) {
	queries := sqlcs.New(tx)
	xusers, err := queries.FindAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("[FindAllUsers] FindAllUsers: %w", err)
	}

	users := lo.Map(xusers, func(xuser sqlcs.User, _ int) User {
		return userFromSqlc(xuser)
	})

	return users, nil
}

func (UserReader) FindByEmail(ctx context.Context, tx sqlcs.DBTX, email string) (User, error) {
	queries := sqlcs.New(tx)
	xuser, err := queries.FindUserByEmail(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("[FindByEmail] FindByEmail: %w", err)
	}

	return userFromSqlc(xuser), nil
}

type UserWriter struct{}

func (UserWriter) CreateUser(ctx context.Context, tx sqlcs.DBTX, user User) (User, error) {
	queries := sqlcs.New(tx)
	xuser, err := queries.CreateUser(ctx, sqlcs.CreateUserParams{
		ID:    user.ID,
		Email: user.Email,
	})
	if err != nil {
		return User{}, fmt.Errorf("[CreateUser] CreateUser: %w", err)
	}

	return userFromSqlc(xuser), nil
}
