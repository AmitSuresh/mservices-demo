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
		ServerAddr:         os.Getenv("SERVERADDR"),
		GrpcAddr:           os.Getenv("GRPCADDR"),
		KafkaServers:       os.Getenv("K_SERVERS"),
		KafkaOffset:        os.Getenv("K_OFFSET"),
		KafkaConsumerGroup: os.Getenv("K_CONSUMER_GROUP"),
		KafkaAcks:          os.Getenv("K_ACKS"),
	}
	l.Info("in config", zap.Any("cfg", cfg))
	return cfg
}
