package services

import (
	"errors"
	"incident-report/config"
	"incident-report/models"
	"incident-report/utils"

	"gorm.io/gorm"
)

// RoomService handles all room-related business logic
type RoomService struct{}

// NewRoomService creates a new instance of RoomService
func NewRoomService() *RoomService {
	return &RoomService{}
}

// CreateRoom creates a new room in the database
func (rs *RoomService) CreateRoom(req *utils.CreateRoomRequest) (*utils.RoomResponse, error) {
	if req.FloorID == 0 || req.Code == "" || req.Name == "" {
		return nil, errors.New("floor_id, code, and name are required")
	}

	// Verify floor exists
	var floor models.Floor
	if err := config.DB.First(&floor, req.FloorID).Error; err != nil {
		return nil, errors.New("floor not found")
	}

	room := models.Room{
		FloorID: req.FloorID,
		Code:    req.Code,
		Name:    req.Name,
	}

	result := config.DB.Create(&room)
	if result.Error != nil {
		return nil, result.Error
	}

	return &utils.RoomResponse{
		ID:        room.ID,
		FloorID:   room.FloorID,
		Code:      room.Code,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}, nil
}

// GetRoomByID retrieves a room by its ID
func (rs *RoomService) GetRoomByID(id uint) (*utils.RoomResponse, error) {
	var room models.Room

	result := config.DB.First(&room, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, result.Error
	}

	return &utils.RoomResponse{
		ID:        room.ID,
		FloorID:   room.FloorID,
		Code:      room.Code,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}, nil
}

// GetRoomsByFloorID retrieves all rooms on a floor
func (rs *RoomService) GetRoomsByFloorID(floorID uint, page, pageSize int) ([]utils.RoomResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var rooms []models.Room
	var total int64

	if err := config.DB.Where("floor_id = ?", floorID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Where("floor_id = ?", floorID).Offset(offset).Limit(pageSize).Find(&rooms)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.RoomResponse
	for _, room := range rooms {
		responses = append(responses, utils.RoomResponse{
			ID:        room.ID,
			FloorID:   room.FloorID,
			Code:      room.Code,
			Name:      room.Name,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		})
	}

	return responses, total, nil
}

// GetAllRooms retrieves all rooms with pagination and floor/building info
func (rs *RoomService) GetAllRooms(page, pageSize int) ([]utils.RoomResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var rooms []models.Room
	var total int64

	if err := config.DB.Model(&models.Room{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	result := config.DB.Preload("Floor.Building").Offset(offset).Limit(pageSize).Find(&rooms)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var responses []utils.RoomResponse
	for _, room := range rooms {
		buildingResp := &utils.BuildingResponse{
			ID:        room.Floor.Building.ID,
			Code:      room.Floor.Building.Code,
			Name:      room.Floor.Building.Name,
			Location:  room.Floor.Building.Location,
			CreatedAt: room.Floor.Building.CreatedAt,
		}
		floorResp := &utils.FloorResponse{
			ID:          room.Floor.ID,
			BuildingID:  room.Floor.BuildingID,
			Building:    buildingResp,
			FloorNumber: room.Floor.Number,
			Name:        room.Floor.Name,
			CreatedAt:   room.Floor.CreatedAt,
			UpdatedAt:   room.Floor.UpdatedAt,
		}
		responses = append(responses, utils.RoomResponse{
			ID:        room.ID,
			FloorID:   room.FloorID,
			Floor:     floorResp,
			Code:      room.Code,
			Name:      room.Name,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		})
	}

	return responses, total, nil
}

// UpdateRoom updates an existing room
func (rs *RoomService) UpdateRoom(id uint, req *utils.UpdateRoomRequest) (*utils.RoomResponse, error) {
	var room models.Room

	result := config.DB.First(&room, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, result.Error
	}

	if req.Code != "" {
		room.Code = req.Code
	}
	if req.Name != "" {
		room.Name = req.Name
	}

	if err := config.DB.Save(&room).Error; err != nil {
		return nil, err
	}

	return &utils.RoomResponse{
		ID:        room.ID,
		FloorID:   room.FloorID,
		Code:      room.Code,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}, nil
}

// DeleteRoom performs a soft delete of a room
func (rs *RoomService) DeleteRoom(id uint) error {
	var room models.Room

	result := config.DB.First(&room, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("room not found")
		}
		return result.Error
	}

	result = config.DB.Delete(&room)
	return result.Error
}

// GetRoomWithComponents retrieves a room with all its components
func (rs *RoomService) GetRoomWithComponents(id uint) (*models.Room, error) {
	var room models.Room

	result := config.DB.Preload("Components").First(&room, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, result.Error
	}

	return &room, nil
}
