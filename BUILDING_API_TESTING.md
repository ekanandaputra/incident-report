# Building Management API Testing Guide

This guide provides comprehensive instructions for testing all building management API endpoints. The API follows RESTful conventions with proper HTTP status codes and error handling.

## Base URL

```
http://localhost:8080/api/v1
```

## API Endpoints Overview

### Building Endpoints
- `POST /buildings` - Create a new building
- `GET /buildings` - Get all buildings (paginated)
- `GET /buildings/:id` - Get a specific building
- `PUT /buildings/:id` - Update a building
- `DELETE /buildings/:id` - Delete a building

### Floor Endpoints
- `POST /floors` - Create a new floor
- `GET /floors/:id` - Get a specific floor
- `GET /buildings/:buildingId/floors` - Get all floors in a building (paginated)
- `PUT /floors/:id` - Update a floor
- `DELETE /floors/:id` - Delete a floor

### Room Endpoints
- `POST /rooms` - Create a new room
- `GET /rooms/:id` - Get a specific room
- `GET /floors/:floorId/rooms` - Get all rooms on a floor (paginated)
- `PUT /rooms/:id` - Update a room
- `DELETE /rooms/:id` - Delete a room

### Component Category Endpoints
- `POST /component-categories` - Create a new component category
- `GET /component-categories` - Get all component categories (paginated)
- `GET /component-categories/:id` - Get a specific component category
- `PUT /component-categories/:id` - Update a component category
- `DELETE /component-categories/:id` - Delete a component category

### Component Endpoints
- `POST /components` - Create a new component
- `GET /components/:id` - Get a specific component
- `GET /rooms/:roomId/components` - Get all components in a room (paginated)
- `GET /component-categories/:categoryId/components` - Get all components in a category (paginated)
- `PUT /components/:id` - Update a component
- `DELETE /components/:id` - Delete a component

## Testing with cURL

### 1. Create a Building

```bash
curl -X POST http://localhost:8080/api/v1/buildings \
  -H "Content-Type: application/json" \
  -d '{
    "code": "BLD001",
    "name": "Main Building",
    "location": "Downtown"
  }'
```

**Expected Response (201 Created):**
```json
{
  "status": "success",
  "message": "Building created successfully",
  "data": {
    "id": 1,
    "code": "BLD001",
    "name": "Main Building",
    "location": "Downtown",
    "created_at": "2024-01-10T10:30:00Z",
    "updated_at": "2024-01-10T10:30:00Z"
  }
}
```

### 2. Get All Buildings (with Pagination)

```bash
curl http://localhost:8080/api/v1/buildings?page=1&pageSize=10
```

**Query Parameters:**
- `page` - Page number (default: 1)
- `pageSize` - Number of records per page (default: 10, max: 100)

**Expected Response (200 OK):**
```json
{
  "status": "success",
  "message": "Buildings retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "code": "BLD001",
        "name": "Main Building",
        "location": "Downtown",
        "created_at": "2024-01-10T10:30:00Z",
        "updated_at": "2024-01-10T10:30:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "pageSize": 10,
    "totalPages": 1
  }
}
```

### 3. Get a Specific Building

```bash
curl http://localhost:8080/api/v1/buildings/1
```

**Expected Response (200 OK):**
```json
{
  "status": "success",
  "message": "Building retrieved successfully",
  "data": {
    "id": 1,
    "code": "BLD001",
    "name": "Main Building",
    "location": "Downtown",
    "created_at": "2024-01-10T10:30:00Z",
    "updated_at": "2024-01-10T10:30:00Z"
  }
}
```

### 4. Update a Building

```bash
curl -X PUT http://localhost:8080/api/v1/buildings/1 \
  -H "Content-Type: application/json" \
  -d '{
    "code": "BLD001-UPD",
    "name": "Main Building Updated",
    "location": "Downtown - Updated"
  }'
```

**Expected Response (200 OK):**
```json
{
  "status": "success",
  "message": "Building updated successfully",
  "data": {
    "id": 1,
    "code": "BLD001-UPD",
    "name": "Main Building Updated",
    "location": "Downtown - Updated",
    "created_at": "2024-01-10T10:30:00Z",
    "updated_at": "2024-01-10T10:35:00Z"
  }
}
```

