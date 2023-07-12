package service

import (
	"database/sql"
	"errors"
)

func Transaction(db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}

	return tx.Commit()
}

func isNotFoundErr(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
