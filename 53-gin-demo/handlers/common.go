package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ping": "pong"})
}

func Health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
