package config

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	host := os.Getenv("REDISHOST")
	password := os.Getenv("REDISPASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}
