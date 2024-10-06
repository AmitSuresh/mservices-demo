package queue

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

var MsgStr string = "Shipping_Info"

func ProduceMsg(p *kafka.Producer, m *kafka.Message, e chan kafka.Event, l *zap.Logger) error {

	err := p.Produce(m, nil)
	if err != nil {
		l.Error("error producing", zap.Error(err))
		return err
	}

	return nil
}

func NewMsg(b []byte, t *kafka.TopicPartition) *kafka.Message {
	return &kafka.Message{
		TopicPartition: *t,
		Value:          b,
	}
}

func NewPartition(s *string, n int32) *kafka.TopicPartition {
	return &kafka.TopicPartition{
		Topic:     s,
		Partition: kafka.PartitionAny,
	}
}
