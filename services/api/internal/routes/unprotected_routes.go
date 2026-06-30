package routes

import (
	"job-monitoring-platform/api/internal/users"

	"github.com/gin-gonic/gin"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	router.POST("/users", users.CreateUser())
	router.GET("/user/:id", users.GetUser())
}
