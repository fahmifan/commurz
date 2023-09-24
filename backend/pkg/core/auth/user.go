package auth

import (
	"github.com/fahmifan/flycasbin/acl"
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Email string
	Role  acl.Role
}
