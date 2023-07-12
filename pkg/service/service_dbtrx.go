package service

import (
	"database/sql"
	"errors"
	"fmt"
)

func Transaction(db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
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

func isNotFoundErr(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
