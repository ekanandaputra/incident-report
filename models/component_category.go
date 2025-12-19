package models

import "gorm.io/gorm"

// ComponentCategory represents a category of components
type ComponentCategory struct {
	// Primary key with auto increment
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	// Category code - unique identifier
	Code string `gorm:"type:varchar(100);uniqueIndex;not null" json:"code" binding:"required"`

	// Category name
	Name string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`

	// Category description (optional)
	Description string `gorm:"type:text" json:"description"`

	// Relationship: ComponentCategory has many Components
	Components []Component `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"components,omitempty"`

	// Timestamps
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for ComponentCategory model
func (ComponentCategory) TableName() string {
	return "component_categories"
}
