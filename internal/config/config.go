package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	App  AppConfig
	Http HttpConfig
	DB   DatabaseConfig
}

type AppConfig struct {
	Name string
	Env  string
}

type HttpConfig struct {
	Port string
}

type DatabaseConfig struct {
	Url             string
	MaxConns        int
	MinConns        int
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

func LoadConfig() (Config, error) {
	cfg := Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", ""),
			Env:  getEnv("APP_ENV", "development"),
		},
		Http: HttpConfig{
			Port: getEnv("PORT", ":8080"),
		},
		DB: DatabaseConfig{
			Url:             getEnv("DB", ""),
			MaxConns:        10,
			MinConns:        1,
			MaxConnLifetime: 1 * time.Hour,
			MaxConnIdleTime: 30 * time.Minute,
		},
	}

	if cfg.App.Name == "" {
		return Config{}, fmt.Errorf("App name is required, check env")
	}

	if cfg.Http.Port == "" {
		return Config{}, fmt.Errorf("Http port is necessary, check env")
	}

	if cfg.DB.Url == "" {
		return Config{}, fmt.Errorf("DB connection string is required")
	}

	return cfg, nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
