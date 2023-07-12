// Package service is the API entry point,
// it orchestrates between business logic, 3rd party services, and persistence storage.
package service

import (
	"database/sql"
)

type Config struct {
	db *sql.DB
}

func NewConfig(db *sql.DB) *Config {
	return &Config{db}
}

type Service struct {
	*Config
}

func NewService(config *Config) *Service {
	return &Service{config}
}
