package redis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"micro_apps/micro_app/config"
)

type itemRepository struct {
	client *redis.Client
}

var repository itemRepository

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func init() {
	client, err := newRedisClient(config.GetRedisUrl())
	if err != nil {
		log.Error(errors.Wrap(err, "repository.NewItemRepository"))
	}
	repository.client = client
}

func GetItemRepository() *itemRepository {
	return &repository
}

func (r *itemRepository) Find() error {
	return nil
}

func (r *itemRepository) Store() error {
	return nil
}
