package users

import (
	"context"
	"job-monitoring-platform/api/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid user input"})
			return
		}

		if err := database.DB.WithContext(ctx).Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		id := c.Param("id")
		var user User

		if err := database.DB.WithContext(ctx).First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
