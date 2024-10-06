package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	KafkaServers       string
	KafkaOffset        string
	KafkaConsumerGroup string
	KafkaAcks          string
	ServerAddr         string
	DbDSN              string
}

func NewConfig(l *zap.Logger) *Config {

	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, "../../.env")
	_ = godotenv.Load(envPath)
	l.Info(envPath)
	return &Config{
		KafkaServers:       os.Getenv("K_SERVERS"),
		KafkaOffset:        os.Getenv("K_OFFSET"),
		KafkaConsumerGroup: os.Getenv("K_CONSUMER_GROUP"),
		KafkaAcks:          os.Getenv("K_ACKS"),
		ServerAddr:         os.Getenv("SERVER_ADDR"),
		DbDSN:              os.Getenv("DB_DSN"),
	}
}
