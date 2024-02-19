package env

import (
	"fmt"
	"github.com/Murat993/auth/internal/config"
	"os"
	"strconv"
)

var _ config.PGConfig = (*pgConfig)(nil)

const (
	PG_HOST          = "PG_HOST"
	PG_DATABASE_NAME = "PG_DATABASE_NAME"
	PG_USER          = "PG_USER"
	PG_PASSWORD      = "PG_PASSWORD"
	PG_PORT          = "PG_PORT"
)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (config.PGConfig, error) {
	dbHost := os.Getenv(PG_HOST)
	dbName := os.Getenv(PG_DATABASE_NAME)
	dbUser := os.Getenv(PG_USER)
	dbPassword := os.Getenv(PG_PASSWORD)
	dbPort, _ := strconv.Atoi(os.Getenv(PG_PORT))

	return &pgConfig{
		dsn: fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s", dbHost, dbPort, dbUser, dbName, "disable", dbPassword),
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
