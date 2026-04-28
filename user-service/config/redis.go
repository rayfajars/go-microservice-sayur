package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("[NewRedisClient-1]: %v", err)
	}

	log.Println("[NewRedisClient-2]: Redis connected")

	return client
}
