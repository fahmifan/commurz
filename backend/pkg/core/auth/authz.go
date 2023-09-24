package auth

import "github.com/fahmifan/flycasbin/acl"

const (
	RoleCustomer        acl.Role = "customer"
	RoleBackofficeAdmin acl.Role = "admin_backoffice"
)

// actions
const (
	Read   acl.Action = "read"
	Manage acl.Action = "manage"
)

// resources
const (
	Product acl.Resource = "product"
)

var policies = []acl.Policy{
	Policy(RoleBackofficeAdmin, Read, Product),
	Policy(RoleBackofficeAdmin, Manage, Product),
}

func NewACL() (*acl.ACL, error) {
	return acl.NewACL(policies)
}

func Policy(role acl.Role, act acl.Action, rsc acl.Resource) acl.Policy {
	return acl.Policy{
		Role:     role,
		Resource: rsc,
		Action:   act,
	}
}
