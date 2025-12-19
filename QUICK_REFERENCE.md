# Quick Reference Guide

## 🚀 Getting Started (5 Minutes)

### 1. Prerequisites
```bash
# Check Go installation
go version    # Should be 1.21+

# Check MySQL is running
mysql --version
```

### 2. Setup Database
```bash
# Connect to MySQL
mysql -u root -p

# Create database
CREATE DATABASE incident_report;

# Exit MySQL
exit
```

### 3. Configure Environment
```bash
# Edit .env file
# Update DB credentials to match your setup
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=incident_report
```

### 4. Install & Run
```bash
# Download dependencies
go mod tidy

# Run the application
go run cmd/main.go

# Or use Make
make run
```

### 5. Test API
```bash
# Create a user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@example.com"}'

# Get all users
curl http://localhost:8080/api/v1/users

# Health check
curl http://localhost:8080/api/v1/health
```

---

## 📡 API Endpoints Quick Reference

### User Operations
```
POST   /api/v1/users              Create user
GET    /api/v1/users              Get all (paginated)
GET    /api/v1/users/:id          Get by ID
PUT    /api/v1/users/:id          Update user
DELETE /api/v1/users/:id          Delete user
GET    /api/v1/health             Health check
```

### Query Parameters
```
?page=1              # Page number (default: 1)
?page_size=10        # Records per page (default: 10, max: 100)
```

---

## 🗂️ Project Structure at a Glance

```
cmd/main.go           → Application startup
config/database.go    → Database configuration
models/user.go        → User entity
services/user_service.go → Business logic
controllers/user_controller.go → HTTP handlers
routes/routes.go      → API routing
middleware/           → Cross-cutting concerns
utils/                → Response helpers & DTOs
.env                  → Configuration
```

---

## 🔧 Make Commands

```bash
make build    # Compile binary
make run      # Compile and run
make dev      # Run in development mode
make clean    # Remove binaries
make deps     # Download dependencies
make fmt      # Format code
make test     # Run tests
make help     # Show all commands
```

---

## 📝 Create User Request Example

### Request
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

### Response (201 Created)
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

---

## 📋 Get All Users with Pagination

### Request
```bash
curl "http://localhost:8080/api/v1/users?page=1&page_size=10"
```

### Response (200 OK)
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

---

## ❌ Error Response Example

### Request (Invalid Email)
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John",
    "email": "invalid-email"
  }'
```

### Response (400 Bad Request)
```json
{
  "success": false,
  "message": "Validation failed",
  "error": "Key: 'CreateUserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

---

## 🔄 Common Operations

### 1. Create New User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane","email":"jane@example.com"}'
```

### 2. Get User by ID
```bash
curl http://localhost:8080/api/v1/users/1
```

### 3. Update User Name
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe"}'
```

### 4. Update User Email
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"email":"jane.doe@example.com"}'
```

### 5. Delete User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

### 6. Get Page 2 with 5 Records
```bash
curl "http://localhost:8080/api/v1/users?page=2&page_size=5"
```

---

## 🐛 Troubleshooting

### Database Connection Error
```
Error: Failed to connect to database

Fix:
1. Ensure MySQL is running: systemctl start mysql
2. Check .env credentials match your MySQL setup
3. Verify database exists: CREATE DATABASE incident_report;
```

### Port Already in Use
```
Error: listen tcp :8080: bind: address already in use

Fix:
1. Change SERVER_PORT in .env to a different port
2. Or kill the process: lsof -i :8080 | kill -9 <PID>
```

### Missing Dependencies
```
Error: cannot find package

Fix:
go mod tidy
go mod download
```

### Auto-Migration Not Working
```
Fix:
1. Ensure database exists
2. Check user has proper permissions
3. Review error logs in console
```

---

## 📚 Documentation Index

| Document | Purpose |
|----------|---------|
| README.md | Project overview & setup |
| ARCHITECTURE.md | Design & patterns |
| API_TESTING_GUIDE.md | Testing examples |
| PROJECT_COMPLETION.md | Status & checklist |
| FILE_MANIFEST.md | File listing & statistics |
| QUICK_REFERENCE.md | This file |

---

## 🎯 Important Concepts

### Clean Architecture
```
HTTP Layer (Controllers)
    ↓