### 5. Delete a Building

```bash
curl -X DELETE http://localhost:8080/api/v1/buildings/1
```

**Expected Response (204 No Content):**
```
[Empty response body]
```

### 6. Create a Floor (requires existing building)

```bash
curl -X POST http://localhost:8080/api/v1/floors \
  -H "Content-Type: application/json" \
  -d '{
    "building_id": 1,
    "number": 1,
    "name": "Ground Floor"
  }'
```

**Expected Response (201 Created):**
```json
{
  "status": "success",
  "message": "Floor created successfully",
  "data": {
    "id": 1,
    "building_id": 1,
    "number": 1,
    "name": "Ground Floor",
    "created_at": "2024-01-10T10:30:00Z",
    "updated_at": "2024-01-10T10:30:00Z"
  }
}
```

### 7. Get All Floors in a Building

```bash
curl http://localhost:8080/api/v1/buildings/1/floors?page=1&pageSize=10
```

**Expected Response (200 OK):**
```json
{
  "status": "success",
  "message": "Floors retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "building_id": 1,
        "number": 1,
        "name": "Ground Floor",
        "created_at": "2024-01-10T10:30:00Z",
        "updated_at": "2024-01-10T10:30:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "pageSize": 10,
    "totalPages": 1
  }
}
```

### 8. Create a Room (requires existing floor)

```bash
curl -X POST http://localhost:8080/api/v1/rooms \
  -H "Content-Type: application/json" \
  -d '{
    "floor_id": 1,
    "code": "ROOM001",
    "name": "Conference Room A"
  }'
```

**Expected Response (201 Created):**
```json
{
  "status": "success",
  "message": "Room created successfully",
  "data": {
    "id": 1,
    "floor_id": 1,
    "code": "ROOM001",
    "name": "Conference Room A",
    "created_at": "2024-01-10T10:30:00Z",
    "updated_at": "2024-01-10T10:30:00Z"
  }
}
```

### 9. Get All Rooms on a Floor

```bash
curl http://localhost:8080/api/v1/floors/1/rooms?page=1&pageSize=10
```

**Expected Response (200 OK):**
```json
{
  "status": "success",
  "message": "Rooms retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "floor_id": 1,
        "code": "ROOM001",
        "name": "Conference Room A",
        "created_at": "2024-01-10T10:30:00Z",
        "updated_at": "2024-01-10T10:30:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "pageSize": 10,
    "totalPages": 1
  }
}
```

### 10. Create a Component Category

```bash
curl -X POST http://localhost:8080/api/v1/component-categories \
  -H "Content-Type: application/json" \
  -d '{
    "code": "CAT001",
    "name": "HVAC Systems",
    "description": "Heating, Ventilation, and Air Conditioning equipment"
  }'
```

**Expected Response (201 Created):**
```json
{
  "status": "success",
  "message": "Component category created successfully",
  "data": {
    "id": 1,
    "code": "CAT001",
    "name": "HVAC Systems",
    "description": "Heating, Ventilation, and Air Conditioning equipment",
    "created_at": "2024-01-10T10:30:00Z",
    "updated_at": "2024-01-10T10:30:00Z"
  }
}
```

### 11. Get All Component Categories

```bash
curl http://localhost:8080/api/v1/component-categories?page=1&pageSize=10
```

**Expected Response (200 OK):**
```json
{
  "status": "success",
  "message": "Component categories retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "code": "CAT001",
        "name": "HVAC Systems",
        "description": "Heating, Ventilation, and Air Conditioning equipment",
        "created_at": "2024-01-10T10:30:00Z",
        "updated_at": "2024-01-10T10:30:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "pageSize": 10,
    "totalPages": 1
  }
}
```

### 12. Create a Component (requires existing room and category)

```bash
curl -X POST http://localhost:8080/api/v1/components \
  -H "Content-Type: application/json" \
  -d '{
    "room_id": 1,
    "category_id": 1,
    "code": "COMP001",
    "name": "AC Unit",
    "brand": "Daikin",
    "specification": "5000 BTU",
    "procurement_year": 2022
  }'
```

