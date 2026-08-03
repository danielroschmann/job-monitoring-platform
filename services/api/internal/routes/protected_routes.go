package routes

import (
	"job-monitoring-platform/api/internal/jobs"
	"job-monitoring-platform/api/internal/middleware"
	"job-monitoring-platform/api/internal/users"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	protected := router.Group("/")
	protected.Use(middleware.RequireAuth())
	protected.Use(middleware.RequireAuth())
	protected.POST("/jobs", jobs.CreateJob())
	protected.GET("/job/:id", jobs.GetJob())
	protected.DELETE("/job/:id", jobs.DeleteJob())
	protected.POST("/logout", users.LogoutUser())
	protected.GET("/me", func(c *gin.Context) {
		userId := c.MustGet("userId")
		c.JSON(200, gin.H{"message": "Authenticated", "userId": userId})
	})
}
