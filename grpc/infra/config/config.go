package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoadConfig(l *zap.Logger) *Config {
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, ".env")

	_ = godotenv.Load(envPath)

	cfg := &Config{
		GrpcAddr: os.Getenv("GRPCADDR"),
		Dsn:      os.Getenv("DSN"),
	}
	l.Info("in config", zap.Any("cfg", cfg))
	return cfg
}
