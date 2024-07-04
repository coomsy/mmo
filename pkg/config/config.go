package config

import (
	"github.com/caarlos0/env"
)

var AppConfig Config

type Config struct {
	// is port
	Port        int    `env:"MMO_PORT" envDefault:"8080"`
	Environment string `env:"MMO_ENVIRONMENT" envDefault:"dev"`
	Debug       bool   `env:"MMO_DEBUG" envDefault:"false"`

	DBDriver string `env:"MMO_DB_DRIVER" envDefault:"sqllite"`
	DBUser   string `env:"MMO_DB_USER" envDefault:"gm"`
	DBPass   string `env:"MMO_DB_PASS" envDefault:"password"`
}

func InitAppConfig() error {
	// TODO: std Flags has a env load function, use it
	err := env.Parse(&AppConfig)
	if err != nil {
		return err
	}

	return nil
}
