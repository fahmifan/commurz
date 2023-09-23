package auth

import (
	"context"

	"github.com/fahmifan/authme/auth"
	"github.com/fahmifan/flycasbin/acl"
)

type CtxKey string

const UserCtxKey CtxKey = "user"

type UserAuth struct {
	auth.UserSession
	Role acl.Role
}

func UserFromCtx(ctx context.Context) (UserAuth, bool) {
	user, ok := ctx.Value(UserCtxKey).(UserAuth)
	return user, ok
}

func CtxWithUser(ctx context.Context, user UserAuth) context.Context {
	return context.WithValue(ctx, UserCtxKey, user)
}
