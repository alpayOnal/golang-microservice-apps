package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	KafkaHost string
	KafkaPort string
)

func ReadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	KafkaHost = os.Getenv("KAFKA_HOST")
	KafkaPort = os.Getenv("KAFKA_PORT")
}

func Load() {
	ReadConfig()
}
