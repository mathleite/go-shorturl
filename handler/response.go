package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

func sendSuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"time": time.Now().Unix(),
		"data": data,
	})
}

func sendNotFoundResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"time":    time.Now().Unix(),
		"message": "URL not found",
	})
}
