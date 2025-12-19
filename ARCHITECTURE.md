# Architecture Documentation

## Overview

This document provides a comprehensive overview of the Incident Report RESTful API architecture, design patterns, and technical decisions.

## Architecture Pattern: Clean Architecture

The project follows **Clean Architecture** principles which emphasize:
- Independence of frameworks
- Testability
- Independence of UI
- Independence of database
- Independence of any external agency

### Architecture Layers

```
┌─────────────────────────────────────────────────────────┐
│                    Presentation Layer                    │
│  (HTTP Handlers, Request/Response Formatting)            │
│  - controllers/user_controller.go                        │
├─────────────────────────────────────────────────────────┤
│                    Application Layer                      │
│  (Business Logic, Validation, Orchestration)             │
│  - services/user_service.go                              │
├─────────────────────────────────────────────────────────┤
│                    Data Access Layer                      │
│  (Database Models, Query Logic)                          │
│  - models/user.go                                        │
│  - config/database.go                                    │
├─────────────────────────────────────────────────────────┤
│                    Infrastructure                        │
│  (MySQL Database)                                        │
└─────────────────────────────────────────────────────────┘
```

## Directory Structure Explained

```
incident-report/
│
├── cmd/
│   └── main.go
│       Purpose: Application entry point
│       Responsibility:
│       - Load environment variables
│       - Initialize database
│       - Configure Gin router
│       - Start HTTP server
│       - Handle graceful shutdown
│
├── config/
│   ├── database.go
│   │   Purpose: Database connection and migration
│   │   Responsibility:
│   │   - MySQL connection setup with GORM
│   │   - Auto-migration for models
│   │   - Connection pooling configuration
│   │   - Graceful shutdown
│   │
│   └── database_queries.go
│       Purpose: Reference guide for GORM queries
│       Responsibility:
│       - Example queries and patterns
│       - Best practices documentation
│
├── models/
│   └── user.go
│       Purpose: Database entity definition
│       Responsibility:
│       - Define User struct with database tags
│       - Specify column constraints (unique, index, etc.)
│       - Support for soft deletes with deleted_at
│       - Timestamps (created_at, updated_at)
│
├── controllers/
│   └── user_controller.go
│       Purpose: HTTP request handlers
│       Responsibility:
│       - Receive HTTP requests
│       - Parse and validate request parameters
│       - Call service layer for business logic
│       - Format responses
│       - Return appropriate HTTP status codes
│       - Error handling at HTTP level
│
├── services/
│   └── user_service.go
│       Purpose: Business logic layer
│       Responsibility:
│       - Implement core business rules
│       - Perform validation logic
│       - Orchestrate database operations
│       - Return DTOs (Data Transfer Objects)
│       - Handle service-level errors
│       - Keep controllers thin and focused
│
├── routes/
│   └── routes.go
│       Purpose: API routing configuration
│       Responsibility:
│       - Define API versioning (/api/v1)
│       - Register all route handlers
│       - Apply middleware to routes
│       - Organize endpoints logically
│       - Document endpoint structure
│
├── middleware/
│   ├── error_handler.go
│   │   Purpose: Error handling and recovery
│   │   Responsibility:
│   │   - Recover from panics
│   │   - Format error responses
│   │   - Log errors
│   │
│   └── auth.go
│       Purpose: Authentication and authorization (template)
│       Responsibility:
│       - JWT token validation (future)
│       - CORS configuration
│       - Request logging
│       - Rate limiting template
│       - Content-Type validation
│
├── utils/
│   ├── response.go
│   │   Purpose: Response formatting utilities
│   │   Responsibility:
│   │   - Standardized success responses
│   │   - Standardized error responses
│   │   - HTTP status code mapping
│   │
│   └── dto.go
│       Purpose: Data Transfer Objects
│       Responsibility:
│       - CreateUserRequest - input validation
│       - UpdateUserRequest - partial update validation
│       - UserResponse - output formatting
│       - PaginationQuery - query parameter parsing
│       - PaginatedResponse - list response format
│
├── .env
│   Purpose: Environment configuration
│   Contains:
│   - Database credentials
│   - Server settings
│   - Environment mode (dev/prod)
│
├── go.mod
│   Purpose: Go module definition
│   Contains:
│   - Module name (incident-report)
│   - Go version
│   - Direct dependencies
│
├── Makefile
│   Purpose: Build automation and task management
│   Commands: build, run, clean, deps, fmt, test, dev
│
├── .gitignore
│   Purpose: Git version control exclusions
│   Excludes: binaries, logs, .env, IDE files, etc.
│
├── README.md
│   Purpose: Project documentation
│   Contains: Features, setup, API endpoints, examples
│
└── API_TESTING_GUIDE.md
    Purpose: Comprehensive testing documentation
    Contains: cURL examples, Postman guide, test scripts
```

