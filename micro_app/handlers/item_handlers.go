package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"micro_apps/micro_app/config"
	"micro_apps/micro_app/types"
)

func AddItem(c echo.Context) error {
	item := types.NewItem()
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

func GetItem(c echo.Context) error {
	id := c.Param("id")

	mc := config.GetMongodbClient()
	collection := mc.Database("testing").Collection("items")

	var item types.Item
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objID}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		fmt.Println(err)
	}
	item.Id = id
	return c.JSON(http.StatusOK, item)
}

func GetItems(c echo.Context) error {
	mc := config.GetMongodbClient()
	collection := mc.Database("testing").Collection("items")

	//var item types.Item
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	var itemList []types.Item
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		var item types.Item
		record, _ := json.Marshal(result)
		err = json.Unmarshal(record, &item)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		itemList = append(itemList, item)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	//resp := json.Marshal(itemList)
	return c.JSON(http.StatusOK, itemList)
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
