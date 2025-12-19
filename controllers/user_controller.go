package controllers

import (
	"incident-report/services"
	"incident-report/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController handles HTTP requests for user operations
type UserController struct {
	userService *services.UserService
}

// NewUserController creates a new instance of UserController with dependency injection
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser handles POST /api/v1/users request to create a new user
// @param c *gin.Context
// Request body: CreateUserRequest (name, email)
// Response: UserResponse with HTTP 201 Created
func (uc *UserController) CreateUser(c *gin.Context) {
	var req utils.CreateUserRequest

	// Bind and validate request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Call service to create user
	user, err := uc.userService.CreateUser(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

// GetUser handles GET /api/v1/users/:id request to retrieve a specific user
// @param c *gin.Context with :id parameter
// Response: UserResponse with HTTP 200 OK
func (uc *UserController) GetUser(c *gin.Context) {
	// Extract user ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", "ID must be a valid number")
		return
	}

	// Call service to fetch user
	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

// GetAllUsers handles GET /api/v1/users request to retrieve all users with pagination
// @param c *gin.Context with optional query parameters: page, page_size
// Response: PaginatedResponse with array of users and HTTP 200 OK
func (uc *UserController) GetAllUsers(c *gin.Context) {
	var pagination utils.PaginationQuery

	// Bind query parameters with default values
	if err := c.ShouldBindQuery(&pagination); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	// Set defaults if not provided
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}

	// Call service to fetch paginated users
	users, total, err := uc.userService.GetAllUsers(pagination.Page, pagination.PageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}

	// Calculate total pages
	totalPage := (int(total) + pagination.PageSize - 1) / pagination.PageSize

	// Create paginated response
	response := utils.PaginatedResponse{
		Data:      users,
		Page:      pagination.Page,
		PageSize:  pagination.PageSize,
		Total:     total,
		TotalPage: totalPage,
	}

	utils.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", response)
}

// UpdateUser handles PUT /api/v1/users/:id request to update a user
// @param c *gin.Context with :id parameter
// Request body: UpdateUserRequest (partial fields)
// Response: UserResponse with HTTP 200 OK
func (uc *UserController) UpdateUser(c *gin.Context) {
	// Extract user ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", "ID must be a valid number")
		return
	}

	var req utils.UpdateUserRequest

	// Bind and validate request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Call service to update user
	user, err := uc.userService.UpdateUser(uint(id), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update user", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User updated successfully", user)
}

// DeleteUser handles DELETE /api/v1/users/:id request to delete a user
// @param c *gin.Context with :id parameter
// Response: HTTP 204 No Content on success
func (uc *UserController) DeleteUser(c *gin.Context) {
	// Extract user ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", "ID must be a valid number")
		return
	}

	// Call service to delete user
	err = uc.userService.DeleteUser(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to delete user", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}
