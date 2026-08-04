package main

import (
	"job-monitoring-platform/api/internal/database"
	"job-monitoring-platform/api/internal/jobs"
	"job-monitoring-platform/api/internal/messaging"
	"job-monitoring-platform/api/internal/middleware"
	"job-monitoring-platform/api/internal/redis"
	"job-monitoring-platform/api/internal/routes"
	"job-monitoring-platform/api/internal/users"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found. Using exisisting enviroment variables")
	}
	database.Connect()
	redis.Connect()
	rabbit, err := messaging.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()

	if err := rabbit.SetupInfrastructure(); err != nil {
		log.Fatal(err)
	}
	database.DB.AutoMigrate(&jobs.Job{}, &users.User{})

	router := gin.Default()
	router.Use(middleware.SessionMiddleware())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.SetupProtectedRoutes(router)
	routes.SetupUnprotectedRoutes(router)
	router.Run()
}
