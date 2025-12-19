package controllers

import (
	"net/http"
	"strconv"

	"incident-report/services"
	"incident-report/utils"

	"github.com/gin-gonic/gin"
)

// FloorController handles HTTP requests for floor operations
type FloorController struct {
	service *services.FloorService
}

// NewFloorController creates a new instance of FloorController
func NewFloorController() *FloorController {
	return &FloorController{
		service: services.NewFloorService(),
	}
}

// CreateFloor handles POST /api/v1/floors
// @Summary Create a new floor
// @Description Creates a new floor with the provided information
// @Accept json
// @Produce json
// @Param request body utils.CreateFloorRequest true "Floor details"
// @Success 201 {object} utils.FloorResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/floors [post]
func (fc *FloorController) CreateFloor(c *gin.Context) {
	var req utils.CreateFloorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	floor, err := fc.service.CreateFloor(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create floor", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Floor created successfully", floor)
}

// GetFloor handles GET /api/v1/floors/:id
// @Summary Get floor by ID
// @Description Retrieves a specific floor by its ID
// @Produce json
// @Param id path int true "Floor ID"
// @Success 200 {object} utils.FloorResponse
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/floors/{id} [get]
func (fc *FloorController) GetFloor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid floor ID", err.Error())
		return
	}

	floor, err := fc.service.GetFloorByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Floor not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Floor retrieved successfully", floor)
}

// GetFloorsByBuilding handles GET /api/v1/buildings/:buildingId/floors
// @Summary Get all floors in a building
// @Description Retrieves all floors in a specific building with pagination
// @Produce json
// @Param buildingId path int true "Building ID"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/buildings/{buildingId}/floors [get]
func (fc *FloorController) GetFloorsByBuilding(c *gin.Context) {
	buildingID, err := strconv.ParseUint(c.Param("buildingId"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid building ID", err.Error())
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

	floors, total, err := fc.service.GetFloorsByBuildingID(uint(buildingID), page, pageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve floors", err.Error())
		return
	}

	response := utils.PaginatedResponse{
		Data:      floors,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	utils.SuccessResponse(c, http.StatusOK, "Floors retrieved successfully", response)
}

// UpdateFloor handles PUT /api/v1/floors/:id
// @Summary Update a floor
// @Description Updates an existing floor
// @Accept json
// @Produce json
// @Param id path int true "Floor ID"
// @Param request body utils.UpdateFloorRequest true "Floor details to update"
// @Success 200 {object} utils.FloorResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/floors/{id} [put]
func (fc *FloorController) UpdateFloor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid floor ID", err.Error())
		return
	}

	var req utils.UpdateFloorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	floor, err := fc.service.UpdateFloor(uint(id), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update floor", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Floor updated successfully", floor)
}

// DeleteFloor handles DELETE /api/v1/floors/:id
// @Summary Delete a floor
// @Description Deletes (soft delete) an existing floor
// @Produce json
// @Param id path int true "Floor ID"
// @Success 204 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/floors/{id} [delete]
func (fc *FloorController) DeleteFloor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid floor ID", err.Error())
		return
	}

	err = fc.service.DeleteFloor(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Floor not found", err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
