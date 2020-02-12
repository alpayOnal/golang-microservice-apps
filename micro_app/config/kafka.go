package config

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	Host            string
	Port            string
	GroupId         string
	AutoOffsetReset string
}

func GetKafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{"bootstrap.servers": GetKafkaHost()}
}

func GetKafkaHost() string {
	return fmt.Sprintf("%s:%s", configuration.Kafka.Host, configuration.Kafka.Port)
}
