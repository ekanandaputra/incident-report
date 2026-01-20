package main

import (
	"fmt"
	"incident-report/config"
	"incident-report/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	if err := config.InitDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Drop and recreate report table
	if err := config.DB.Migrator().DropTable(&models.Report{}); err != nil {
		log.Fatal("Failed to drop report table:", err)
	}
	fmt.Println("✓ Dropped report table")

	// Recreate with new schema
	if err := config.DB.AutoMigrate(&models.Report{}); err != nil {
		log.Fatal("Failed to migrate report table:", err)
	}
	fmt.Println("✓ Recreated report table with nullable user_id")

	fmt.Println("\nDatabase reset complete!")
}
