package controllers

import (
	"net/http"
	"strconv"

	"incident-report/services"
	"incident-report/utils"

	"github.com/gin-gonic/gin"
)

// ComponentCategoryController handles HTTP requests for component category operations
type ComponentCategoryController struct {
	service *services.ComponentCategoryService
}

// NewComponentCategoryController creates a new instance of ComponentCategoryController
func NewComponentCategoryController() *ComponentCategoryController {
	return &ComponentCategoryController{
		service: services.NewComponentCategoryService(),
	}
}

// CreateComponentCategory handles POST /api/v1/component-categories
// @Summary Create a new component category
// @Description Creates a new component category with the provided information
// @Accept json
// @Produce json
// @Param request body utils.CreateComponentCategoryRequest true "Category details"
// @Success 201 {object} utils.ComponentCategoryResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/component-categories [post]
func (ccc *ComponentCategoryController) CreateComponentCategory(c *gin.Context) {
	var req utils.CreateComponentCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	category, err := ccc.service.CreateComponentCategory(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create component category", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Component category created successfully", category)
}

// GetComponentCategory handles GET /api/v1/component-categories/:id
// @Summary Get component category by ID
// @Description Retrieves a specific component category by its ID
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} utils.ComponentCategoryResponse
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/component-categories/{id} [get]
func (ccc *ComponentCategoryController) GetComponentCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	category, err := ccc.service.GetComponentCategoryByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Component category not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Component category retrieved successfully", category)
}

// GetAllComponentCategories handles GET /api/v1/component-categories
// @Summary Get all component categories
// @Description Retrieves all component categories with pagination
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/component-categories [get]
func (ccc *ComponentCategoryController) GetAllComponentCategories(c *gin.Context) {
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

	categories, total, err := ccc.service.GetAllComponentCategories(page, pageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve component categories", err.Error())
		return
	}

	response := utils.PaginatedResponse{
		Data:      categories,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	utils.SuccessResponse(c, http.StatusOK, "Component categories retrieved successfully", response)
}

// UpdateComponentCategory handles PUT /api/v1/component-categories/:id
// @Summary Update a component category
// @Description Updates an existing component category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param request body utils.UpdateComponentCategoryRequest true "Category details to update"
// @Success 200 {object} utils.ComponentCategoryResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/component-categories/{id} [put]
func (ccc *ComponentCategoryController) UpdateComponentCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	var req utils.UpdateComponentCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	category, err := ccc.service.UpdateComponentCategory(uint(id), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update component category", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Component category updated successfully", category)
}

// DeleteComponentCategory handles DELETE /api/v1/component-categories/:id
// @Summary Delete a component category
// @Description Deletes (soft delete) an existing component category
// @Produce json
// @Param id path int true "Category ID"
// @Success 204 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/component-categories/{id} [delete]
func (ccc *ComponentCategoryController) DeleteComponentCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	err = ccc.service.DeleteComponentCategory(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Component category not found", err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
