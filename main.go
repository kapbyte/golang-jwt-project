package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	routes "github.com/kapbyte/golang-jwt-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// router.GET("/api-1", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong-api-1",
	// 	})
	// })

	router.Run(":" + port)
}
