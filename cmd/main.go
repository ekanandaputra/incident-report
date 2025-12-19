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
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}

	// Log startup message
	log.Printf("🚀 Server starting on http://%s:%s", host, port)
	log.Println("📝 API Documentation:")
	log.Printf("   🔗 Swagger UI: http://localhost:%s/swagger/index.html", port)
	log.Println("   POST   /api/v1/users           - Create a new user")
	log.Println("   GET    /api/v1/users           - Get all users (with pagination)")
	log.Println("   GET    /api/v1/users/:id       - Get a specific user")
	log.Println("   PUT    /api/v1/users/:id       - Update a specific user")
	log.Println("   DELETE /api/v1/users/:id       - Delete a specific user")
	log.Println("   GET    /api/v1/health         - Health check")

	// Start the server
	// If TLS cert and key are provided via env vars or present as cert.pem/key.pem,
	// run the server with TLS. Otherwise fall back to plain HTTP.
	certFile := os.Getenv("TLS_CERT_FILE")
	keyFile := os.Getenv("TLS_KEY_FILE")

	// If env vars not set, look for cert.pem/key.pem in project root
	if certFile == "" || keyFile == "" {
		if _, err := os.Stat("cert.pem"); err == nil {
			certFile = "cert.pem"
		}
		if _, err := os.Stat("key.pem"); err == nil {
			keyFile = "key.pem"
		}
	}

	addr := host + ":" + port
	if certFile != "" && keyFile != "" {
		log.Printf("🔒 Starting HTTPS server on https://%s", addr)
		if err := router.RunTLS(addr, certFile, keyFile); err != nil {
			log.Fatalf("Failed to start HTTPS server: %v", err)
		}
	} else {
		log.Printf("Starting HTTP server on http://%s", addr)
		if err := router.Run(addr); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}
}
