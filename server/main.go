package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Get the host and port env variables
	url := os.Getenv("URL")
	if url == "" {
		url = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize Gin
	router := gin.Default()

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"POST"},
    AllowHeaders:     []string{"Origin"},
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
	addr := fmt.Sprintf("%s:%s", url, "3000")
	router.Run(addr)
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
