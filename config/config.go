package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	TimeZone string `env:"TIME_ZONE" validate:"required"`

	AppMode   string `env:"APP_MODE"   validate:"required"`
	AppConfig string `env:"APP_CONFIG" validate:"required"`

	DBHost string `env:"DB_HOST" validate:"required"`
	DBPort int    `env:"DB_PORT" validate:"required,min=0,max=65535"`
	DBUser string `env:"DB_USER" validate:"required"`
	DBPass string `env:"DB_PASS" validate:"required"`
	DBName string `env:"DB_NAME" validate:"required"`

	ModemHost      string `env:"MODEM_HOST"       validate:"required,url"`
	ModemLoginUser string `env:"MODEM_LOGIN_USER" validate:"required"`
	ModemLoginPass string `env:"MODEM_LOGIN_PASS" validate:"required"`
}

func New() (Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %w", err)
	}

	if err := config.validate(); err != nil {
		return Config{}, fmt.Errorf("failed to validate config: %w", err)
	}

	return config, nil
}

func (c Config) validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(c)
	if err != nil {
		return fmt.Errorf("failed to validate config %v: %w", c, err)
	}

	return nil
}
