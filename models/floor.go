package models

import "gorm.io/gorm"

// Floor represents a floor in a building
type Floor struct {
	// Primary key with auto increment
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	// Foreign key to Building
	BuildingID uint `gorm:"not null;index" json:"building_id" binding:"required"`

	// Floor number
	Number int `gorm:"not null" json:"number" binding:"required"`

	// Floor name/description
	Name string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`

	// Relationship: Floor belongs to Building
	Building Building `gorm:"foreignKey:BuildingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"building,omitempty"`

	// Relationship: Floor has many Rooms
	Rooms []Room `gorm:"foreignKey:FloorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"rooms,omitempty"`

	// Timestamps
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for Floor model
func (Floor) TableName() string {
	return "floors"
}
