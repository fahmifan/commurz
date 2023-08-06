package pkguser

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/ulids"
	"github.com/samber/lo"
)

type User struct {
	ID    ulids.ULID
	Email string
}

func NewUser(email string) (User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return User{}, fmt.Errorf("[NewUser] ParseAddress: %w", err)
	}

	user := User{
		ID:    ulids.New(),
		Email: email,
	}

	return user, nil
}

func userFromSqlc(xuser sqlcs.User) User {
	return User{
		ID:    pkgutil.WeakParseULID(xuser.ID),
		Email: xuser.Email,
	}
}

type UserReader struct{}

func (UserReader) FindUserByID(ctx context.Context, tx sqlcs.DBTX, id ulids.ULID) (User, error) {
	queries := sqlcs.New(tx)
	xuser, err := queries.FindUserByID(ctx, id.String())
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

type UserWriter struct{}

func (UserWriter) CreateUser(ctx context.Context, tx sqlcs.DBTX, user User) (User, error) {
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
