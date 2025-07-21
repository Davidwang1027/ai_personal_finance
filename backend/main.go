package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	// Import the config package directly
	"github.com/davidwang/go-finance-api/go-finance-api/config"
	"github.com/davidwang/go-finance-api/go-finance-api/db"

	// Import our new plaid package
	"github.com/davidwang/go-finance-api/go-finance-api/handlers"
	"github.com/davidwang/go-finance-api/go-finance-api/plaid"
)

func main() {
	// Load configuration
	log.Println("Loading configuration...")
	cfg := config.Load()
	log.Println("Configuration loaded successfully")

	// Check if we should skip database initialization (for development)
	skipDB := os.Getenv("SKIP_DB") == "true"
	var database *db.Database

	if !skipDB {
		// Initialize database
		log.Println("Connecting to database...")
		var err error
		database, err = db.NewDatabase(cfg.DB)
		if err != nil {
			log.Printf("Failed to connect to database: %v", err)
			log.Println("To skip database initialization, set SKIP_DB=true")
			log.Fatal("Exiting due to database connection failure")
		}
		defer database.Close()

		// Run database migrations
		log.Println("Running database migrations...")
		if err := database.Setup(); err != nil {
			log.Fatalf("Failed to set up database: %v", err)
		}
		log.Println("Database migrations completed successfully")
	} else {
		log.Println("Skipping database initialization (SKIP_DB=true)")
	}

	// Initialize Plaid client
	log.Println("Initializing Plaid client...")
	plaidClient := plaid.NewClient(cfg)
	log.Println("Plaid client initialized successfully")

	// Initialize handlers
	plaidHandler := handlers.NewPlaidHandler(plaidClient)

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
				"status":      "operational",
				"env":         cfg.PlaidEnv,
				"plaid_ready": plaidClient != nil,
				"db_ready":    skipDB || (database != nil && database.Ping() == nil),
			})
		})

		// Plaid endpoints
		plaidRoutes := api.Group("/plaid")
		{
			// Link and access token endpoints
			plaidRoutes.POST("/create_link_token", plaidHandler.CreateLinkToken)
			plaidRoutes.POST("/exchange_public_token", plaidHandler.ExchangePublicToken)

			// Account and transaction endpoints
			plaidRoutes.GET("/accounts", plaidHandler.GetAccounts)
			plaidRoutes.POST("/transactions", plaidHandler.GetTransactions)
			plaidRoutes.POST("/transactions/sync", plaidHandler.SyncTransactions)

			// Item management endpoints
			plaidRoutes.GET("/item", plaidHandler.GetItem)
			plaidRoutes.POST("/item/webhook", plaidHandler.UpdateItemWebhook)

			// Webhook endpoint
			plaidRoutes.POST("/webhook", plaidHandler.HandleWebhook)
		}
	}

	// Start the server
	port := ":" + cfg.Port
	log.Printf("Server starting on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
