package main

import (
	"incident-report/config"
	"incident-report/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main is the entry point of the application
func main() {
	// Load environment variables from .env file
	// godotenv allows us to use a .env file for configuration management
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize database connection
	// This establishes MySQL connection using GORM and runs auto-migration
	if err := config.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Ensure database connection is closed when application exits
	defer func() {
		if err := config.CloseDatabase(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	// Set Gin mode based on environment
	// Use "debug" for development, "release" for production
	environment := os.Getenv("ENVIRONMENT")
	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a new Gin router instance
	router := gin.Default()

	// Register all API routes
	routes.RegisterRoutes(router)

	// Get server configuration from environment variables
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	// Set defaults if environment variables are not set
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8080"
	}

	// Log startup message
	log.Printf("🚀 Server starting on http://%s:%s", host, port)
	log.Println("📝 API Documentation:")
	log.Println("   POST   /api/v1/users           - Create a new user")
	log.Println("   GET    /api/v1/users           - Get all users (with pagination)")
	log.Println("   GET    /api/v1/users/:id       - Get a specific user")
	log.Println("   PUT    /api/v1/users/:id       - Update a specific user")
	log.Println("   DELETE /api/v1/users/:id       - Delete a specific user")
	log.Println("   GET    /api/v1/health         - Health check")

	// Start the server
	// ListenAndServe blocks until the server is stopped or encounters an error
	if err := router.Run(host + ":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