Service Layer (Business Logic)
    ↓
Data Layer (GORM & Models)
    ↓
Database (MySQL)
```

### Response Format
```json
{
  "success": boolean,
  "message": "User-friendly message",
  "data": {...},
  "error": "Error details (if error)"
}
```

### HTTP Status Codes
- 200: Success
- 201: Created
- 400: Bad Request/Validation Error
- 404: Not Found
- 500: Server Error

---

## 🔒 Before Production

- [ ] Set `ENVIRONMENT=production` in .env
- [ ] Use strong database passwords
- [ ] Implement JWT authentication
- [ ] Add CORS configuration
- [ ] Enable HTTPS
- [ ] Set up rate limiting
- [ ] Configure logging
- [ ] Monitor error logs
- [ ] Test load handling
- [ ] Backup database strategy

---

## 💡 Development Tips

1. **Hot Reload:** Use `make dev` for auto-restart on file changes
2. **Format Code:** Run `make fmt` before committing
3. **Clean Builds:** Use `make clean && make build`
4. **Check Status:** Monitor logs in terminal during testing
5. **Test Endpoints:** Use cURL or Postman
6. **Database Queries:** Reference `config/database_queries.go`
7. **Error Handling:** Always check response status
8. **Validation:** Refer to `utils/dto.go` for rules

---

## 📊 Quick Stats

- **Go Files:** 8
- **Documentation:** 4 files
- **Lines of Code:** ~1000+
- **API Endpoints:** 6
- **Database Models:** 1 (User)
- **Services:** 1 (UserService)
- **Controllers:** 1 (UserController)
- **Middleware:** 2
- **DTOs:** 5

---

## 🚀 Deploy Checklist

### Before Deployment
- [ ] Test all endpoints locally
- [ ] Verify database migration works
- [ ] Set production environment variables
- [ ] Review error handling
- [ ] Check logging output
- [ ] Test pagination
- [ ] Validate email constraints
- [ ] Test error scenarios

### Deployment Steps
1. Build binary: `make build`
2. Copy to server: `scp incident-report user@host:/app/`
3. Set environment: Update `.env` on server
4. Ensure database exists on server
5. Run: `./incident-report`
6. Monitor logs
7. Test endpoints from client

### Post-Deployment
- [ ] Verify endpoints accessible
- [ ] Check database connection
- [ ] Monitor for errors
- [ ] Test with real data
- [ ] Set up monitoring/alerts
- [ ] Document any issues

---

## 📞 Quick Links

- **Main Documentation:** See README.md
- **Testing Guide:** See API_TESTING_GUIDE.md
- **Architecture Details:** See ARCHITECTURE.md
- **File Listing:** See FILE_MANIFEST.md
- **Status & Checklist:** See PROJECT_COMPLETION.md

---

## ⭐ Key Features

✅ **RESTful API** - Standard HTTP conventions
✅ **CRUD Operations** - Full Create, Read, Update, Delete
✅ **Pagination** - Efficient data retrieval
✅ **Validation** - Input validation at multiple layers
✅ **Error Handling** - Comprehensive error responses
✅ **Clean Code** - Following Go best practices
✅ **Database** - MySQL with GORM ORM
✅ **Environment Config** - .env file support
✅ **Auto-Migration** - Automatic schema creation
✅ **Documentation** - Comprehensive guides

---

## 🎓 Learning Resources

This project demonstrates:
- Go programming fundamentals
- RESTful API design
- Clean Architecture principles
- SOLID principles
- Database design with GORM
- Dependency injection pattern
- Error handling best practices
- API testing strategies

---

**Ready to code? Start with `go run cmd/main.go`! 🚀**

