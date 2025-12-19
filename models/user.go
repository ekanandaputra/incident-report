package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the User entity in the database
type User struct {
	// Primary key with auto increment
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	// User's full name
	Name string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`

	// User's email - unique constraint
	Email string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" binding:"required,email"`

	// Timestamps for tracking user creation and updates
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // Soft delete
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
