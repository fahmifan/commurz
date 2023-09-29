package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fahmifan/authme"
	"github.com/fahmifan/authme/backend/httphandler"
	"github.com/fahmifan/authme/backend/smtpmail"
	"github.com/fahmifan/authme/register"
	"github.com/fahmifan/commurz/pkg/config"
	"github.com/fahmifan/commurz/pkg/core/auth"
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

	acl, err := auth.NewACL()
	if err != nil {
		return fmt.Errorf("auth.NewACL: %w", err)
	}

	svc := service.NewService(&service.Config{
		DB:  db,
		ACL: acl,
	})

	mailComposer := register.NewDefaultMailComposer("cs@commurz.com", "commurz")
	smtpClient, err := smtpmail.NewSmtpClient(&smtpmail.Config{
		Host: "localhost",
		Port: 1025,
	})
	if err != nil {
		return fmt.Errorf("smtpmail.NewSmtpClient: %w", err)
	}

	accountHandler := httphandler.NewAccountHandler(httphandler.NewAccountHandlerArg{
		VerificationBaseURL: config.BaseURL() + httphandler.PathVerifyRegister,
		DB:                  db,
		MailComposer:        mailComposer,
		Mailer:              smtpClient,
		Locker:              authme.NewDefaultLocker(),
		CSRFSecret:          []byte(config.CSRFSecret()),
		RegisterRedirectURL: config.FEBaseURL() + "/",
		CSRFSecure:          false,
	})
	cookieHandler := httphandler.NewCookieAuthHandler(httphandler.NewCookieAuthHandlerArg{
		AccountHandler: accountHandler,
		RoutePrefix:    "/api",
		CookieSecret:   []byte(config.CookieSecret()),
	})

	ws := webserver.NewWebserver(svc, config.Port(), cookieHandler)

	return ws.Run()
}
