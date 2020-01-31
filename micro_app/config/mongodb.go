package config

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongodbClient *mongo.Client
	once          sync.Once
)

func ConnectMongodb() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	mongodbClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = mongodbClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return mongodbClient
}

func GetMongodbClient() *mongo.Client {
	once.Do(func() {
		mongodbClient = ConnectMongodb()
		err := mongodbClient.Ping(context.Background(), readpref.Primary())
		if err != nil {
			log.Fatal("Couldn't connect to the Mongodb ", err)
		} else {
			log.Println("Mongodb Connected!")
		}
	})
	return mongodbClient
}
func init() {
	GetMongodbClient()
}
