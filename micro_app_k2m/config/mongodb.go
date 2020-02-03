package config

import (
	"context"
	"fmt"
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

type MongoDB struct {
	Host string
	Port string
}

func ConnectMongodb() *mongo.Client {
	clientOptions := options.Client().ApplyURI(GetMongodbUri())
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
		}
	})
	return mongodbClient
}

func GetMongodbUri() string {
	return fmt.Sprintf("mongodb://%s:%s", configuration.MongoDB.Host, configuration.MongoDB.Port)
}

func init() {
	GetMongodbClient()
}
