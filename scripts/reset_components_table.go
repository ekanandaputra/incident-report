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
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize database
	if err := config.InitDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Drop and recreate components table
	if err := config.DB.Migrator().DropTable(&models.Component{}); err != nil {
		log.Fatal("Failed to drop components table:", err)
	}
	fmt.Println("✓ Dropped components table")

	// Recreate with new schema
	if err := config.DB.AutoMigrate(&models.Component{}); err != nil {
		log.Fatal("Failed to migrate components table:", err)
	}
	fmt.Println("✓ Recreated components table with nullable room_id")

	fmt.Println("\nDatabase reset complete!")
}
