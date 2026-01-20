package controllers

import (
	"net/http"
	"strconv"

	"incident-report/services"
	"incident-report/utils"

	"github.com/gin-gonic/gin"
)

// RoomController handles HTTP requests for room operations
type RoomController struct {
	service *services.RoomService
}

// NewRoomController creates a new instance of RoomController
func NewRoomController() *RoomController {
	return &RoomController{
		service: services.NewRoomService(),
	}
}

// CreateRoom handles POST /api/v1/rooms
// @Summary Create a new room
// @Description Creates a new room with the provided information
// @Accept json
// @Produce json
// @Param request body utils.CreateRoomRequest true "Room details"
// @Success 201 {object} utils.RoomResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/rooms [post]
func (rc *RoomController) CreateRoom(c *gin.Context) {
	var req utils.CreateRoomRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	room, err := rc.service.CreateRoom(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create room", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Room created successfully", room)
}

// GetAllRooms handles GET /api/v1/rooms
// @Summary Get all rooms
// @Description Retrieves all rooms with floor and building information and pagination
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Records per page" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/rooms [get]
func (rc *RoomController) GetAllRooms(c *gin.Context) {
	var pagination utils.PaginationQuery

	if err := c.ShouldBindQuery(&pagination); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}

	rooms, total, err := rc.service.GetAllRooms(pagination.Page, pagination.PageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch rooms", err.Error())
		return
	}

	totalPage := (int(total) + pagination.PageSize - 1) / pagination.PageSize

	response := utils.PaginatedResponse{
		Data:      rooms,
		Page:      pagination.Page,
		PageSize:  pagination.PageSize,
		Total:     total,
		TotalPage: totalPage,
	}

	utils.SuccessResponse(c, http.StatusOK, "Rooms retrieved successfully", response)
}

// GetRoom handles GET /api/v1/rooms/:id
// @Summary Get room by ID
// @Description Retrieves a specific room by its ID
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} utils.RoomResponse
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/rooms/{id} [get]
func (rc *RoomController) GetRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid room ID", err.Error())
		return
	}

	room, err := rc.service.GetRoomByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Room not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Room retrieved successfully", room)
}

// GetRoomsByFloor handles GET /api/v1/floors/:floorId/rooms
// @Summary Get all rooms on a floor
// @Description Retrieves all rooms on a specific floor with pagination
// @Produce json
// @Param floorId path int true "Floor ID"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/floors/{floorId}/rooms [get]
func (rc *RoomController) GetRoomsByFloor(c *gin.Context) {
	floorID, err := strconv.ParseUint(c.Param("floorId"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid floor ID", err.Error())
		return
	}

	var query utils.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	page := query.Page
	pageSize := query.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	rooms, total, err := rc.service.GetRoomsByFloorID(uint(floorID), page, pageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve rooms", err.Error())
		return
	}

	response := utils.PaginatedResponse{
		Data:      rooms,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	utils.SuccessResponse(c, http.StatusOK, "Rooms retrieved successfully", response)
}

// UpdateRoom handles PUT /api/v1/rooms/:id
// @Summary Update a room
// @Description Updates an existing room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Param request body utils.UpdateRoomRequest true "Room details to update"
// @Success 200 {object} utils.RoomResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/rooms/{id} [put]
func (rc *RoomController) UpdateRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid room ID", err.Error())
		return
	}

	var req utils.UpdateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	room, err := rc.service.UpdateRoom(uint(id), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update room", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Room updated successfully", room)
}

// DeleteRoom handles DELETE /api/v1/rooms/:id
// @Summary Delete a room
// @Description Deletes (soft delete) an existing room
// @Produce json
// @Param id path int true "Room ID"
// @Success 204 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/rooms/{id} [delete]
func (rc *RoomController) DeleteRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid room ID", err.Error())
		return
	}

	err = rc.service.DeleteRoom(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Room not found", err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
