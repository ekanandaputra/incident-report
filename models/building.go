package models

import "gorm.io/gorm"

// Building represents a physical building
type Building struct {
	// Primary key with auto increment
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	// Building code - unique identifier
	Code string `gorm:"type:varchar(100);uniqueIndex;not null" json:"code" binding:"required"`

	// Building name
	Name string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`

	// Building location (optional)
	Location string `gorm:"type:varchar(500)" json:"location"`

	// Relationship: Building has many Floors
	Floors []Floor `gorm:"foreignKey:BuildingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"floors,omitempty"`

	// Timestamps
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for Building model
func (Building) TableName() string {
	return "buildings"
}
