package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	MongoDB MongoDB
}

var configuration Configuration

func LoadConfig() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

func init() {
	LoadConfig()
}