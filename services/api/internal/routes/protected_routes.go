package routes

import (
	"job-monitoring-platform/api/internal/jobs"
	"job-monitoring-platform/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Group("/")
	router.Use(middleware.RequireAuth())
	router.POST("/jobs", jobs.CreateJob())
	router.GET("/job/:id", jobs.GetJob())
	router.DELETE("/job/:id", jobs.DeleteJob())
}
