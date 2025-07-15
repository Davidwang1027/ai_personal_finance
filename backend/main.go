package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// Import the config package directly
	"github.com/davidwang/go-finance-api/go-finance-api/config"
)

func main() {
	// Load configuration
	log.Println("Loading configuration...")
	cfg := config.Load()
	log.Println("Configuration loaded successfully")

	// Set up Gin router
	log.Println("Setting up Gin router...")
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Finance API is running",
		})
	})

	// API routes group
	api := router.Group("/api")
	{
		// Add a simple test endpoint
		api.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "operational",
				"env":    cfg.PlaidEnv,
			})
		})
	}

	// Start the server
	port := ":" + cfg.Port
	log.Printf("Server starting on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
