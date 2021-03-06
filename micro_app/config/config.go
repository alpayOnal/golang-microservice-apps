package config

import (


	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Kafka   Kafka
	MongoDB MongoDB
	Redis   Redis
}

var (
	configuration Configuration
)

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

func GetConfiguration() *Configuration {
	return &configuration
}

func init() {
	LoadConfig()
}
