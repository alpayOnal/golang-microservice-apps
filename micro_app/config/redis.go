package config

import (
	"sync"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	redisClient *redis.Client
	onceRedisClient sync.Once
)

type Redis struct {
	Url string
}

func GetRedisClient() *redis.Client {
	onceRedisClient.Do(func () {
		log.Info("connect to the Redis ")

		opts, err := redis.ParseURL(GetRedisUrl())
		if err != nil {
			log.Fatal("Couldn't connect to the Redis ", err)
		}
		redisClient := redis.NewClient(opts)
		_, err = redisClient.Ping().Result()
		if err != nil {
			log.Fatal("Couldn't connect to the Redis ", err)
		}
	})
	return redisClient
}

func GetRedisUrl() string {
	return configuration.Redis.Url
}

func init() {
	GetRedisClient()
}