**Expected Response (201 Created):**
```json
{
  "status": "success",
  "message": "Component created successfully",
  "data": {
    "id": 1,
    "room_id": 1,
    "category_id": 1,
    "code": "COMP001",
    "name": "AC Unit",
    "brand": "Daikin",
    "specification": "5000 BTU",
    "procurement_year": 2022,
    "created_at": "2024-01-10T10:30:00Z",
    "updated_at": "2024-01-10T10:30:00Z"
  }
}
```

### 13. Get All Components in a Room

```bash
curl http://localhost:8080/api/v1/rooms/1/components?page=1&pageSize=10
```

**Expected Response (200 OK):**
```json
{
  "status": "success",
  "message": "Components retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "room_id": 1,
        "category_id": 1,
        "code": "COMP001",
        "name": "AC Unit",
        "brand": "Daikin",
        "specification": "5000 BTU",
        "procurement_year": 2022,
        "created_at": "2024-01-10T10:30:00Z",
        "updated_at": "2024-01-10T10:30:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "pageSize": 10,
    "totalPages": 1
  }
}
```

### 14. Get All Components in a Category

```bash
curl http://localhost:8080/api/v1/component-categories/1/components?page=1&pageSize=10
```

**Expected Response (200 OK):**
```json
{
  "status": "success",
  "message": "Components retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "room_id": 1,
        "category_id": 1,
        "code": "COMP001",
        "name": "AC Unit",
        "brand": "Daikin",
        "specification": "5000 BTU",
        "procurement_year": 2022,
        "created_at": "2024-01-10T10:30:00Z",
        "updated_at": "2024-01-10T10:30:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "pageSize": 10,
    "totalPages": 1
  }
}
```

## Error Responses

### 400 Bad Request
```json
{
  "status": "error",
  "message": "Invalid request format",
  "error": "validation error details"
}
```

### 404 Not Found
```json
{
  "status": "error",
  "message": "Building not found",
  "error": "building not found"
}
```

### 500 Internal Server Error
```json
{
  "status": "error",
  "message": "Failed to create building",
  "error": "database error details"
}
```

## Testing with Postman

1. Create a new Postman Collection
2. Import the endpoints from the URL structure above
3. Set the base URL: `{{base_url}}/api/v1` where `base_url` is `http://localhost:8080`
4. Create environment variables for IDs used in subsequent requests
5. Use the POST responses to extract IDs for dependent operations

## Testing Order Recommendation

For complete end-to-end testing, follow this order:

1. Create a Building
2. Create a Floor (using building_id from step 1)
3. Create a Component Category
4. Create a Room (using floor_id from step 2)
5. Create a Component (using room_id from step 4 and category_id from step 3)
6. Retrieve all buildings with pagination
7. Retrieve all floors in the building
8. Retrieve all rooms in the floor
9. Retrieve all components in the room
10. Retrieve all components in the category
11. Update building details
12. Update floor details
13. Update room details
14. Update component details
15. Delete operations (reverse order: component, room, floor, building)

## Data Validation Rules

### Building
- `code` - Required, unique string
- `name` - Required, non-empty string
- `location` - Required, non-empty string

### Floor
- `building_id` - Required, must exist in buildings table
- `number` - Required, positive integer
- `name` - Required, non-empty string

### Room
- `floor_id` - Required, must exist in floors table
- `code` - Required, unique string
- `name` - Required, non-empty string

### Component Category
- `code` - Required, unique string
- `name` - Required, non-empty string
- `description` - Optional, string

### Component
- `room_id` - Required, must exist in rooms table
- `category_id` - Required, must exist in component_categories table
- `code` - Required, unique string
- `name` - Required, non-empty string
- `brand` - Optional, string
- `specification` - Optional, string
- `procurement_year` - Optional, positive integer

## Performance Notes

- All list endpoints support pagination with default page size of 10 and maximum 100
- Foreign key constraints ensure referential integrity
- Cascading deletes remove related records automatically
- Soft deletes preserve historical data

## Next Steps

After testing all endpoints:
1. Verify database tables were created correctly
2. Test foreign key constraints
3. Test cascading delete behavior
4. Load test with bulk data
5. Implement frontend integration
