package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectOriginalUrl(ctx *gin.Context) {
	identifier := ctx.Param("identifier")
	cachedData := redisClient.Get(identifier)
	if cachedData.Identifier != identifier {
		sendNotFoundResponse(ctx)
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, cachedData.OriginalUrl)
}
