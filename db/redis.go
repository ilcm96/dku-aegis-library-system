package db

import "github.com/redis/go-redis/v9"

var redisClient *redis.Client

func InitRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "dku-redis:6379",
		DB:   0,
	})
}

func RedisClient() *redis.Client {
	return redisClient
}
