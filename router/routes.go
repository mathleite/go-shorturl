package router

import (
	"mathleite/short-url/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	v1.POST("/", handler.CreateShortUrl)
	v1.GET("/:identifier", handler.RedirectOriginalUrl)
}
