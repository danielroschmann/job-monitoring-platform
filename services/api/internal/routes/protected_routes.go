package routes

import (
	"job-monitoring-platform/api/internal/jobs"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.POST("/jobs", jobs.CreateJob())
	router.GET("/job/:id", jobs.GetJob())
}
