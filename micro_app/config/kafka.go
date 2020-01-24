package config

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func GetKafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", KafkaHost, KafkaPort)}
}

