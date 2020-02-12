package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/alpayOnal/golang-microservice-apps/micro_app/repository/kafka"
	"micro_apps/micro_app/repository/mongodb"
)

func main() {
	receiveFromKafka()
}

func receiveFromKafka() {

	fmt.Println("Start receiving from Kafka")
	consumer, err := kafka.NewItemRepository("items-topic1").GetConsumer()
	if err != nil {
		log.Error("Kafka consumer error:", err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			message := string(msg.Value)
			mongodb.GetItemRepository().Store(message)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	consumer.Close()

}
