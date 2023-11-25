package api

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.RouterGroup) {
	r.POST("/sig-up")
	r.POST("/sig-in")
	r.PUT("/user/:id")
	r.GET("/user/:id")
}
