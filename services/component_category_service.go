package services

import (
	"errors"
	"incident-report/config"
	"incident-report/models"
	"incident-report/utils"

	"gorm.io/gorm"
)

// ComponentCategoryService handles all component category-related business logic
type ComponentCategoryService struct{}

// NewComponentCategoryService creates a new instance of ComponentCategoryService
func NewComponentCategoryService() *ComponentCategoryService {
	return &ComponentCategoryService{}
}

// CreateComponentCategory creates a new component category in the database
func (ccs *ComponentCategoryService) CreateComponentCategory(req *utils.CreateComponentCategoryRequest) (*utils.ComponentCategoryResponse, error) {
	if req.Code == "" || req.Name == "" {
		return nil, errors.New("code and name are required")
	}

	category := models.ComponentCategory{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
	}

	result := config.DB.Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return &utils.ComponentCategoryResponse{
		ID:          category.ID,
		Code:        category.Code,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

// GetComponentCategoryByID retrieves a component category by its ID
func (ccs *ComponentCategoryService) GetComponentCategoryByID(id uint) (*utils.ComponentCategoryResponse, error) {
	var category models.ComponentCategory

	result := config.DB.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("component category not found")
		}
		return nil, result.Error
	}

	return &utils.ComponentCategoryResponse{
		ID:          category.ID,
		Code:        category.Code,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

// GetAllComponentCategories retrieves all component categories with pagination
func (ccs *ComponentCategoryService) GetAllComponentCategories(page, pageSize int) ([]utils.ComponentCategoryResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var categories []models.ComponentCategory
	var total int64

	if err := config.DB.Model(&models.ComponentCategory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Offset(offset).Limit(pageSize).Find(&categories)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.ComponentCategoryResponse
	for _, category := range categories {
		responses = append(responses, utils.ComponentCategoryResponse{
			ID:          category.ID,
			Code:        category.Code,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		})
	}

	return responses, total, nil
}

// UpdateComponentCategory updates an existing component category
func (ccs *ComponentCategoryService) UpdateComponentCategory(id uint, req *utils.UpdateComponentCategoryRequest) (*utils.ComponentCategoryResponse, error) {
	var category models.ComponentCategory

	result := config.DB.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("component category not found")
		}
		return nil, result.Error
	}

	if req.Code != "" {
		category.Code = req.Code
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}

	if err := config.DB.Save(&category).Error; err != nil {
		return nil, err
	}

	return &utils.ComponentCategoryResponse{
		ID:          category.ID,
		Code:        category.Code,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

// DeleteComponentCategory performs a soft delete of a component category
func (ccs *ComponentCategoryService) DeleteComponentCategory(id uint) error {
	var category models.ComponentCategory

	result := config.DB.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("component category not found")
		}
		return result.Error
	}

	result = config.DB.Delete(&category)
	return result.Error
}

// GetCategoryWithComponents retrieves a category with all its components
func (ccs *ComponentCategoryService) GetCategoryWithComponents(id uint) (*models.ComponentCategory, error) {
	var category models.ComponentCategory

	result := config.DB.Preload("Components").First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("component category not found")
		}
		return nil, result.Error
	}

	return &category, nil
}
