package middleware

import (
	"incident-report/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware recovers from panics and returns error response
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.ErrorResponse(
					c,
					http.StatusInternalServerError,
					"Internal Server Error",
					"An unexpected error occurred",
				)
			}
		}()
		c.Next()
	}
}

// ValidationErrorHandler returns validation error response
func ValidationErrorHandler(c *gin.Context, fieldName string, message string) {
	utils.ErrorResponse(
		c,
		http.StatusBadRequest,
		"Validation Error",
		fieldName+" - "+message,
	)
}
