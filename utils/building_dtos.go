package utils

// ===== Building DTOs =====

// CreateBuildingRequest represents building creation request
type CreateBuildingRequest struct {
	Code     string `json:"code" binding:"required,min=1,max=100"`
	Name     string `json:"name" binding:"required,min=2,max=255"`
	Location string `json:"location" binding:"omitempty,max=500"`
}

// UpdateBuildingRequest represents building update request (partial)
type UpdateBuildingRequest struct {
	Code     string `json:"code" binding:"omitempty,min=1,max=100"`
	Name     string `json:"name" binding:"omitempty,min=2,max=255"`
	Location string `json:"location" binding:"omitempty,max=500"`
}

// BuildingResponse represents building response
type BuildingResponse struct {
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// ===== Floor DTOs =====

// CreateFloorRequest represents floor creation request
type CreateFloorRequest struct {
	BuildingID  uint   `json:"building_id" binding:"required"`
	FloorNumber int    `json:"floor_number" binding:"required"`
	Name        string `json:"name" binding:"required,min=2,max=255"`
}

// UpdateFloorRequest represents floor update request (partial)
type UpdateFloorRequest struct {
	FloorNumber int    `json:"floor_number" binding:"omitempty"`
	Name        string `json:"name" binding:"omitempty,min=2,max=255"`
}

// FloorResponse represents floor response
type FloorResponse struct {
	ID          uint              `json:"id"`
	BuildingID  uint              `json:"building_id"`
	Building    *BuildingResponse `json:"building,omitempty"`
	FloorNumber int               `json:"floor_number"`
	Name        string            `json:"name"`
	CreatedAt   int64             `json:"created_at"`
	UpdatedAt   int64             `json:"updated_at"`
}

// ===== Room DTOs =====

// CreateRoomRequest represents room creation request
type CreateRoomRequest struct {
	FloorID uint   `json:"floor_id" binding:"required"`
	Code    string `json:"code" binding:"required,min=1,max=100"`
	Name    string `json:"name" binding:"required,min=2,max=255"`
}

// UpdateRoomRequest represents room update request (partial)
type UpdateRoomRequest struct {
	Code string `json:"code" binding:"omitempty,min=1,max=100"`
	Name string `json:"name" binding:"omitempty,min=2,max=255"`
}

// RoomResponse represents room response
type RoomResponse struct {
	ID        uint           `json:"id"`
	FloorID   uint           `json:"floor_id"`
	Floor     *FloorResponse `json:"floor,omitempty"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
}

// ===== ComponentCategory DTOs =====

// CreateComponentCategoryRequest represents component category creation request
type CreateComponentCategoryRequest struct {
	Code        string `json:"code" binding:"required,min=1,max=100"`
	Name        string `json:"name" binding:"required,min=2,max=255"`
	Description string `json:"description" binding:"omitempty"`
}

// UpdateComponentCategoryRequest represents component category update request (partial)
type UpdateComponentCategoryRequest struct {
	Code        string `json:"code" binding:"omitempty,min=1,max=100"`
	Name        string `json:"name" binding:"omitempty,min=2,max=255"`
	Description string `json:"description" binding:"omitempty"`
}

// ComponentCategoryResponse represents component category response
type ComponentCategoryResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// ===== Component DTOs =====

// CreateComponentRequest represents component creation request
type CreateComponentRequest struct {
	RoomID          uint   `json:"room_id" binding:"required"`
	CategoryID      uint   `json:"category_id" binding:"required"`
	Code            string `json:"code" binding:"required,min=1,max=100"`
	Name            string `json:"name" binding:"required,min=2,max=255"`
	Brand           string `json:"brand" binding:"omitempty,max=255"`
	Specification   string `json:"specification" binding:"omitempty"`
	ProcurementYear int    `json:"procurement_year" binding:"omitempty"`
}

// UpdateComponentRequest represents component update request (partial)
type UpdateComponentRequest struct {
	Code            string `json:"code" binding:"omitempty,min=1,max=100"`
	Name            string `json:"name" binding:"omitempty,min=2,max=255"`
	Brand           string `json:"brand" binding:"omitempty,max=255"`
	Specification   string `json:"specification" binding:"omitempty"`
	ProcurementYear int    `json:"procurement_year" binding:"omitempty"`
}

// ComponentResponse represents component response
type ComponentResponse struct {
	ID              uint   `json:"id"`
	RoomID          uint   `json:"room_id"`
	CategoryID      uint   `json:"category_id"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	Brand           string `json:"brand"`
	Specification   string `json:"specification"`
	ProcurementYear int    `json:"procurement_year"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
}
