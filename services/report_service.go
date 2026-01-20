package services

import (
	"errors"
	"incident-report/config"
	"incident-report/models"
	"incident-report/utils"

	"gorm.io/gorm"
)

// ReportService handles all report-related business logic
type ReportService struct{}

// NewReportService creates a new instance of ReportService
func NewReportService() *ReportService {
	return &ReportService{}
}

// CreateReport creates a new report in the database
func (rs *ReportService) CreateReport(req *utils.CreateReportRequest) (*utils.ReportResponse, error) {
	// Validate input
	if req.Name == "" || req.RoomID == 0 || req.ComponentID == 0 {
		return nil, errors.New("name, room_id, and component_id are required")
	}

	// Create report model instance
	report := models.Report{
		Name:        req.Name,
		RoomID:      req.RoomID,
		UserID:      req.UserID,
		ComponentID: req.ComponentID,
		Status:      models.ReportStatus(req.Status),
	}

	// Save to database
	result := config.DB.Create(&report)
	if result.Error != nil {
		return nil, result.Error
	}

	// Return report response DTO
	return &utils.ReportResponse{
		ID:          report.ID,
		Name:        report.Name,
		RoomID:      report.RoomID,
		UserID:      report.UserID,
		ComponentID: report.ComponentID,
		Status:      string(report.Status),
		CreatedAt:   report.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   report.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// GetReportByID retrieves a report by their ID
func (rs *ReportService) GetReportByID(id uint) (*utils.ReportResponse, error) {
	var report models.Report

	// Query database for report with given ID
	result := config.DB.First(&report, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("report not found")
		}
		return nil, result.Error
	}

	return &utils.ReportResponse{
		ID:          report.ID,
		Name:        report.Name,
		RoomID:      report.RoomID,
		UserID:      report.UserID,
		ComponentID: report.ComponentID,
		Status:      string(report.Status),
		CreatedAt:   report.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   report.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// GetAllReports retrieves all reports with pagination support
func (rs *ReportService) GetAllReports(page int, pageSize int) ([]utils.ReportResponse, int64, error) {
	// Set default pagination values
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	var reports []models.Report
	var total int64

	// Count total records
	if err := config.DB.Model(&models.Report{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset for pagination
	offset := (page - 1) * pageSize

	// Fetch paginated results
	result := config.DB.Offset(offset).Limit(pageSize).Find(&reports)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	// Convert to response DTOs
	var responses []utils.ReportResponse
	for _, report := range reports {
		responses = append(responses, utils.ReportResponse{
			ID:          report.ID,
			Name:        report.Name,
			RoomID:      report.RoomID,
			UserID:      report.UserID,
			ComponentID: report.ComponentID,
			Status:      string(report.Status),
			CreatedAt:   report.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   report.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return responses, total, nil
}

// UpdateReport updates an existing report's information
func (rs *ReportService) UpdateReport(id uint, req *utils.UpdateReportRequest) (*utils.ReportResponse, error) {
	// Find report first
	var report models.Report
	result := config.DB.First(&report, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("report not found")
		}
		return nil, result.Error
	}

	// Update only provided fields
	if req.Name != "" {
		report.Name = req.Name
	}
	if req.RoomID != 0 {
		report.RoomID = req.RoomID
	}
	if req.UserID != nil {
		report.UserID = req.UserID
	}
	if req.ComponentID != 0 {
		report.ComponentID = req.ComponentID
	}
	if req.Status != "" {
		report.Status = models.ReportStatus(req.Status)
	}

	// Save changes to database
	if err := config.DB.Save(&report).Error; err != nil {
		return nil, err
	}

	return &utils.ReportResponse{
		ID:          report.ID,
		Name:        report.Name,
		RoomID:      report.RoomID,
		UserID:      report.UserID,
		ComponentID: report.ComponentID,
		Status:      string(report.Status),
		CreatedAt:   report.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   report.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// DeleteReport deletes a report from the database
func (rs *ReportService) DeleteReport(id uint) error {
	// Find report first to ensure it exists
	var report models.Report
	result := config.DB.First(&report, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("report not found")
		}
		return result.Error
	}

	// Perform hard delete
	result = config.DB.Unscoped().Delete(&report)
	return result.Error
}

// AssignUserToReport assigns a user to an existing report
func (rs *ReportService) AssignUserToReport(reportID uint, userID uint) (*utils.ReportResponse, error) {
	// Find report first
	var report models.Report
	result := config.DB.First(&report, reportID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("report not found")
		}
		return nil, result.Error
	}

	// Check if user exists
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Assign user to report
	report.UserID = &userID

	// Save changes to database
	if err := config.DB.Save(&report).Error; err != nil {
		return nil, err
	}

	return &utils.ReportResponse{
		ID:          report.ID,
		Name:        report.Name,
		RoomID:      report.RoomID,
		UserID:      report.UserID,
		ComponentID: report.ComponentID,
		Status:      string(report.Status),
		CreatedAt:   report.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   report.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
