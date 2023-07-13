package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fahmifan/commurz/pkg/service"
	"github.com/fahmifan/commurz/pkg/web/webserver"
	_ "modernc.org/sqlite"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func run() error {
	dsnURI := "file::memory:?mode=memory&cache=shared&journal_mode=wal&_fk=1"
	db, err := sql.Open("sqlite", dsnURI)
	if err != nil {
		return err
	}

	if err := migrate(db); err != nil {
		return err
	}

	svc := service.NewService(&service.Config{
		DB: db,
	})

	ws := webserver.NewWebserver(svc, 8080)

	return ws.Run()
}

func migrate(db *sql.DB) error {
	filename := "sqlcdef/schema.sql"
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(buf)); err != nil {
		return err
	}

	return nil
}
