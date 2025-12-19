package services

import (
	"errors"
	"incident-report/config"
	"incident-report/models"
	"incident-report/utils"

	"gorm.io/gorm"
)

// BuildingService handles all building-related business logic
type BuildingService struct{}

// NewBuildingService creates a new instance of BuildingService
func NewBuildingService() *BuildingService {
	return &BuildingService{}
}

// CreateBuilding creates a new building in the database
func (bs *BuildingService) CreateBuilding(req *utils.CreateBuildingRequest) (*utils.BuildingResponse, error) {
	if req.Code == "" || req.Name == "" {
		return nil, errors.New("code and name are required")
	}

	building := models.Building{
		Code:     req.Code,
		Name:     req.Name,
		Location: req.Location,
	}

	result := config.DB.Create(&building)
	if result.Error != nil {
		return nil, result.Error
	}

	return &utils.BuildingResponse{
		ID:        building.ID,
		Code:      building.Code,
		Name:      building.Name,
		Location:  building.Location,
		CreatedAt: building.CreatedAt,
		UpdatedAt: building.UpdatedAt,
	}, nil
}

// GetBuildingByID retrieves a building by its ID
func (bs *BuildingService) GetBuildingByID(id uint) (*utils.BuildingResponse, error) {
	var building models.Building

	result := config.DB.First(&building, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("building not found")
		}
		return nil, result.Error
	}

	return &utils.BuildingResponse{
		ID:        building.ID,
		Code:      building.Code,
		Name:      building.Name,
		Location:  building.Location,
		CreatedAt: building.CreatedAt,
		UpdatedAt: building.UpdatedAt,
	}, nil
}

// GetAllBuildings retrieves all buildings with pagination
func (bs *BuildingService) GetAllBuildings(page, pageSize int) ([]utils.BuildingResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var buildings []models.Building
	var total int64

	if err := config.DB.Model(&models.Building{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Offset(offset).Limit(pageSize).Find(&buildings)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.BuildingResponse
	for _, building := range buildings {
		responses = append(responses, utils.BuildingResponse{
			ID:        building.ID,
			Code:      building.Code,
			Name:      building.Name,
			Location:  building.Location,
			CreatedAt: building.CreatedAt,
			UpdatedAt: building.UpdatedAt,
		})
	}

	return responses, total, nil
}

// UpdateBuilding updates an existing building
func (bs *BuildingService) UpdateBuilding(id uint, req *utils.UpdateBuildingRequest) (*utils.BuildingResponse, error) {
	var building models.Building

	result := config.DB.First(&building, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("building not found")
		}
		return nil, result.Error
	}

	if req.Code != "" {
		building.Code = req.Code
	}
	if req.Name != "" {
		building.Name = req.Name
	}
	if req.Location != "" {
		building.Location = req.Location
	}

	if err := config.DB.Save(&building).Error; err != nil {
		return nil, err
	}

	return &utils.BuildingResponse{
		ID:        building.ID,
		Code:      building.Code,
		Name:      building.Name,
		Location:  building.Location,
		CreatedAt: building.CreatedAt,
		UpdatedAt: building.UpdatedAt,
	}, nil
}

// DeleteBuilding performs a soft delete of a building
func (bs *BuildingService) DeleteBuilding(id uint) error {
	var building models.Building

	result := config.DB.First(&building, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("building not found")
		}
		return result.Error
	}

	result = config.DB.Delete(&building)
	return result.Error
}

// GetBuildingWithFloors retrieves a building with all its floors
func (bs *BuildingService) GetBuildingWithFloors(id uint) (*models.Building, error) {
	var building models.Building

	result := config.DB.Preload("Floors").First(&building, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("building not found")
		}
		return nil, result.Error
	}

	return &building, nil
}
