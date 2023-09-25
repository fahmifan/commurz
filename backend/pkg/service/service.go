// Package service is the API entry point,
// it orchestrates between business logic, 3rd party services, and persistence storage.
package service

import (
	"database/sql"

	"github.com/fahmifan/commurz/pkg/core"
	"github.com/fahmifan/commurz/pkg/core/auth/auth_query"
	"github.com/fahmifan/commurz/pkg/core/order_inventory/order_inventory_cmd"
	"github.com/fahmifan/commurz/pkg/core/order_inventory/order_inventory_query"
	"github.com/fahmifan/commurz/pkg/core/pkgmoney"
	"github.com/fahmifan/commurz/pkg/core/storefront/storefront_query"
	"github.com/fahmifan/commurz/pkg/core/user_profile/user_profile_query"
	"github.com/fahmifan/commurz/pkg/pb/commurz/v1/commurzv1connect"
	"github.com/fahmifan/flycasbin/acl"
)

type Config struct {
	DB  *sql.DB
	ACL *acl.ACL
}

var _ commurzv1connect.CommurzServiceHandler = &Service{}

type Service struct {
	_ struct{}
	*Config
	*order_inventory_cmd.OrderInventoryCmd
	*order_inventory_query.OrderInventoryQuery
	*storefront_query.StoreFrontQuery
	*user_profile_query.UserQuery
	*auth_query.AuthQuery
}

func init() {
	pkgmoney.Divider = 100
}

func NewService(config *Config) *Service {
	coreCtx := &core.Ctx{
		DB:  config.DB,
		ACL: config.ACL,
	}

	return &Service{
		Config:              config,
		OrderInventoryCmd:   &order_inventory_cmd.OrderInventoryCmd{Ctx: coreCtx},
		StoreFrontQuery:     &storefront_query.StoreFrontQuery{Ctx: coreCtx},
		OrderInventoryQuery: &order_inventory_query.OrderInventoryQuery{Ctx: coreCtx},
		UserQuery:           &user_profile_query.UserQuery{Ctx: coreCtx},
		AuthQuery:           &auth_query.AuthQuery{Ctx: coreCtx},
	}
}
