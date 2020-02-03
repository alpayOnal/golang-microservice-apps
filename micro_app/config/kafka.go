package config

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	Host string
	Port string
}

func GetKafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", configuration.Kafka.Host, configuration.Kafka.Port)}
}