## Data Flow

### Create User Request Flow

```
1. HTTP Request (POST /api/v1/users)
   ↓
2. Gin Router (routes.go)
   ↓
3. UserController.CreateUser (controllers/user_controller.go)
   - Parse request body
   - Validate input
   ↓
4. UserService.CreateUser (services/user_service.go)
   - Validate business rules
   - Create model instance
   ↓
5. GORM (config/database.go)
   - Execute INSERT query
   - Handle database errors
   ↓
6. MySQL Database
   - Store user record
   ↓
7. Convert to DTO (utils/dto.go)
   ↓
8. Format Response (utils/response.go)
   ↓
9. HTTP Response (JSON with 201 Created)
```

### Get All Users Request Flow

```
1. HTTP Request (GET /api/v1/users?page=1&page_size=10)
   ↓
2. Gin Router (routes.go)
   ↓
3. UserController.GetAllUsers
   - Parse query parameters
   - Validate pagination
   ↓
4. UserService.GetAllUsers
   - Count total records
   - Calculate offset
   - Fetch paginated results
   ↓
5. GORM
   - Execute SELECT with LIMIT/OFFSET
   - Get COUNT(*)
   ↓
6. MySQL Database
   - Return records and count
   ↓
7. Convert to DTOs and PaginatedResponse
   ↓
8. Format Response
   ↓
9. HTTP Response (JSON with 200 OK)
```

## Design Patterns Used

### 1. Dependency Injection
```go
// Constructor takes dependencies
func NewUserController(userService *UserService) *UserController {
    return &UserController{userService: userService}
}

// Usage in routes
userService := services.NewUserService()
userController := controllers.NewUserController(userService)
```

**Benefits:**
- Easier to test (mock dependencies)
- Loose coupling
- Flexible configuration
- Follows SOLID principles

### 2. Data Transfer Objects (DTOs)
```go
// Separate input/output objects from database models
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

type UserResponse struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

**Benefits:**
- Decouples API contract from database schema
- Enables request validation
- Controls what data is exposed
- Enables transformation logic

### 3. Service Layer Pattern
```go
// Business logic separated from HTTP handling
type UserService struct {}

func (us *UserService) CreateUser(req *CreateUserRequest) (*UserResponse, error) {
    // Validation and business logic
    // Database operations via GORM
    // Error handling
}
```

**Benefits:**
- Testable business logic
- Reusable across controllers
- Centralized business rules
- Easier maintenance

### 4. Repository Pattern (via GORM)
```go
// GORM acts as repository
DB.Create(&user)        // Create
DB.First(&user, id)     // Read
DB.Save(&user)          // Update
DB.Delete(&user)        // Delete
```

**Benefits:**
- Abstract database operations
- Easy to switch databases
- Consistent query interface
- Transaction support

### 5. Middleware Pattern
```go
// Cross-cutting concerns
router.Use(middleware.ErrorHandlerMiddleware())

// Applied to route groups
v1 := router.Group("/api/v1")
v1.GET("/users", userController.GetAllUsers)
```

**Benefits:**
- Separation of concerns
- Reusable across routes
- Clean controller logic
- Global exception handling

## SOLID Principles Implementation

### Single Responsibility Principle
```
- Models: Define database schema
- Services: Implement business logic
- Controllers: Handle HTTP layer
- Routes: Configure routing
```

### Open/Closed Principle
```
- Services are open for extension via inheritance
- Middleware can be added without modifying existing code
```

### Liskov Substitution Principle
```
- Services implement consistent interfaces
- Easy to swap implementations for testing
```

### Interface Segregation Principle
```
- Controllers depend on specific service methods
- Not forced to depend on unused methods
```

### Dependency Inversion Principle
```
- Controllers depend on service abstractions
- Not concrete implementations
- Dependency injection via constructors
```

## Error Handling Strategy

### At Controller Level
```go
// HTTP-level errors
utils.ErrorResponse(c, http.StatusBadRequest, "message", "error")
```

### At Service Level
```go
// Business logic errors
return nil, errors.New("validation failed")
```

### At Middleware Level
```go
// Panic recovery and cross-cutting errors
defer func() {
    if err := recover(); err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "error", "")
    }
}()
```

## Response Format

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { /* actual data */ }
}
```

