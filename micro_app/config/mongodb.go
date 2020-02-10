package config

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	log "github.com/sirupsen/logrus"
)

var (
	mongodbClient *mongo.Client
	once          sync.Once
)

type MongoDB struct {
	Host string
	Port string
}

func newMongoClient() *mongo.Client {
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
		mongodbClient = newMongoClient()
		err := mongodbClient.Ping(context.Background(), readpref.Primary())
		if err != nil {
			log.Fatal("Couldn't connect to the Mongodb ", err)
		} else {
			log.Info("Connected to the Mongodb ")
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
