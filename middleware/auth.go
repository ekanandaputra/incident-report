package middleware

import (
	"incident-report/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a placeholder for JWT-based authentication middleware
// This serves as a template for implementing JWT authentication in the future
//
// Typical JWT flow:
// 1. Client sends Authorization header: "Bearer <token>"
// 2. Middleware extracts and validates the token
// 3. Token claims are added to request context
// 4. Handler accesses claims from context
//
// Future implementation should:
// - Validate JWT signature
// - Check token expiration
// - Extract user claims from token
// - Store claims in gin.Context for handlers to access
//
// Example implementation structure:
// func AuthMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         // Extract token from Authorization header
//         authHeader := c.GetHeader("Authorization")
//         if authHeader == "" {
//             utils.ErrorResponse(c, http.StatusUnauthorized, "Missing Authorization Header", "")
//             c.Abort()
//             return
//         }
//
//         // Token format: "Bearer <token>"
//         parts := strings.Split(authHeader, " ")
//         if len(parts) != 2 || parts[0] != "Bearer" {
//             utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid Authorization Header Format", "")
//             c.Abort()
//             return
//         }
//
//         token := parts[1]
//
//         // Validate token and extract claims
//         // claims, err := ValidateToken(token)
//         // if err != nil {
//         //     utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid Token", err.Error())
//         //     c.Abort()
//         //     return
//         // }
//
//         // Store claims in context for handlers to use
//         // c.Set("user_claims", claims)
//
//         c.Next()
//     }
// }

// CORSMiddleware configures CORS headers for cross-origin requests
// Currently a template - implement based on your requirements
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		allowed := os.Getenv("CORS_ALLOW_ORIGINS")

		if strings.TrimSpace(allowed) == "" {
			allowed = "*"
		}

		// If wildcard is present, allow all origins and set credentials to false
		if strings.Contains(allowed, "*") {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
		} else {
			matched := false
			for _, o := range strings.Split(allowed, ",") {
				if strings.TrimSpace(o) == origin {
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
					matched = true
					break
				}
			}
			if !matched {
				first := strings.Split(allowed, ",")[0]
				c.Writer.Header().Set("Access-Control-Allow-Origin", strings.TrimSpace(first))
			}
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// RequestLoggingMiddleware logs incoming requests and outgoing responses
// Useful for debugging and monitoring API usage
// Future enhancement: integrate with logging library (e.g., zap, logrus)
func RequestLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request details
		// log.Printf("METHOD: %s, PATH: %s, IP: %s", c.Request.Method, c.Request.URL.Path, c.ClientIP())

		// Continue processing
		c.Next()

		// Log response details
		// log.Printf("STATUS: %d, METHOD: %s, PATH: %s", c.Writer.Status(), c.Request.Method, c.Request.URL.Path)
	}
}

// RateLimitMiddleware is a placeholder for implementing rate limiting
// Future implementation should:
// - Track requests per IP/user
// - Limit requests per time window
// - Return 429 Too Many Requests when limit exceeded
// - Use Redis or in-memory store for tracking
//
// Example using a simple counter:
// func RateLimitMiddleware(requestsPerSecond int) gin.HandlerFunc {
//     limiter := time.NewTicker(time.Second / time.Duration(requestsPerSecond))
//     return func(c *gin.Context) {
//         select {
//         case <-limiter.C:
//             c.Next()
//         default:
//             utils.ErrorResponse(c, http.StatusTooManyRequests, "Rate limit exceeded", "")
//             c.Abort()
//         }
//     }
// }

// ValidateContentTypeMiddleware ensures requests have proper Content-Type header
// Useful for POST/PUT/PATCH endpoints that expect JSON
func ValidateContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only validate for requests with body
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			contentType := c.GetHeader("Content-Type")
			if !strings.Contains(contentType, "application/json") {
				utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Content-Type", "Content-Type must be application/json")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
