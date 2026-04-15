package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	return rdb
}
