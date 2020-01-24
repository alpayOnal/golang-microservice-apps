package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"


	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"micro_apps/micro_app_k2m/config"
	"micro_apps/micro_app_k2m/types"
)

func main() {

	//Create MongoDB session

	receiveFromKafka()

}

func receiveFromKafka() {

	fmt.Println("Start receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"items-topic1"}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			job := string(msg.Value)
			saveItemToMongo(job)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()

}

func saveItemToMongo(jobString string) {

	c := config.GetMongodbClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	fmt.Println(jobString)
	fmt.Println("Save to MongoDB")

	collection := c.Database("testing").Collection("items")

	//Save data into Job struct
	var item types.Item
	b := []byte(jobString)
	err = json.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}

	//Insert item into MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, item)
	if res != nil {
		fmt.Println("ItemId : %s", res.InsertedID)
	}

	fmt.Println("Saved to MongoDB : %s", jobString)
}
