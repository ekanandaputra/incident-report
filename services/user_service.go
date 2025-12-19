package services

import (
	"errors"
	"gorm.io/gorm"
	"incident-report/config"
	"incident-report/models"
	"incident-report/utils"
)

// UserService handles all user-related business logic
type UserService struct{}

// NewUserService creates a new instance of UserService
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser creates a new user in the database
func (us *UserService) CreateUser(req *utils.CreateUserRequest) (*utils.UserResponse, error) {
	// Validate input
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("name and email are required")
	}

	// Create user model instance
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	// Save to database
	result := config.DB.Create(&user)
	if result.Error != nil {
		// Check for duplicate email error
		if result.Error.Error() == "UNIQUE constraint failed: users.email" {
			return nil, errors.New("email already exists")
		}
		return nil, result.Error
	}

	// Return user response DTO
	return &utils.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// GetUserByID retrieves a user by their ID
func (us *UserService) GetUserByID(id uint) (*utils.UserResponse, error) {
	var user models.User

	// Query database for user with given ID
	result := config.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &utils.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// GetAllUsers retrieves all users with pagination support
func (us *UserService) GetAllUsers(page int, pageSize int) ([]utils.UserResponse, int64, error) {
	// Set default pagination values
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var users []models.User
	var total int64

	// Count total records
	if err := config.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset for pagination
	offset := (page - 1) * pageSize

	// Fetch paginated results
	result := config.DB.Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	// Convert to response DTOs
	var responses []utils.UserResponse
	for _, user := range users {
		responses = append(responses, utils.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return responses, total, nil
}

// UpdateUser updates an existing user's information
func (us *UserService) UpdateUser(id uint, req *utils.UpdateUserRequest) (*utils.UserResponse, error) {
	// Find user first
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	// Update only provided fields
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	// Save changes to database
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &utils.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// DeleteUser performs a soft delete of a user (marks as deleted, doesn't remove from DB)
func (us *UserService) DeleteUser(id uint) error {
	// Find user first to ensure it exists
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}

	// Perform soft delete (sets deleted_at timestamp)
	result = config.DB.Delete(&user)
	return result.Error
}
