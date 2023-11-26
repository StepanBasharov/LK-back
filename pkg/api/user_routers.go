package api

import (
	"LK_back/pkg/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouterUsers(r *gin.Engine) {
	r.POST("/sig-up", handlers.CreateUserHandler)
	r.POST("/sig-in", handlers.Authorization)
	r.PUT("/user/:id")
	r.GET("/user/:id")
}
