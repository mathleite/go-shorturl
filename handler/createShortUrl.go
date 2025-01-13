package handler

import (
	"fmt"
	"mathleite/short-url/internal/identifier"
	"mathleite/short-url/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateShortUrl(ctx *gin.Context) {
	request := CreateShortUrlRequest{}
	ctx.BindJSON(&request)

	if error := request.Validate(); error != nil {
		ctx.JSON(400, error)
		return
	}
	identifier := identifier.CreateIdentifier(request.Url)
	cachedData := redisClient.Get(identifier)
	if cachedData.Identifier != identifier {
		input := redis.RedisDataInput{
			Identifier:  identifier,
			OriginalUrl: request.Url,
			BaseUrl:     baseUrl,
			Ttl:         300 * time.Second,
		}
		redisClient.Store(input)
	}

	sendSuccessResponse(ctx, CreateShortUrlResponse{ShortUrl: fmt.Sprintf("%s/%s", baseUrl, identifier)})
}
