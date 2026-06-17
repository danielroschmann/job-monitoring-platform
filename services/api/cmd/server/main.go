package main

import (
	"github.com/gin-gonic/gin"
	"job-monitoring-platform/api/internal/database"
)

func main() {
	database.Connect()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
