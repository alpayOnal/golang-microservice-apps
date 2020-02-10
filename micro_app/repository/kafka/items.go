package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"micro_apps/micro_app/config"
	"micro_apps/micro_app/models"
)

var itemRepository ItemRepository

type ItemRepository struct {
	config *kafka.ConfigMap
	topic  string
}

func init() {
	itemRepository = ItemRepository{
		config.GetKafkaConfig(),
		"items-topic1",
	}
}

func GetItemRepository() *ItemRepository {
	return &itemRepository
}

func (i *ItemRepository) Add(item models.Item) error {

	jsonString, _ := json.Marshal(item)
	itemString := string(jsonString)
	p, err := kafka.NewProducer(i.config)
	if err != nil {
		return err
	}
	// Produce messages to topic (asynchronously)
	topic := i.topic
	for _, word := range []string{string(itemString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	return nil
}
