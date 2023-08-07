package config

import (
	"context"
	"os"
	"sync"

	"github.com/fahmifan/commurz/pkg/logs"
	"github.com/joho/godotenv"
)

type Config struct {
	CSRFSecret       string `env:"CSRF_SECRET"`
	EnableSecureCSRF string `env:"CSRF_ENABLE_SECURE"`
	CookieSecret     string `env:"COOKIE_SECRET"`
	Debug            string `env:"DEBUG"`
	ENV              string `env:"ENV"`
	Port             string `env:"PORT"`
	BaseURL          string `env:"BASE_URL"`
}

var once sync.Once
var cfg Config

func Parse(filename string) {
	once.Do(func() {
		if filename == "" {
			filename = ".env"
		}
		if loadErr := godotenv.Load(filename); loadErr != nil {
			logs.InfoCtx(context.Background(), "config", "failed load env from", filename, loadErr.Error())
		}

		cfg.CSRFSecret = os.Getenv("CSRF_SECRET")
		cfg.EnableSecureCSRF = os.Getenv("CSRF_ENABLE_SECURE")
		cfg.CookieSecret = os.Getenv("COOKIE_SECRET")
		cfg.Debug = os.Getenv("DEBUG")
		cfg.ENV = os.Getenv("ENV")
		cfg.Port = os.Getenv("PORT")
		cfg.BaseURL = os.Getenv("BASE_URL")
	})
}

func CookieSecret() string {
	return cfg.CookieSecret
}

func Debug() bool {
	return cfg.Debug == "true"
}

func CSRFSecret() string {
	return cfg.CSRFSecret
}

func CSRFEnableSecure() bool {
	return cfg.EnableSecureCSRF == "true"
}

func Env() string {
	return cfg.ENV
}

func IsDevENV() bool {
	return cfg.ENV == "development" && cfg.ENV != "production"
}

func Port() string {
	return cfg.Port
}

func BaseURL() string {
	return cfg.BaseURL
}
