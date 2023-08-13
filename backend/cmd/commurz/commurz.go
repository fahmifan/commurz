package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fahmifan/authme/backend/httphandler"
	"github.com/fahmifan/authme/backend/smtpmail"
	"github.com/fahmifan/authme/register"
	"github.com/fahmifan/commurz/pkg/config"
	"github.com/fahmifan/commurz/pkg/service"
	"github.com/fahmifan/commurz/pkg/web/webserver"
	_ "github.com/lib/pq"
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

	if err := migrate(db); err != nil {
		return err
	}

	svc := service.NewService(&service.Config{
		DB: db,
	})

	mailComposer := register.NewDefaultMailComposer("cs@commurz.com", "commurz")
	smtpClient, err := smtpmail.NewSmtpClient(&smtpmail.Config{
		Host: "localhost",
		Port: 1025,
	})
	if err != nil {
		return fmt.Errorf("smtpmail.NewSmtpClient: %w", err)
	}

	authHandler := httphandler.NewHTTPHandler(httphandler.NewHTTPHandlerArg{
		RedisHost:           config.RedisHost(),
		DB:                  db,
		JWTSecret:           []byte("secret"),
		VerificationBaseURL: "http://localhost:8080" + httphandler.PathVerifyRegister,
		MailComposer:        mailComposer,
		Mailer:              smtpClient,
	})

	if err := authHandler.MigrateUp(); err != nil {
		return fmt.Errorf("authHandler.MigrateUp: %w", err)
	}

	authRouter, err := authHandler.Router()
	if err != nil {
		return fmt.Errorf("authHandler.Router: %w", err)
	}

	ws := webserver.NewWebserver(svc, 8080, authRouter)

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
