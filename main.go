package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	service := gin.Default()
	service.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Service Work": "OK"})
	})
	service.Run("0.0.0.0:8080")
}
