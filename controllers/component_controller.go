package controllers

import (
	"net/http"
	"strconv"

	"incident-report/services"
	"incident-report/utils"

	"github.com/gin-gonic/gin"
)

// ComponentController handles HTTP requests for component operations
type ComponentController struct {
	service *services.ComponentService
}

// NewComponentController creates a new instance of ComponentController
func NewComponentController() *ComponentController {
	return &ComponentController{
		service: services.NewComponentService(),
	}
}

// CreateComponent handles POST /api/v1/components
// @Summary Create a new component
// @Description Creates a new component with the provided information
// @Accept json
// @Produce json
// @Param request body utils.CreateComponentRequest true "Component details"
// @Success 201 {object} utils.ComponentResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/components [post]
func (cc *ComponentController) CreateComponent(c *gin.Context) {
	var req utils.CreateComponentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	component, err := cc.service.CreateComponent(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create component", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Component created successfully", component)
}

// GetComponent handles GET /api/v1/components/:id
// @Summary Get component by ID
// @Description Retrieves a specific component by its ID
// @Produce json
// @Param id path int true "Component ID"
// @Success 200 {object} utils.ComponentResponse
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/components/{id} [get]
func (cc *ComponentController) GetComponent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid component ID", err.Error())
		return
	}

	component, err := cc.service.GetComponentByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Component not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Component retrieved successfully", component)
}

// GetComponentsByRoom handles GET /api/v1/rooms/:roomId/components
// @Summary Get all components in a room
// @Description Retrieves all components in a specific room with pagination
// @Produce json
// @Param roomId path int true "Room ID"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/rooms/{roomId}/components [get]
func (cc *ComponentController) GetComponentsByRoom(c *gin.Context) {
	roomID, err := strconv.ParseUint(c.Param("roomId"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid room ID", err.Error())
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

	components, total, err := cc.service.GetComponentsByRoomID(uint(roomID), page, pageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve components", err.Error())
		return
	}

	response := utils.PaginatedResponse{
		Data:      components,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	utils.SuccessResponse(c, http.StatusOK, "Components retrieved successfully", response)
}

// GetComponentsByCategory handles GET /api/v1/component-categories/:categoryId/components
// @Summary Get all components in a category
// @Description Retrieves all components in a specific category with pagination
// @Produce json
// @Param categoryId path int true "Category ID"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/component-categories/{categoryId}/components [get]
func (cc *ComponentController) GetComponentsByCategory(c *gin.Context) {
	categoryID, err := strconv.ParseUint(c.Param("categoryId"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err.Error())
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

	components, total, err := cc.service.GetComponentsByCategoryID(uint(categoryID), page, pageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve components", err.Error())
		return
	}

	response := utils.PaginatedResponse{
		Data:      components,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	utils.SuccessResponse(c, http.StatusOK, "Components retrieved successfully", response)
}

// UpdateComponent handles PUT /api/v1/components/:id
// @Summary Update a component
// @Description Updates an existing component
// @Accept json
// @Produce json
// @Param id path int true "Component ID"
// @Param request body utils.UpdateComponentRequest true "Component details to update"
// @Success 200 {object} utils.ComponentResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/components/{id} [put]
func (cc *ComponentController) UpdateComponent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid component ID", err.Error())
		return
	}

	var req utils.UpdateComponentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	component, err := cc.service.UpdateComponent(uint(id), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update component", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Component updated successfully", component)
}

// DeleteComponent handles DELETE /api/v1/components/:id
// @Summary Delete a component
// @Description Deletes (soft delete) an existing component
// @Produce json
// @Param id path int true "Component ID"
// @Success 204 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/components/{id} [delete]
func (cc *ComponentController) DeleteComponent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid component ID", err.Error())
		return
	}

	err = cc.service.DeleteComponent(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Component not found", err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
