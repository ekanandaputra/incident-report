# Building Management API - Quick Reference

## Project Structure Overview

```
Models (data structures)
    ↓
Services (business logic)
    ↓
Controllers (HTTP handlers)
    ↓
Routes (endpoint registration)
```

## Entity Relationships

```
Building (1) ──→ (*) Floor
                     ↓
                  (1) ──→ (*) Room
                              ↓
                           (1) ──→ (*) Component
                                       ↑
                                       │
                       ComponentCategory (1) ──→ (*)
```

## Running the Application

### Prerequisites
- Go 1.21+
- MySQL 5.7+
- Environment variables configured in `.env`

### Setup .env File
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=incident_report
```

### Start Server
```bash
# Development mode
go run cmd/main.go

# Build executable
go build -o app cmd/main.go
./app

# Using Makefile (if available)
make run
```

Server runs on: `http://localhost:8080`

## API Response Format

### Success Response
```json
{
  "status": "success",
  "message": "Operation completed successfully",
  "data": {
    // Response payload
  }
}
```

### Error Response
```json
{
  "status": "error",
  "message": "User-friendly error message",
  "error": "Detailed error information"
}
```

## Common HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | OK - Request succeeded |
| 201 | Created - Resource created successfully |
| 204 | No Content - Successful deletion |
| 400 | Bad Request - Invalid input |
| 404 | Not Found - Resource doesn't exist |
| 500 | Internal Error - Server error |

## Pagination

All list endpoints support pagination:

```bash
curl "http://localhost:8080/api/v1/buildings?page=2&pageSize=20"
```

Parameters:
- `page` - Page number (default: 1)
- `pageSize` - Items per page (default: 10, max: 100)

Response includes:
```json
{
  "data": [...],
  "total": 100,
  "page": 2,
  "page_size": 20,
  "total_page": 5
}
```

## CRUD Operations Pattern

### Create (POST)
```bash
curl -X POST http://localhost:8080/api/v1/buildings \
  -H "Content-Type: application/json" \
  -d '{
    "code": "BLD001",
    "name": "Building Name",
    "location": "Location"
  }'
```

### Read (GET)
```bash
# Single record
curl http://localhost:8080/api/v1/buildings/1

# All records (paginated)
curl http://localhost:8080/api/v1/buildings?page=1&pageSize=10
```

### Update (PUT)
```bash
curl -X PUT http://localhost:8080/api/v1/buildings/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Name"
  }'
```

### Delete (DELETE)
```bash
curl -X DELETE http://localhost:8080/api/v1/buildings/1
```

## Building Management Endpoints

### Buildings
```
POST   /api/v1/buildings
GET    /api/v1/buildings
GET    /api/v1/buildings/:id
PUT    /api/v1/buildings/:id
DELETE /api/v1/buildings/:id
GET    /api/v1/buildings/:id/floors          # Get floors in building
```

### Floors
```
POST   /api/v1/floors
GET    /api/v1/floors/:id
PUT    /api/v1/floors/:id
DELETE /api/v1/floors/:id
GET    /api/v1/buildings/:buildingId/floors  # Get floors by building
GET    /api/v1/floors/:id/rooms              # Get rooms on floor
```

### Rooms
```
POST   /api/v1/rooms
GET    /api/v1/rooms/:id
PUT    /api/v1/rooms/:id
DELETE /api/v1/rooms/:id
GET    /api/v1/floors/:floorId/rooms         # Get rooms by floor
GET    /api/v1/rooms/:id/components          # Get components in room
```

### Component Categories
```
POST   /api/v1/component-categories
GET    /api/v1/component-categories
GET    /api/v1/component-categories/:id
PUT    /api/v1/component-categories/:id
DELETE /api/v1/component-categories/:id
GET    /api/v1/component-categories/:id/components  # Get components by category
```

### Components
```
POST   /api/v1/components
GET    /api/v1/components/:id
PUT    /api/v1/components/:id
DELETE /api/v1/components/:id
GET    /api/v1/rooms/:roomId/components             # Get components by room
GET    /api/v1/component-categories/:categoryId/components  # Get components by category
```

## Required Fields by Model

### Building
- `code` (string, unique) - Building identifier
- `name` (string) - Building name
- `location` (string) - Building location

### Floor
- `building_id` (integer) - Parent building ID
- `number` (integer) - Floor number
- `name` (string) - Floor name

### Room
- `floor_id` (integer) - Parent floor ID
- `code` (string, unique) - Room identifier
- `name` (string) - Room name

### Component Category
- `code` (string, unique) - Category identifier
- `name` (string) - Category name
- `description` (string, optional) - Category description

### Component
- `room_id` (integer) - Parent room ID
- `category_id` (integer) - Parent category ID
- `code` (string, unique) - Component identifier
- `name` (string) - Component name
- `brand` (string, optional) - Equipment brand
- `specification` (string, optional) - Technical specs
- `procurement_year` (integer, optional) - Year purchased

