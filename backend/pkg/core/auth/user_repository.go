package auth

import (
	"context"
	"fmt"

	"github.com/fahmifan/commurz/pkg/sqlcs"
	"github.com/fahmifan/flycasbin/acl"
	"github.com/google/uuid"
)

type UserReader struct{}

func (UserReader) FindUserByID(ctx context.Context, tx sqlcs.DBTX, id uuid.UUID) (User, error) {
	queries := sqlcs.New(tx)
	xuser, err := queries.FindUserByID(ctx, id)
	if err != nil {
		return User{}, fmt.Errorf("[FindUserByID] FindUserByID: %w", err)
	}

	return userFromSqlc(xuser), nil
}

func userFromSqlc(xuser sqlcs.User) User {
	return User{
		ID:    xuser.ID,
		Email: xuser.Email,
		Role:  acl.Role(xuser.Role),
	}
}
