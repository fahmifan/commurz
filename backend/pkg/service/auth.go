package service

import (
	"context"

	"github.com/fahmifan/ulids"
)

type User struct {
	ID ulids.ULID
}
type CtxKey string

const UserCtxKey CtxKey = "user"

func UserFromCtx(ctx context.Context) (User, bool) {
	user, ok := ctx.Value(UserCtxKey).(User)
	if !ok {
		return User{}, false
	}

	return user, true
}

func CtxWithUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, UserCtxKey, user)
}
