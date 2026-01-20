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
	if req.CategoryID == 0 || req.Code == "" || req.Name == "" {
		return nil, errors.New("category_id, code, and name are required")
	}

	// Verify room exists if provided
	if req.RoomID != nil {
		var room models.Room
		if err := config.DB.First(&room, *req.RoomID).Error; err != nil {
			return nil, errors.New("room not found")
		}
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

// GetAllComponents retrieves all components with pagination and nested building, floor, and room info
func (cs *ComponentService) GetAllComponents(page, pageSize int) ([]utils.ComponentResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var components []models.Component
	var total int64

	if err := config.DB.Model(&models.Component{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Preload("Room.Floor.Building").Offset(offset).Limit(pageSize).Find(&components)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.ComponentResponse
	for _, component := range components {
		response := utils.ComponentResponse{
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
		}

		// if component.Room != nil && component.Room.ID != 0 {
		// 	response.Room = &utils.RoomResponse{
		// 		ID:        component.Room.ID,
		// 		FloorID:   component.Room.FloorID,
		// 		Code:      component.Room.Code,
		// 		Name:      component.Room.Name,
		// 		CreatedAt: component.Room.CreatedAt,
		// 		UpdatedAt: component.Room.UpdatedAt,
		// 	}
		// 	if component.Room.Floor != nil && component.Room.Floor.ID != 0 {
		// 		response.Room.Floor = &utils.FloorResponse{
		// 			ID:          component.Room.Floor.ID,
		// 			BuildingID:  component.Room.Floor.BuildingID,
		// 			FloorNumber: component.Room.Floor.Number,
		// 			Name:        component.Room.Floor.Name,
		// 			CreatedAt:   component.Room.Floor.CreatedAt,
		// 			UpdatedAt:   component.Room.Floor.UpdatedAt,
		// 		}
		// 		if component.Room.Floor.Building != nil && component.Room.Floor.Building.ID != 0 {
		// 			response.Room.Floor.Building = &utils.BuildingResponse{
		// 				ID:        component.Room.Floor.Building.ID,
		// 				Code:      component.Room.Floor.Building.Code,
		// 				Name:      component.Room.Floor.Building.Name,
		// 				Location:  component.Room.Floor.Building.Location,
		// 				CreatedAt: component.Room.Floor.Building.CreatedAt,
		// 				UpdatedAt: component.Room.Floor.Building.UpdatedAt,
		// 			}
		// 		}
		// 	}
		// }

		responses = append(responses, response)
	}

	return responses, total, nil
}
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

// AssignRoomToComponent assigns a room to an existing component
func (cs *ComponentService) AssignRoomToComponent(componentID uint, req *utils.AssignRoomRequest) (*utils.ComponentResponse, error) {
	var component models.Component

	result := config.DB.First(&component, componentID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("component not found")
		}
		return nil, result.Error
	}

	// Verify room exists
	var room models.Room
	if err := config.DB.First(&room, req.RoomID).Error; err != nil {
		return nil, errors.New("room not found")
	}

	component.RoomID = &req.RoomID
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
