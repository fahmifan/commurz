package service

import (
	"context"

	"github.com/fahmifan/commurz/pkg/auth"
	"github.com/fahmifan/flycasbin/acl"
)

func (service *Service) can(ctx context.Context, act acl.Action, rsc acl.Resource) error {
	user, ok := auth.UserFromCtx(ctx)
	if !ok {
		return ErrUnauthenticated
	}

	err := service.ACL.Can(user.Role, act, rsc)
	if err != nil {
		return ErrUnauthorized
	}

	return nil
}
