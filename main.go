package main

import (
	"LK_back/pkg/api"
	"LK_back/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := settings.InitConfigService()
	bind := fmt.Sprintf("%s:%s", config.Host, config.Port)
	service := gin.Default()
	service.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Service Work": "OK"})
	})
	api.SetupRouterUsers(service)
	service.Run(bind)
}
