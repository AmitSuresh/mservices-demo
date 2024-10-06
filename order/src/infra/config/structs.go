package config

type Config struct {
	ServerAddr         string
	GrpcAddr           string
	KafkaServers       string
	KafkaOffset        string
	KafkaConsumerGroup string
	KafkaAcks          string
}
