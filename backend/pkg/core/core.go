// Package core contains the core logic of the application.
// It sub packages are prefix with pkg to avoid name collision with variable from the pkg,
// e.g. with this prefix we can create a variable like this `product := pkgproduct.Product{}`
package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/fahmifan/commurz/pkg/core/auth"
	"github.com/fahmifan/flycasbin/acl"
)

var (
	ErrSvcInternal        = connect.NewError(connect.CodeInternal, errors.New("internal error"))
	ErrSvcNotFound        = connect.NewError(connect.CodeNotFound, errors.New("not found"))
	ErrSvcUnauthorized    = connect.NewError(connect.CodePermissionDenied, errors.New("unauthorized"))
	ErrSvcUnauthenticated = connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
)

type Ctx struct {
	DB  *sql.DB
	ACL *acl.ACL
}

func (service *Ctx) CanUser(ctx context.Context, act acl.Action, rsc acl.Resource) error {
	user, ok := auth.UserFromCtx(ctx)
	if !ok {
		return ErrSvcUnauthenticated
	}

	err := service.ACL.Can(user.Role, act, rsc)
	if err != nil {
		return ErrSvcUnauthorized
	}

	return nil
}

func Transaction(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	txErr := fn(tx)
	// success
	if txErr == nil {
		return tx.Commit()
	}

	if err := tx.Rollback(); err != nil {
		return fmt.Errorf("txErr: %w: rollabck err: %w", txErr, err)
	}

	return txErr
}

func IsDBNotFoundErr(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func PageLimit(limit int32) int32 {
	if limit <= 0 {
		limit = 10
	}
	return int32(limit)
}

func PageOffset(page, size int32) int32 {
	if size <= 0 {
		size = 10
	}
	return int32(PageLimit(page)-1) * int32(size)
}

func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}
