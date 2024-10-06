package queue

import (
	"github.com/AmitSuresh/orderapi/src/infra/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

func NewConsumer(cfg *config.Config, l *zap.Logger) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaServers,
		"group.id":          cfg.KafkaConsumerGroup,
		"auto.offset.reset": cfg.KafkaOffset,
	})

	if err != nil {
		l.Error(err.Error())
		return nil, err
	}

	return consumer, nil
}

func NewProducer(cfg *config.Config, l *zap.Logger) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaServers,
		"acks":              cfg.KafkaAcks,
	})

	if err != nil {
		l.Error(err.Error())
		return nil, err
	}

	return producer, nil
}
