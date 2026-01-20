package services

import (
	"errors"
	"incident-report/config"
	"incident-report/models"
	"incident-report/utils"

	"gorm.io/gorm"
)

// FloorService handles all floor-related business logic
type FloorService struct{}

// NewFloorService creates a new instance of FloorService
func NewFloorService() *FloorService {
	return &FloorService{}
}

// CreateFloor creates a new floor in the database
func (fs *FloorService) CreateFloor(req *utils.CreateFloorRequest) (*utils.FloorResponse, error) {
	if req.BuildingID == 0 || req.FloorNumber == 0 || req.Name == "" {
		return nil, errors.New("building_id, floor_number, and name are required")
	}

	// Verify building exists
	var building models.Building
	if err := config.DB.First(&building, req.BuildingID).Error; err != nil {
		return nil, errors.New("building not found")
	}

	floor := models.Floor{
		BuildingID: req.BuildingID,
		Number:     req.FloorNumber,
		Name:       req.Name,
	}

	result := config.DB.Create(&floor)
	if result.Error != nil {
		return nil, result.Error
	}

	return &utils.FloorResponse{
		ID:          floor.ID,
		BuildingID:  floor.BuildingID,
		FloorNumber: floor.Number,
		Name:        floor.Name,
		CreatedAt:   floor.CreatedAt,
		UpdatedAt:   floor.UpdatedAt,
	}, nil
}

// GetFloorByID retrieves a floor by its ID
func (fs *FloorService) GetFloorByID(id uint) (*utils.FloorResponse, error) {
	var floor models.Floor

	result := config.DB.First(&floor, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("floor not found")
		}
		return nil, result.Error
	}

	return &utils.FloorResponse{
		ID:          floor.ID,
		BuildingID:  floor.BuildingID,
		FloorNumber: floor.Number,
		Name:        floor.Name,
		CreatedAt:   floor.CreatedAt,
		UpdatedAt:   floor.UpdatedAt,
	}, nil
}

// GetFloorsByBuildingID retrieves all floors in a building
func (fs *FloorService) GetFloorsByBuildingID(buildingID uint, page, pageSize int) ([]utils.FloorResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var floors []models.Floor
	var total int64

	if err := config.DB.Where("building_id = ?", buildingID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Where("building_id = ?", buildingID).Offset(offset).Limit(pageSize).Find(&floors)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.FloorResponse
	for _, floor := range floors {
		responses = append(responses, utils.FloorResponse{
			ID:          floor.ID,
			BuildingID:  floor.BuildingID,
			FloorNumber: floor.Number,
			Name:        floor.Name,
			CreatedAt:   floor.CreatedAt,
			UpdatedAt:   floor.UpdatedAt,
		})
	}

	return responses, total, nil
}

// GetAllFloors retrieves all floors with pagination and building info
func (fs *FloorService) GetAllFloors(page, pageSize int) ([]utils.FloorResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var floors []models.Floor
	var total int64

	if err := config.DB.Model(&models.Floor{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Preload("Building").Offset(offset).Limit(pageSize).Find(&floors)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.FloorResponse
	for _, floor := range floors {
		buildingResp := &utils.BuildingResponse{
			ID:        floor.Building.ID,
			Code:      floor.Building.Code,
			Name:      floor.Building.Name,
			Location:  floor.Building.Location,
			CreatedAt: floor.Building.CreatedAt,
		}
		responses = append(responses, utils.FloorResponse{
			ID:          floor.ID,
			BuildingID:  floor.BuildingID,
			Building:    buildingResp,
			FloorNumber: floor.Number,
			Name:        floor.Name,
			CreatedAt:   floor.CreatedAt,
			UpdatedAt:   floor.UpdatedAt,
		})
	}

	return responses, total, nil
}

// UpdateFloor updates an existing floor
func (fs *FloorService) UpdateFloor(id uint, req *utils.UpdateFloorRequest) (*utils.FloorResponse, error) {
	var floor models.Floor

	result := config.DB.First(&floor, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("floor not found")
		}
		return nil, result.Error
	}

	if req.FloorNumber != 0 {
		floor.Number = req.FloorNumber
	}
	if req.Name != "" {
		floor.Name = req.Name
	}

	if err := config.DB.Save(&floor).Error; err != nil {
		return nil, err
	}

	return &utils.FloorResponse{
		ID:          floor.ID,
		BuildingID:  floor.BuildingID,
		FloorNumber: floor.Number,
		Name:        floor.Name,
		CreatedAt:   floor.CreatedAt,
		UpdatedAt:   floor.UpdatedAt,
	}, nil
}

// DeleteFloor performs a soft delete of a floor
func (fs *FloorService) DeleteFloor(id uint) error {
	var floor models.Floor

	result := config.DB.First(&floor, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("floor not found")
		}
		return result.Error
	}

	result = config.DB.Delete(&floor)
	return result.Error
}

// GetFloorWithRooms retrieves a floor with all its rooms
func (fs *FloorService) GetFloorWithRooms(id uint) (*models.Floor, error) {
	var floor models.Floor

	result := config.DB.Preload("Rooms").First(&floor, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("floor not found")
		}
		return nil, result.Error
	}

	return &floor, nil
}
