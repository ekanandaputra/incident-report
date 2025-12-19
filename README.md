# Incident Report RESTful API

A production-ready RESTful API built with Go, featuring MySQL database integration, clean architecture principles, and comprehensive API documentation.

## 📋 Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Usage Examples](#usage-examples)
- [Architecture Overview](#architecture-overview)
- [Future Enhancements](#future-enhancements)

## ✨ Features

- **RESTful API** with versioned endpoints (`/api/v1`)
- **CRUD Operations** for User management
- **Pagination Support** for list endpoints
- **MySQL Database** with GORM ORM
- **Soft Deletes** for data preservation
- **Request Validation** with detailed error messages
- **Middleware Support** for error handling
- **Environment Configuration** with `.env` file
- **Clean Architecture** with separation of concerns
- **Production-Ready Code** with best practices
- **Auto-Migration** for database schema

## 🛠️ Technology Stack

- **Go** 1.21+ - Programming language
- **Gin** - HTTP web framework for routing and middleware
- **GORM** - ORM library for database operations
- **MySQL** - Relational database
- **godotenv** - Environment variable management
- **Go SQL Driver for MySQL** - Database driver

## 📁 Project Structure

```
incident-report/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── database.go             # Database configuration and initialization
├── models/
│   └── user.go                 # User entity model
├── controllers/
│   └── user_controller.go      # HTTP request handlers
├── services/
│   └── user_service.go         # Business logic layer
├── routes/
│   └── routes.go               # API routing configuration
├── middleware/
│   └── error_handler.go        # Error handling middleware
├── utils/
│   ├── response.go             # Response utility functions
│   └── dto.go                  # Data Transfer Objects (DTOs)
├── .env                        # Environment variables
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
├── Makefile                    # Build automation
└── README.md                   # This file
```

## 📦 Prerequisites

- Go 1.21 or higher
- MySQL 5.7 or higher
- Git (optional, for version control)

## 🚀 Installation

1. **Clone or navigate to the project directory:**
   ```bash
   cd incident-report
   ```

2. **Download Go dependencies:**
   ```bash
   go mod download
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

## ⚙️ Configuration

1. **Create a MySQL database:**
   ```sql
   CREATE DATABASE incident_report CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   ```

2. **Update the `.env` file** with your database credentials:
   ```env
   # Server Configuration
   SERVER_HOST=localhost
   SERVER_PORT=8080
   ENVIRONMENT=development

   # Database Configuration
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=your_password
   DB_NAME=incident_report
   ```

3. **Save and verify** the `.env` file is in the project root directory.

## ▶️ Running the Application

### Using Go directly:
```bash
go run cmd/main.go
```

### Using Make (if available):
```bash
make run
```

### Using build binary:
```bash
go build -o incident-report cmd/main.go
./incident-report
```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### Health Check
- **GET** `/api/v1/health`
  - Returns server health status
  - Response: `200 OK`

### User Management

#### Create User
- **POST** `/api/v1/users`
- **Request Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```
- **Response:** `201 Created`
  ```json
  {
    "success": true,
    "message": "User created successfully",
    "data": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com"
    }
  }
  ```

#### Get All Users (with Pagination)
- **GET** `/api/v1/users?page=1&page_size=10`
- **Query Parameters:**
  - `page` (optional, default: 1) - Page number
  - `page_size` (optional, default: 10, max: 100) - Records per page
- **Response:** `200 OK`
  ```json
  {
    "success": true,
    "message": "Users retrieved successfully",
    "data": {
      "data": [
        {
          "id": 1,
          "name": "John Doe",
          "email": "john@example.com"
        }
      ],
      "page": 1,
      "page_size": 10,
      "total": 1,
      "total_page": 1
    }
  }
  ```

#### Get User by ID
- **GET** `/api/v1/users/:id`
- **Response:** `200 OK`
  ```json
  {
    "success": true,
    "message": "User retrieved successfully",
    "data": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com"
    }
  }
  ```

#### Update User
- **PUT** `/api/v1/users/:id`
- **Request Body (partial):**
  ```json
  {
    "name": "Jane Doe",
    "email": "jane@example.com"
  }
  ```
- **Response:** `200 OK`
  ```json
  {
    "success": true,
    "message": "User updated successfully",
    "data": {
      "id": 1,
      "name": "Jane Doe",
      "email": "jane@example.com"
    }
  }
  ```

#### Delete User
- **DELETE** `/api/v1/users/:id`
- **Response:** `200 OK`
  ```json
  {
    "success": true,
    "message": "User deleted successfully"
  }
  ```

## 💡 Usage Examples

### Using cURL

**Create a user:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

**Get all users:**
```bash
curl http://localhost:8080/api/v1/users?page=1&page_size=10
```

**Get specific user:**
```bash
curl http://localhost:8080/api/v1/users/1
```

**Update user:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com"
  }'
```

**Delete user:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

### Using Postman

1. Import the API endpoints
2. Set the base URL to `http://localhost:8080`
3. Create requests for each endpoint with appropriate HTTP methods
4. Use the examples above for request bodies and parameters

## 🏗️ Architecture Overview

The project follows **Clean Architecture** principles with clear separation of concerns:

### Layers

1. **HTTP Layer** (`controllers/`)
   - Handles HTTP requests and responses
   - Validates input parameters
   - Calls service layer for business logic
   - Returns appropriate HTTP status codes

2. **Service Layer** (`services/`)
   - Contains core business logic
   - Performs validation
   - Orchestrates database operations
   - Returns DTOs to controllers

3. **Data Layer** (`models/`, `config/`)
   - Database models with GORM tags
   - Database connection management
   - Auto-migration logic

4. **Support Layers**
   - **Routes** (`routes/`) - API routing and versioning
   - **Middleware** (`middleware/`) - Cross-cutting concerns
   - **Utils** (`utils/`) - Response helpers and DTOs

### Data Flow

```
HTTP Request
    ↓
Gin Router
    ↓
Controller (HTTP Handler)
    ↓
Service (Business Logic)
    ↓
Model (Database Entity)
    ↓
GORM → MySQL
    ↓
Response DTO
    ↓
HTTP Response
```

## 🔐 Future Enhancements

### Authentication & Authorization
- JWT token-based authentication
- Role-based access control (RBAC)
- Password hashing with bcrypt
- Token refresh mechanism

### Advanced Features
- Request logging middleware
- Rate limiting
- CORS configuration
- API versioning strategy
- Comprehensive error codes
- Request/Response caching
- Database transaction support

### Testing
- Unit tests for services
- Integration tests for controllers
- Mock database tests
- API endpoint tests with testify

### DevOps & Deployment
- Docker containerization
- Docker Compose for local development
- GitHub Actions CI/CD pipeline
- Database migration scripts
- Environment-specific configurations

### Monitoring & Analytics
- Request/Response logging
- Performance metrics
- Error tracking and reporting
- Health check endpoints

### Code Quality
- Code coverage analysis
- Linting with golangci-lint
- Formatting with gofmt
- Documentation generation

## 📝 Best Practices Implemented

✅ **Clean Architecture** - Separation of concerns with distinct layers
✅ **SOLID Principles** - Single responsibility, Open/closed, etc.
✅ **Dependency Injection** - Services injected into controllers
✅ **Error Handling** - Comprehensive error handling with meaningful messages
✅ **Validation** - Input validation at controller level
✅ **DTOs** - Request/Response objects separate from models
✅ **Soft Deletes** - Data preservation with logical deletes
✅ **Pagination** - Efficient data retrieval with limits
✅ **Environment Configuration** - Externalized configuration with .env
✅ **Logging** - Application startup and important events
✅ **HTTP Standards** - Proper status codes and REST conventions
✅ **Code Comments** - Clear documentation of important logic

## 🐛 Troubleshooting

### Database Connection Error
**Error:** `Failed to connect to database`
- **Solution:** 
  1. Verify MySQL is running
  2. Check database credentials in `.env`
  3. Ensure the database exists
  4. Verify network connectivity

### Port Already in Use
**Error:** `listen tcp :8080: bind: address already in use`
- **Solution:**
  1. Change `SERVER_PORT` in `.env`
  2. Or kill the process using the port

### Missing Dependencies
**Error:** `cannot find package`
- **Solution:**
  ```bash
  go mod tidy
  go mod download
  ```

## 📄 License

This project is provided as-is for educational and development purposes.

## 🤝 Contributing

Feel free to extend this project with additional features, improvements, and enhancements following the established architecture and best practices.

---

**Happy Coding! 🚀**
