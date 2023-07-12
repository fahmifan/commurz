package pkguser

import (
	"context"
	"fmt"

	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/oklog/ulid/v2"
)

type User struct {
	ID    ulids.ULID
	Email string
}

func NewUser(email string) User {
	return User{
		ID:    ulids.New(),
		Email: email,
	}
}

func userFromSqlc(xuser sqlcs.User) User {
	return User{
		ID:    mustParseULID(xuser.ID),
		Email: xuser.Email,
	}
}

type UserRepository struct{}

func (UserRepository) FindUserByID(ctx context.Context, tx sqlcs.DBTX, id ulids.ULID) (User, error) {
	queries := sqlcs.New(tx)
	xuser, err := queries.FindUserByID(ctx, id.String())
	if err != nil {
		return User{}, fmt.Errorf("[FindUserByID] FindUserByID: %w", err)
	}

	return userFromSqlc(xuser), nil
}

func (UserRepository) CreateUser(ctx context.Context, tx sqlcs.DBTX, user User) (User, error) {
	queries := sqlcs.New(tx)
	xuser, err := queries.CreateUser(ctx, sqlcs.CreateUserParams{
		ID:    user.ID.String(),
		Email: user.Email,
	})
	if err != nil {
		return User{}, fmt.Errorf("[CreateUser] CreateUser: %w", err)
	}

	return userFromSqlc(xuser), nil
}

func mustParseULID(s string) ulids.ULID {
	return ulids.ULID{ULID: ulid.MustParse(s)}
}
