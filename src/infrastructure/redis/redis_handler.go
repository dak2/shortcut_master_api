package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type RedisHandler struct {
	client *redis.Client
}

func NewRedisHandler() *RedisHandler {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	return &RedisHandler{client: client}
}

func (handler *RedisHandler) SET(key string, value string) error {
	ctx := context.Background()
	return handler.client.Set(ctx, key, value, 0).Err()
}

func (handler *RedisHandler) GET(key string) (string, error) {
	ctx := context.Background()
	return handler.client.Get(ctx, key).Result()
}

func (handler *RedisHandler) DEL(key string) error {
	ctx := context.Background()
	return handler.client.Del(ctx, key).Err()
}
