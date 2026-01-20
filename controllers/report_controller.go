package controllers

import (
	"incident-report/services"
	"incident-report/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReportController handles HTTP requests for report operations
type ReportController struct {
	reportService *services.ReportService
}

// NewReportController creates a new instance of ReportController with dependency injection
func NewReportController(reportService *services.ReportService) *ReportController {
	return &ReportController{
		reportService: reportService,
	}
}

// CreateReport handles POST /api/v1/reports request to create a new report
// @param c *gin.Context
// Request body: CreateReportRequest (name, room_id, user_id, status)
// Response: ReportResponse with HTTP 201 Created
func (rc *ReportController) CreateReport(c *gin.Context) {
	var req utils.CreateReportRequest

	// Bind and validate request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Call service to create report
	report, err := rc.reportService.CreateReport(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create report", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Report created successfully", report)
}

// GetReport handles GET /api/v1/reports/:id request to retrieve a specific report
// @param c *gin.Context with :id parameter
// Response: ReportResponse with HTTP 200 OK
func (rc *ReportController) GetReport(c *gin.Context) {
	// Extract report ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid report ID", "ID must be a valid number")
		return
	}

	// Call service to fetch report
	report, err := rc.reportService.GetReportByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Report not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Report retrieved successfully", report)
}

// GetAllReports handles GET /api/v1/reports request to retrieve all reports with pagination
// @param c *gin.Context with optional query parameters: page, page_size
// Response: PaginatedResponse with array of reports and HTTP 200 OK
func (rc *ReportController) GetAllReports(c *gin.Context) {
	var pagination utils.PaginationQuery

	// Bind query parameters with default values
	if err := c.ShouldBindQuery(&pagination); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	// Set defaults if not provided
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}

	// Call service to fetch paginated reports
	reports, total, err := rc.reportService.GetAllReports(pagination.Page, pagination.PageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch reports", err.Error())
		return
	}

	// Calculate total pages
	totalPage := (int(total) + pagination.PageSize - 1) / pagination.PageSize

	// Create paginated response
	response := utils.PaginatedResponse{
		Data:      reports,
		Page:      pagination.Page,
		PageSize:  pagination.PageSize,
		Total:     total,
		TotalPage: totalPage,
	}

	utils.SuccessResponse(c, http.StatusOK, "Reports retrieved successfully", response)
}

// UpdateReport handles PUT /api/v1/reports/:id request to update a report
// @param c *gin.Context with :id parameter
// Request body: UpdateReportRequest (partial fields)
// Response: ReportResponse with HTTP 200 OK
func (rc *ReportController) UpdateReport(c *gin.Context) {
	// Extract report ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid report ID", "ID must be a valid number")
		return
	}

	var req utils.UpdateReportRequest

	// Bind and validate request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Call service to update report
	report, err := rc.reportService.UpdateReport(uint(id), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update report", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Report updated successfully", report)
}

// DeleteReport handles DELETE /api/v1/reports/:id request to delete a report
// @param c *gin.Context with :id parameter
// Response: HTTP 204 No Content on success
func (rc *ReportController) DeleteReport(c *gin.Context) {
	// Extract report ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid report ID", "ID must be a valid number")
		return
	}

	// Call service to delete report
	err = rc.reportService.DeleteReport(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to delete report", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Report deleted successfully", nil)
}

// AssignUserToReport handles PUT /api/v1/reports/:id/assign-user request to assign a user to a report
// @param c *gin.Context with :id parameter
// Request body: AssignUserRequest (user_id)
// Response: ReportResponse with HTTP 200 OK
func (rc *ReportController) AssignUserToReport(c *gin.Context) {
	// Extract report ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid report ID", "ID must be a valid number")
		return
	}

	var req utils.AssignUserRequest

	// Bind and validate request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Call service to assign user to report
	report, err := rc.reportService.AssignUserToReport(uint(id), req.UserID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to assign user to report", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User assigned to report successfully", report)
}
