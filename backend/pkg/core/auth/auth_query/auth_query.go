package auth_query

import (
	"context"

	"github.com/fahmifan/commurz/pkg/core"
	"github.com/fahmifan/commurz/pkg/core/auth"
	"github.com/fahmifan/commurz/pkg/logs"
	"github.com/fahmifan/flycasbin/acl"
	"github.com/google/uuid"
)

type AuthQuery struct {
	*core.Ctx
}

type Perm struct {
	Action   acl.Action
	Resource acl.Resource
}

func (service *AuthQuery) InternalHasAccess(
	ctx context.Context,
	userID uuid.UUID,
	perms []Perm,
) error {
	userRepo := auth.UserReader{}

	user, err := userRepo.FindUserByID(ctx, service.DB, userID)
	if err != nil {
		return logs.ErrWrapCtx(ctx, err, "[InternalHasAccess] FindUserByID")
	}

	for _, perm := range perms {
		err = service.ACL.Can(user.Role, perm.Action, perm.Resource)
		if err != nil {
			return logs.ErrWrapCtx(ctx, err, "[InternalHasAccess] service.ACL.Can")
		}
	}

	return err
}

func (service *AuthQuery) InternalFindUserByID(ctx context.Context, id uuid.UUID) (auth.User, error) {
	userRepo := auth.UserReader{}

	user, err := userRepo.FindUserByID(ctx, service.DB, id)
	if err != nil {
		return auth.User{}, logs.ErrWrapCtx(ctx, err, "[InternalFindUserByID] FindUserByID")
	}

	return user, nil
}
