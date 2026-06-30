package jobs

import (
	"context"
	"net/http"
	"time"

	"job-monitoring-platform/api/internal/database"

	"github.com/gin-gonic/gin"
)

func GetJob() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		id := c.Param("id")
		var job Job

		if err := database.DB.WithContext(ctx).First(&job, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
			return
		}

		c.JSON(http.StatusOK, job)
	}
}

func CreateJob() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		var job Job

		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": "Invalid input"})
			return
		}

		if err := database.DB.WithContext(ctx).Create(&job).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
			return
		}

		c.JSON(http.StatusCreated, job)
	}
}
