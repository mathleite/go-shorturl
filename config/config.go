package config

import (
	"mathleite/short-url/pkg/redis"
	"os"
)

var (
	redisClient *redis.RedisClient
	baseUrl     string
)

func GetRedisClient() *redis.RedisClient {
	redisClient = redis.NewRedis()
	return redisClient
}

func GetBaseUrl() string {
	baseUrl = os.Getenv("BASE_URL")
	return baseUrl
}
