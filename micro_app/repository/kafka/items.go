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

func NewItemRepository(topic string) *ItemRepository {
	var itemRepository ItemRepository
	itemRepository.topic = topic
	itemRepository = ItemRepository{
		config.GetKafkaConfig(),
		itemRepository.topic,
	}
	return &itemRepository
}

func (r *ItemRepository) GetConsumer() (*kafka.Consumer, error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.GetKafkaHost(),
		"group.id":          config.GetConfiguration().Kafka.GroupId,
		"auto.offset.reset": config.GetConfiguration().Kafka.AutoOffsetReset,
	})

	if err != nil {
		return nil, err
	}
	c.SubscribeTopics([]string{r.topic}, nil)
	return c, nil
}

func (r *ItemRepository) Store(item models.Item) error {

	jsonString, _ := json.Marshal(item)
	itemString := string(jsonString)
	p, err := kafka.NewProducer(r.config)
	if err != nil {
		return err
	}
	topic := r.topic
	for _, word := range []string{string(itemString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	return nil
}
