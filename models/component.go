package models

import "gorm.io/gorm"

// Component represents a physical component in a room
type Component struct {
	// Primary key with auto increment
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	// Foreign key to Room
	RoomID uint `gorm:"not null;index" json:"room_id" binding:"required"`

	// Foreign key to ComponentCategory
	CategoryID uint `gorm:"not null;index" json:"category_id" binding:"required"`

	// Component code - unique identifier
	Code string `gorm:"type:varchar(100);uniqueIndex;not null" json:"code" binding:"required"`

	// Component name
	Name string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`

	// Brand/manufacturer (optional)
	Brand string `gorm:"type:varchar(255)" json:"brand"`

	// Specification/model details (optional)
	Specification string `gorm:"type:text" json:"specification"`

	// Year of procurement (optional)
	ProcurementYear int `json:"procurement_year"`

	// Relationship: Component belongs to Room
	Room Room `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room,omitempty"`

	// Relationship: Component belongs to ComponentCategory
	Category ComponentCategory `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category,omitempty"`

	// Timestamps
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for Component model
func (Component) TableName() string {
	return "components"
}