### Error Response
```json
{
  "success": false,
  "message": "User-friendly message",
  "error": "Detailed error information"
}
```

### Paginated Response
```json
{
  "success": true,
  "message": "Users retrieved",
  "data": {
    "data": [ /* records */ ],
    "page": 1,
    "page_size": 10,
    "total": 100,
    "total_page": 10
  }
}
```

## HTTP Status Codes

| Code | Meaning | Use Case |
|------|---------|----------|
| 200 | OK | Successful GET, PUT, DELETE |
| 201 | Created | Successful POST |
| 204 | No Content | Successful DELETE with no response |
| 400 | Bad Request | Validation error, missing fields |
| 401 | Unauthorized | Missing/invalid authentication |
| 404 | Not Found | Resource doesn't exist |
| 409 | Conflict | Duplicate email, constraint violation |
| 500 | Internal Server | Database error, panic |

## Database Schema

### Users Table
```sql
CREATE TABLE users (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
```

## Configuration Management

### Environment Variables (.env)
```env
SERVER_HOST=localhost      # Server host
SERVER_PORT=8080          # Server port
ENVIRONMENT=development   # dev or production
DB_HOST=localhost         # Database host
DB_PORT=3306             # Database port
DB_USER=root             # Database user
DB_PASSWORD=password     # Database password
DB_NAME=incident_report  # Database name
```

### Loading Configuration
```go
if err := godotenv.Load(); err != nil {
    log.Println("Warning: .env file not found")
}

// Accessed via os.Getenv()
```

## Security Considerations

### Current Implementation
✅ Input validation on all endpoints
✅ SQL injection prevention (parameterized queries via GORM)
✅ Email format validation
✅ Soft deletes (data preservation)
✅ Error message obfuscation

### Future Enhancements
🔒 JWT authentication
🔒 CORS configuration
🔒 Rate limiting
🔒 Request logging
🔒 HTTPS support
🔒 Password hashing (bcrypt)
🔒 Database encryption
🔒 Audit logging

## Performance Optimization

### Current Features
✅ Database indexing (email, deleted_at)
✅ Pagination for list endpoints
✅ Selective column queries
✅ Connection pooling (GORM default)

### Future Enhancements
📊 Caching (Redis)
📊 Query optimization
📊 Bulk operations
📊 Database replication
📊 Load balancing
📊 API monitoring

## Testing Strategy

### Current Setup (Ready for Implementation)
- Unit tests for services
- Integration tests for controllers
- Mock database for testing
- API endpoint testing

### Example Test Structure
```go
// services/user_service_test.go
func TestCreateUser(t *testing.T) {
    // Arrange
    req := &utils.CreateUserRequest{...}
    
    // Act
    resp, err := service.CreateUser(req)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "John", resp.Name)
}
```

## Deployment Considerations

### Docker Setup (Future)
```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o incident-report cmd/main.go
EXPOSE 8080
CMD ["./incident-report"]
```

### Environment Differences
- **Development**: Debug logging, hot reload, local database
- **Production**: Release mode, structured logging, remote database

## Monitoring and Logging

### Current Implementation
- Server startup logging
- API endpoint documentation
- Database connection logging

### Future Enhancements
📝 Request/response logging middleware
📝 Error tracking and reporting
📝 Performance metrics
📝 Database query logging
📝 Health check endpoint
📝 Metrics export (Prometheus)

## API Versioning Strategy

### Current Implementation
```
/api/v1/users - Version 1 endpoints
```

### Future Versioning
```
/api/v1/*     - Stable version
/api/v2/*     - New features (non-breaking)
/api/beta/*   - Experimental features
```

## Scalability Considerations

### Horizontal Scaling
- Stateless API design ✅
- External database ✅
- Session management ready for Redis
- Load balancer compatible ✅

### Vertical Scaling
- Connection pooling ✅
- Database indexing ✅
- Query optimization ready
- Caching ready for implementation

## Summary

This architecture provides:
- ✅ Clean, maintainable code structure
- ✅ Separation of concerns
- ✅ Testability
- ✅ Scalability potential
- ✅ Security foundation
- ✅ Production readiness
- ✅ Future enhancement flexibility

The project is designed to grow with your needs while maintaining code quality and architectural integrity.

---

**Architecture Last Updated:** December 15, 2025
