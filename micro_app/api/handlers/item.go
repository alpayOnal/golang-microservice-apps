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
	"micro_apps/micro_app/models"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ItemHandler  represent the httphandler for item
type ItemHandler struct {
}

func NewItemHandler(e *echo.Echo) {
	handler := &ItemHandler{}

	e.GET("/items/:id", handler.GetItem)
	e.GET("/items", handler.GetItems)

	e.POST("/items", handler.AddItem)
}

func (i *ItemHandler) AddItem(c echo.Context) error {
	item := models.NewItem()
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&item)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	SaveItemToKafka(item)
	log.Printf("this is your item %#v", item)
	return c.String(http.StatusOK, "We got your Item!!!")
}

func (i *ItemHandler) GetItem(c echo.Context) error {
	id := c.Param("id")

	mc := config.GetMongodbClient()
	collection := mc.Database("testing").Collection("items")

	var item models.Item
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

func (i *ItemHandler) GetItems(c echo.Context) error {
	mc := config.GetMongodbClient()
	collection := mc.Database("testing").Collection("items")

	//var item models.Item
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	var itemList []models.Item
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		var item models.Item
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

func SaveItemToKafka(item models.Item) {

	jsonString, _ := json.Marshal(item)

	itemString := string(jsonString)
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
