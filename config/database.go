package config

import (
	"fmt"
	"incident-report/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database holds the GORM database instance
var DB *gorm.DB

// InitDatabase initializes the MySQL database connection using GORM
// It reads configuration from environment variables and establishes connection
func InitDatabase() error {
	// Get environment variables for database configuration
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Build DSN (Data Source Name) for MySQL connection
	// Format: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	// Open database connection with MySQL dialect
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	log.Println("Database connection established successfully")

	// Run auto-migration for all models
	// This will create tables if they don't exist
	if err := AutoMigrate(); err != nil {
		log.Printf("Auto migration failed: %v", err)
		return err
	}

	return nil
}

// AutoMigrate runs migrations for all models
// Each model's schema will be created/updated in the database
func AutoMigrate() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	// Migrate the User model - this creates the users table with all columns
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	// Migrate building management models
	if err := DB.AutoMigrate(&models.Building{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&models.Floor{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&models.Room{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&models.ComponentCategory{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&models.Component{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&models.Report{}); err != nil {
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}

// CloseDatabase closes the database connection
func CloseDatabase() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
