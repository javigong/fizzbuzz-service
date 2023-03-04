package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Gin
	router := gin.Default()

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://fizzbuzz-service-client-javier-gongora.vercel.app"},
    AllowMethods:     []string{"POST"},
    AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "https://github.com"
    },
    MaxAge: 12 * time.Hour,
  }))

	// Define route
	router.POST("/fizzbuzz", getFizzbuzzMessage)
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := router.Run(":" + port); err != nil {
        log.Panicf("error: %s", err)
	}
}

// Handle get fizzbuzz message
func getFizzbuzzMessage(c *gin.Context) {
	// Get environment variables
	fizzMessage := os.Getenv("FIZZ_MSG")
	buzzMessage := os.Getenv("BUZZ_MSG")
	fizzBuzzMessage := os.Getenv("FIZZBUZZ_MSG")

	// Read input count number from JSON request body
	var requestBody struct {
		Count int `json:"count"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	// Calculate fizzbuzz message
	number := requestBody.Count
	var result string

	if number%3 == 0 && number%5 == 0 {
		result = fizzBuzzMessage
	} else if number%3 == 0 {
		result = fizzMessage
	} else if number%5 == 0 {
		result = buzzMessage
	} else {
		result = ""
	}

	// Return fizzbuzz message
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
	})
}
