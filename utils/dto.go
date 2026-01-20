package utils

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=255"`
	Email string `json:"email" binding:"required,email"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,min=2,max=255"`
	Email string `json:"email" binding:"omitempty,email"`
}

// UserResponse represents the response payload for a user
type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// PaginationQuery represents pagination parameters
type PaginationQuery struct {
	Page     int `form:"page" binding:"omitempty,min=1"`
	PageSize int `form:"page_size" binding:"omitempty,min=1,max=100"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse struct {
	Data      interface{} `json:"data"`
	Page      int         `json:"page"`
	PageSize  int         `json:"page_size"`
	Total     int64       `json:"total"`
	TotalPage int         `json:"total_page"`
}

// CreateReportRequest represents the request payload for creating a report
type CreateReportRequest struct {
	Name        string `json:"name" binding:"required"`
	RoomID      uint   `json:"room_id" binding:"required"`
	UserID      *uint  `json:"user_id,omitempty"`
	ComponentID uint   `json:"component_id" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=PENDING IN_PROGRESS COMPLETED"`
}

// UpdateReportRequest represents the request payload for updating a report
type UpdateReportRequest struct {
	Name        string `json:"name" binding:"omitempty"`
	RoomID      uint   `json:"room_id" binding:"omitempty"`
	UserID      *uint  `json:"user_id" binding:"omitempty"`
	ComponentID uint   `json:"component_id" binding:"omitempty"`
	Status      string `json:"status" binding:"omitempty,oneof=PENDING IN_PROGRESS COMPLETED"`
}

// ReportResponse represents the response payload for a report
type ReportResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	RoomID      uint   `json:"room_id"`
	UserID      *uint  `json:"user_id,omitempty"`
	ComponentID uint   `json:"component_id"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// AssignUserRequest represents the request payload for assigning a user to a report
type AssignUserRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}
