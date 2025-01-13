package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDataOutput struct {
	Identifier  string `redis:"identifier"`
	OriginalUrl string `redis:"original_url"`
	BaseUrl     string `redis:"base_url"`
}

type RedisDataInput struct {
	Identifier  string        `redis:"identifier"`
	OriginalUrl string        `redis:"original_url"`
	BaseUrl     string        `redis:"base_url"`
	Ttl         time.Duration `redis:"ttl"`
}

type RedisClient struct {
	redis *redis.Client
}
