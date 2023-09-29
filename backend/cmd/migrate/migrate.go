package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fahmifan/commurz/pkg/config"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func run() error {
	config.Parse(".env")

	db, err := sql.Open("postgres", config.PostgresDSN())
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}

	return migrateDB(db)
}

func migrateDB(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "sqlcdef/migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Println("migrate up:", n)

	return nil
}
