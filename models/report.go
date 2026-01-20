package models

import (
	"time"
)

// ReportStatus represents the status of a report
type ReportStatus string

const (
	ReportStatusPending    ReportStatus = "PENDING"
	ReportStatusInProgress ReportStatus = "IN_PROGRESS"
	ReportStatusCompleted  ReportStatus = "COMPLETED"
)

// Report represents the Report entity in the database
type Report struct {
	// Primary key with auto increment
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	// Report name
	Name string `gorm:"type:text;not null" json:"name" binding:"required"`

	// Foreign key to Room
	RoomID uint `gorm:"not null;index" json:"room_id" binding:"required"`

	// Foreign key to User (nullable, assigned later by admin)
	UserID *uint `gorm:"index" json:"user_id,omitempty"`

	// Foreign key to Component
	ComponentID uint `gorm:"not null;index" json:"component_id" binding:"required"`

	// Report status
	Status ReportStatus `gorm:"type:varchar(20);not null;default:'PENDING'" json:"status" binding:"required,oneof=PENDING IN_PROGRESS COMPLETED"`

	// Timestamps for tracking report creation and updates
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Room      Room      `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room,omitempty"`
	User      *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Component Component `gorm:"foreignKey:ComponentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"component,omitempty"`
}

// TableName specifies the table name for the Report model
func (Report) TableName() string {
	return "reports"
}
