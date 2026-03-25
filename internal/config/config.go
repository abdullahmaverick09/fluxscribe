package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	Env  string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() (*Config, error) {
	conf := &Config{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUser:     os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}

	if err := conf.validate(); err != nil {
		return nil, fmt.Errorf("Configuration error: %w", err)
	}

	return conf, nil
}

func (conf *Config) validate() error {
	if conf.Port == "" {
		return fmt.Errorf("PORT is required")
	}

	if conf.Env == "" {
		return fmt.Errorf("ENV is required")
	}

	if conf.DBHost == "" {
		return fmt.Errorf("DB_HOST is required")
	}

	if conf.DBPort == "" {
		return fmt.Errorf("DB_PORT is required")
	}

	if conf.DBUser == "" {
		return fmt.Errorf("DB_USER is required")
	}

	if conf.DBPassword == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	if conf.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	return nil

}
