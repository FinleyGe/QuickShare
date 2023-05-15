package redis

import (
	"github.com/redis/go-redis/v9"
)

// Redis Client
var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
