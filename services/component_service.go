package services

import (
	"errors"
	"incident-report/config"
	"incident-report/models"
	"incident-report/utils"

	"gorm.io/gorm"
)

// ComponentService handles all component-related business logic
type ComponentService struct{}

// NewComponentService creates a new instance of ComponentService
func NewComponentService() *ComponentService {
	return &ComponentService{}
}

// CreateComponent creates a new component in the database
func (cs *ComponentService) CreateComponent(req *utils.CreateComponentRequest) (*utils.ComponentResponse, error) {
	if req.RoomID == 0 || req.CategoryID == 0 || req.Code == "" || req.Name == "" {
		return nil, errors.New("room_id, category_id, code, and name are required")
	}

	// Verify room exists
	var room models.Room
	if err := config.DB.First(&room, req.RoomID).Error; err != nil {
		return nil, errors.New("room not found")
	}

	// Verify category exists
	var category models.ComponentCategory
	if err := config.DB.First(&category, req.CategoryID).Error; err != nil {
		return nil, errors.New("component category not found")
	}

	component := models.Component{
		RoomID:          req.RoomID,
		CategoryID:      req.CategoryID,
		Code:            req.Code,
		Name:            req.Name,
		Brand:           req.Brand,
		Specification:   req.Specification,
		ProcurementYear: req.ProcurementYear,
	}

	result := config.DB.Create(&component)
	if result.Error != nil {
		return nil, result.Error
	}

	return &utils.ComponentResponse{
		ID:              component.ID,
		RoomID:          component.RoomID,
		CategoryID:      component.CategoryID,
		Code:            component.Code,
		Name:            component.Name,
		Brand:           component.Brand,
		Specification:   component.Specification,
		ProcurementYear: component.ProcurementYear,
		CreatedAt:       component.CreatedAt,
		UpdatedAt:       component.UpdatedAt,
	}, nil
}

// GetComponentByID retrieves a component by its ID
func (cs *ComponentService) GetComponentByID(id uint) (*utils.ComponentResponse, error) {
	var component models.Component

	result := config.DB.First(&component, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("component not found")
		}
		return nil, result.Error
	}

	return &utils.ComponentResponse{
		ID:              component.ID,
		RoomID:          component.RoomID,
		CategoryID:      component.CategoryID,
		Code:            component.Code,
		Name:            component.Name,
		Brand:           component.Brand,
		Specification:   component.Specification,
		ProcurementYear: component.ProcurementYear,
		CreatedAt:       component.CreatedAt,
		UpdatedAt:       component.UpdatedAt,
	}, nil
}

// GetComponentsByRoomID retrieves all components in a room
func (cs *ComponentService) GetComponentsByRoomID(roomID uint, page, pageSize int) ([]utils.ComponentResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var components []models.Component
	var total int64

	if err := config.DB.Where("room_id = ?", roomID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Where("room_id = ?", roomID).Offset(offset).Limit(pageSize).Find(&components)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.ComponentResponse
	for _, component := range components {
		responses = append(responses, utils.ComponentResponse{
			ID:              component.ID,
			RoomID:          component.RoomID,
			CategoryID:      component.CategoryID,
			Code:            component.Code,
			Name:            component.Name,
			Brand:           component.Brand,
			Specification:   component.Specification,
			ProcurementYear: component.ProcurementYear,
			CreatedAt:       component.CreatedAt,
			UpdatedAt:       component.UpdatedAt,
		})
	}

	return responses, total, nil
}

// GetComponentsByCategoryID retrieves all components in a category
func (cs *ComponentService) GetComponentsByCategoryID(categoryID uint, page, pageSize int) ([]utils.ComponentResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var components []models.Component
	var total int64

	if err := config.DB.Where("category_id = ?", categoryID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Where("category_id = ?", categoryID).Offset(offset).Limit(pageSize).Find(&components)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.ComponentResponse
	for _, component := range components {
		responses = append(responses, utils.ComponentResponse{
			ID:              component.ID,
			RoomID:          component.RoomID,
			CategoryID:      component.CategoryID,
			Code:            component.Code,
			Name:            component.Name,
			Brand:           component.Brand,
			Specification:   component.Specification,
			ProcurementYear: component.ProcurementYear,
			CreatedAt:       component.CreatedAt,
			UpdatedAt:       component.UpdatedAt,
		})
	}

	return responses, total, nil
}

// UpdateComponent updates an existing component
func (cs *ComponentService) UpdateComponent(id uint, req *utils.UpdateComponentRequest) (*utils.ComponentResponse, error) {
	var component models.Component

	result := config.DB.First(&component, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("component not found")
		}
		return nil, result.Error
	}

	if req.Code != "" {
		component.Code = req.Code
	}
	if req.Name != "" {
		component.Name = req.Name
	}
	if req.Brand != "" {
		component.Brand = req.Brand
	}
	if req.Specification != "" {
		component.Specification = req.Specification
	}
	if req.ProcurementYear != 0 {
		component.ProcurementYear = req.ProcurementYear
	}

	if err := config.DB.Save(&component).Error; err != nil {
		return nil, err
	}

	return &utils.ComponentResponse{
		ID:              component.ID,
		RoomID:          component.RoomID,
		CategoryID:      component.CategoryID,
		Code:            component.Code,
		Name:            component.Name,
		Brand:           component.Brand,
		Specification:   component.Specification,
		ProcurementYear: component.ProcurementYear,
		CreatedAt:       component.CreatedAt,
		UpdatedAt:       component.UpdatedAt,
	}, nil
}

// DeleteComponent performs a soft delete of a component
func (cs *ComponentService) DeleteComponent(id uint) error {
	var component models.Component

	result := config.DB.First(&component, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("component not found")
		}
		return result.Error
	}

	result = config.DB.Delete(&component)
	return result.Error
}
