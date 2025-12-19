package controllers

import (
	"net/http"
	"strconv"

	"incident-report/services"
	"incident-report/utils"

	"github.com/gin-gonic/gin"
)

// BuildingController handles HTTP requests for building operations
type BuildingController struct {
	service *services.BuildingService
}

// NewBuildingController creates a new instance of BuildingController
func NewBuildingController() *BuildingController {
	return &BuildingController{
		service: services.NewBuildingService(),
	}
}

// CreateBuilding handles POST /api/v1/buildings
// @Summary Create a new building
// @Description Creates a new building with the provided information
// @Accept json
// @Produce json
// @Param request body utils.CreateBuildingRequest true "Building details"
// @Success 201 {object} utils.BuildingResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/buildings [post]
func (bc *BuildingController) CreateBuilding(c *gin.Context) {
	var req utils.CreateBuildingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	building, err := bc.service.CreateBuilding(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create building", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Building created successfully", building)
}

// GetBuilding handles GET /api/v1/buildings/:id
// @Summary Get building by ID
// @Description Retrieves a specific building by its ID
// @Produce json
// @Param id path int true "Building ID"
// @Success 200 {object} utils.BuildingResponse
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/buildings/{id} [get]
func (bc *BuildingController) GetBuilding(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid building ID", err.Error())
		return
	}

	building, err := bc.service.GetBuildingByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Building not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Building retrieved successfully", building)
}

// GetAllBuildings handles GET /api/v1/buildings
// @Summary Get all buildings
// @Description Retrieves all buildings with pagination
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/buildings [get]
func (bc *BuildingController) GetAllBuildings(c *gin.Context) {
	var query utils.PaginationQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	// Apply same defaults as service to avoid division by zero in controller
	page := query.Page
	pageSize := query.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	buildings, total, err := bc.service.GetAllBuildings(page, pageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve buildings", err.Error())
		return
	}

	response := utils.PaginatedResponse{
		Data:      buildings,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	utils.SuccessResponse(c, http.StatusOK, "Buildings retrieved successfully", response)
}

// UpdateBuilding handles PUT /api/v1/buildings/:id
// @Summary Update a building
// @Description Updates an existing building
// @Accept json
// @Produce json
// @Param id path int true "Building ID"
// @Param request body utils.UpdateBuildingRequest true "Building details to update"
// @Success 200 {object} utils.BuildingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/buildings/{id} [put]
func (bc *BuildingController) UpdateBuilding(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid building ID", err.Error())
		return
	}

	var req utils.UpdateBuildingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	building, err := bc.service.UpdateBuilding(uint(id), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update building", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Building updated successfully", building)
}

// DeleteBuilding handles DELETE /api/v1/buildings/:id
// @Summary Delete a building
// @Description Deletes (soft delete) an existing building
// @Produce json
// @Param id path int true "Building ID"
// @Success 204 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/buildings/{id} [delete]
func (bc *BuildingController) DeleteBuilding(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid building ID", err.Error())
		return
	}

	err = bc.service.DeleteBuilding(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Building not found", err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
