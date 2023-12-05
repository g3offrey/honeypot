package config

import (
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"3000"`
	DBPath   string `env:"DB_PATH" envDefault:"honeypot.db"`
	logLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

func (c *Config) SlogLevel() slog.Level {
	switch c.logLevel {
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func ParseConfig() (Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("can't parse config: %w", err)
	}

	return cfg, nil
}
