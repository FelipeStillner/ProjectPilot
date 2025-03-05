package adapter

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis() *Redis {
	addr := os.Getenv("PORT_REDIS")
	db := 0
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})

	return &Redis{client: rdb}
}

func (r *Redis) Set(key string, value string, time time.Duration) error {
	ctx := context.Background()
	err := r.client.Set(ctx, key, value, time).Err()
	return err
}

func (r *Redis) Get(key string) (*string, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err != redis.Nil {
		return nil, err
	}
	return &val, nil
}
