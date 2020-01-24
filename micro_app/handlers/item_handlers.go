package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo"

	"micro_apps/micro_app/config"
	"micro_apps/micro_app/types"
)

func AddItem(c echo.Context) error {
	item := types.Item{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&item)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	saveItemToKafka(item)
	log.Printf("this is your item %#v", item)
	return c.String(http.StatusOK, "We got your Item!!!")
}

func saveItemToKafka(item types.Item) {

	jsonString, _ := json.Marshal(item)

	itemString := string(jsonString)
	fmt.Print(config.GetKafkaConfig())
	p, err := kafka.NewProducer(config.GetKafkaConfig())
	if err != nil {
		panic(err)
	}
	// Produce messages to topic (asynchronously)
	topic := "items-topic1"
	for _, word := range []string{string(itemString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
}

