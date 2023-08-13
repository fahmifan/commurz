package pkguser

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword string
}

var (
	ErrHashingPassword = errors.New("error hashing password")
)

func NewUser(email, plainPassword string) (User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return User{}, errors.New("invalid email address")
	}

	// simple rule for now
	if len(plainPassword) < 8 {
		return User{}, errors.New("password must be at least 8 characters")
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("[NewUser] %w: %w", ErrHashingPassword, err)
	}

	user := User{
		ID:             uuid.New(),
		Email:          email,
		HashedPassword: string(bcryptPassword),
	}

	return user, nil
}

func (u User) CanLogin(plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(plainPassword))
	if err != nil {
		return fmt.Errorf("[User-CanLogin] CompareHashAndPassword: %w", err)
	}

	return nil
}

func userFromSqlc(xuser sqlcs.User) User {
	return User{
		ID:    xuser.ID,
		Email: xuser.Email,
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
