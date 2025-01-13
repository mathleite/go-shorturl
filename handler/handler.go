package handler

import (
	"mathleite/short-url/config"
	"mathleite/short-url/pkg/redis"
)

var (
	redisClient redis.RedisClient
	baseUrl     string
)

func InitializeHandler() {
	redisClient = *config.GetRedisClient()
	baseUrl = config.GetBaseUrl()
}
