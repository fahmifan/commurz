package config

import (
	"context"
	"os"
	"strconv"
	"sync"

	"github.com/fahmifan/commurz/pkg/logs"
	"github.com/joho/godotenv"
)

type Config struct {
	CSRFSecret       string `env:"CSRF_SECRET"`
	EnableSecureCSRF string `env:"CSRF_ENABLE_SECURE"`
	CookieSecret     string `env:"COOKIE_SECRET"`
	RedisHost        string `env:"REDIS_HOST"`
	PostgresDSN      string `env:"POSTGRES_DSN"`
	Debug            string `env:"DEBUG"`
	ENV              string `env:"ENV"`
	Port             int    `env:"PORT"`
	BaseURL          string `env:"BASE_URL"`
	FEBaseURL        string `env:"FE_BASE_URL"`
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
		cfg.Port = strInt(os.Getenv("PORT"))
		cfg.BaseURL = os.Getenv("BASE_URL")
		cfg.PostgresDSN = os.Getenv("POSTGRES_DSN")
		cfg.RedisHost = os.Getenv("REDIS_HOST")
		cfg.FEBaseURL = os.Getenv("FE_BASE_URL")
		cfg.CookieSecret = os.Getenv("COOKIE_SECRET")
	})
}

func strInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func CookieSecret() string {
	return cfg.CookieSecret
}

func RedisHost() string {
	return cfg.RedisHost
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

func Port() int {
	return cfg.Port
}

func BaseURL() string {
	return cfg.BaseURL
}

func FEBaseURL() string {
	return cfg.FEBaseURL
}

func PostgresDSN() string {
	return cfg.PostgresDSN
}
