package main

import (
	"job-monitoring-platform/api/internal/database"
	"job-monitoring-platform/api/internal/jobs"
	"job-monitoring-platform/api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found. Using exisisting enviroment variables")
	}
	database.Connect()
	database.DB.AutoMigrate(&jobs.Job{})

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.SetupProtectedRoutes(router)
	router.Run()
}
