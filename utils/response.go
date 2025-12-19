package utils

import (
	"github.com/gin-gonic/gin"
)

// ResponseData represents a standard API response structure
type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse returns a success response with data
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, ResponseData{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse returns an error response with an error message
func ErrorResponse(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, ResponseData{
		Success: false,
		Message: message,
		Error:   err,
	})
}
