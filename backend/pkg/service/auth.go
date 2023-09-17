package service

import (
	"context"

	"github.com/fahmifan/authme/auth"
	"github.com/fahmifan/authme/backend/httphandler"
	"github.com/fahmifan/flycasbin/acl"
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Role acl.Role
}

type CtxKey string

const UserCtxKey CtxKey = "user"

func UserFromCtx(ctx context.Context) (auth.UserSession, bool) {
	user, ok := httphandler.GetUser(ctx)
	if !ok {
		return auth.UserSession{}, false
	}

	return user, true
}

func CtxWithUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, UserCtxKey, user)
}
