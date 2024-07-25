package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config represents the application configuration.
type Config struct {
	App      *AppCfg
	Database *DatabaseCfg
}

type (
	AppCfg struct {
		HttpServerPort int    `envconfig:"HTTP_SERVER_PORT" default:"8888"`
		ServiceName    string `envconfig:"SERVICE_NAME" default:"tix-devops-api"`
		AppEnv         string `envconfig:"APP_ENV" default:"production"`
	}

	DatabaseCfg struct {
		DBName string `envconfig:"MAIN_DBNAME" default:"dbname" required:"true"`
		DBUser string `envconfig:"MAIN_DBUSER" default:"dbuser" required:"true"`
		DBPass string `envconfig:"MAIN_DBPASS" default:"dbpass" required:"true"`
		Host   string `envconfig:"MAIN_HOST" default:"localhost" required:"true"`
		Port   string `envconfig:"MAIN_PORT" default:"27017" required:"true"`
	}
)

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load .env file")
	}

	cfg := &Config{
		App:      &AppCfg{},
		Database: &DatabaseCfg{},
	}

	// Process environment variables for each configuration section
	if err := loadConfig("APP", cfg.App); err != nil {
		return nil, errors.Wrap(err, "failed to load App configuration")
	}

	if err := loadConfig("DATABASE", cfg.Database); err != nil {
		return nil, errors.Wrap(err, "failed to load Database configuration")
	}

	return cfg, nil
}

// loadConfig loads the configuration for a given section using envconfig.
func loadConfig(prefix string, cfg interface{}) error {
	if err := envconfig.Process(prefix, cfg); err != nil {
		return errors.Wrapf(err, "failed to process %s configuration", prefix)
	}
	return nil
}
