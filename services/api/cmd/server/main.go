package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"job-monitoring-platform/api/internal/database"
	"job-monitoring-platform/api/internal/jobs"
)

func main() {
	godotenv.Load()
	database.Connect()
	database.DB.AutoMigrate(&jobs.Job{})

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
