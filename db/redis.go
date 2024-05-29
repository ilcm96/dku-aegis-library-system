package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var redisClient *redis.Client

func InitRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		DB:   0,
	})
}

func RedisClient() *redis.Client {
	return redisClient
}