## Development Workflow

### Adding a New Endpoint

1. **Add to Model** (if needed) - Define struct with GORM tags
2. **Update Service** - Add business logic method
3. **Update Controller** - Add HTTP handler method
4. **Register Route** - Add endpoint to `routes.go`
5. **Test** - Test with cURL or Postman
6. **Document** - Update API documentation

### Testing

```bash
# Test health check
curl http://localhost:8080/api/v1/health

# Test building creation
curl -X POST http://localhost:8080/api/v1/buildings \
  -H "Content-Type: application/json" \
  -d '{"code":"TEST","name":"Test","location":"Test"}'

# View all buildings
curl http://localhost:8080/api/v1/buildings
```

## Database Operations

### Auto-Migration
Runs automatically on startup from `config/database.go`:
- Creates tables if they don't exist
- Updates schema for model changes
- No manual migration needed

### Soft Deletes
- All models support soft delete (DeletedAt timestamp)
- Deleted records are not returned in queries
- Can be permanently deleted if needed

### Relationships
- Cascading deletes via GORM tags
- Automatic foreign key creation
- Lazy loading with `Preload()` method

## Common Tasks

### Create a Building with Floors
```bash
# 1. Create building
BUILDING_ID=$(curl -s -X POST http://localhost:8080/api/v1/buildings \
  -H "Content-Type: application/json" \
  -d '{"code":"BLD1","name":"Building 1","location":"NYC"}' | jq '.data.id')

# 2. Create floor
curl -X POST http://localhost:8080/api/v1/floors \
  -H "Content-Type: application/json" \
  -d "{\"building_id\":$BUILDING_ID,\"number\":1,\"name\":\"Ground Floor\"}"
```

### Find All Components in a Room
```bash
curl http://localhost:8080/api/v1/rooms/1/components
```

### Update Component Details
```bash
curl -X PUT http://localhost:8080/api/v1/components/1 \
  -H "Content-Type: application/json" \
  -d '{"brand":"NewBrand","specification":"Updated specs"}'
```

### Delete Building (cascades to floors, rooms, components)
```bash
curl -X DELETE http://localhost:8080/api/v1/buildings/1
```

## Troubleshooting

### Building fails to start
- Check MySQL is running
- Verify `.env` file has correct credentials
- Check port 8080 is available

### Foreign key errors
- Ensure parent record exists before creating child
- Building ID must exist before creating floor
- Floor ID must exist before creating room

### Validation errors
- Check required fields are provided
- Verify field types (numbers vs strings)
- Check string length constraints

### Database errors
- Review error message in console
- Check database permissions
- Ensure tables exist (auto-migration should create them)

## Performance Tips

1. Use pagination for large datasets
2. Limit page size to 50 or less for better performance
3. Index frequently searched fields (code fields are indexed)
4. Cache frequently accessed data if needed
5. Monitor slow queries in production

## References

- Full Testing Guide: [BUILDING_API_TESTING.md](BUILDING_API_TESTING.md)
- Implementation Details: [BUILDING_SYSTEM_IMPLEMENTATION.md](BUILDING_SYSTEM_IMPLEMENTATION.md)
- Architecture: [ARCHITECTURE.md](ARCHITECTURE.md)
- General README: [README.md](README.md)

## Key Files

| File | Purpose |
|------|---------|
| `models/building.go` | Data structure definitions |
| `services/building_service.go` | Business logic implementation |
| `controllers/building_controller.go` | HTTP request handlers |
| `routes/routes.go` | API endpoint registration |
| `config/database.go` | Database connection and migration |
| `utils/building_dtos.go` | Request/response objects |

## Code Example: Complete Flow

```bash
#!/bin/bash
# Complete CRUD flow example

API="http://localhost:8080/api/v1"

# 1. CREATE
echo "Creating building..."
BUILDING=$(curl -s -X POST $API/buildings \
  -H "Content-Type: application/json" \
  -d '{"code":"EX001","name":"Example","location":"Test"}')
BID=$(echo $BUILDING | grep -o '"id":[0-9]*' | head -1 | grep -o '[0-9]*')

# 2. READ
echo "Reading building $BID..."
curl -s $API/buildings/$BID | jq .

# 3. UPDATE
echo "Updating building..."
curl -s -X PUT $API/buildings/$BID \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Example"}' | jq .

# 4. DELETE
echo "Deleting building..."
curl -s -X DELETE $API/buildings/$BID
```

## Support

For issues or questions:
1. Check this quick reference
2. Review [BUILDING_API_TESTING.md](BUILDING_API_TESTING.md) for detailed examples
3. Check [BUILDING_SYSTEM_IMPLEMENTATION.md](BUILDING_SYSTEM_IMPLEMENTATION.md) for architecture details
