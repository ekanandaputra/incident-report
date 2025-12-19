package models

import "gorm.io/gorm"

// Room represents a room on a floor
type Room struct {
	// Primary key with auto increment
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	// Foreign key to Floor
	FloorID uint `gorm:"not null;index" json:"floor_id" binding:"required"`

	// Room code - unique identifier
	Code string `gorm:"type:varchar(100);uniqueIndex;not null" json:"code" binding:"required"`

	// Room name/description
	Name string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`

	// Relationship: Room belongs to Floor
	Floor Floor `gorm:"foreignKey:FloorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"floor,omitempty"`

	// Relationship: Room has many Components
	Components []Component `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"components,omitempty"`

	// Timestamps
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for Room model
func (Room) TableName() string {
	return "rooms"
}
