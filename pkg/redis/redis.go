package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx context.Context = context.Background()

func NewRedis() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	return &RedisClient{redis: client}
}

func (c *RedisClient) Store(data RedisDataInput) {
	hashFields := []string{
		"identifier", data.Identifier,
		"base_url", data.BaseUrl,
		"original_url", data.OriginalUrl,
	}
	_, err := c.redis.HSet(ctx, data.Identifier, hashFields).Result()
	if err != nil {
		panic(err)
	}
	_, err = c.redis.Expire(ctx, data.Identifier, data.Ttl).Result()
	if err != nil {
		panic(err)
	}
}

func (c *RedisClient) Get(key string) RedisDataOutput {
	var data RedisDataOutput
	err := c.redis.HGetAll(ctx, key).Scan(&data)
	if err != nil {
		panic(err)
	}
	return data
}
