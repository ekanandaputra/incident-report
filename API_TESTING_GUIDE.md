# API Testing Guide

This document provides comprehensive examples for testing the Incident Report API using various tools.

## Prerequisites

- API running on `http://localhost:8080`
- MySQL database configured and running
- cURL, Postman, or any REST client installed

## 1. Health Check Endpoint

### cURL
```bash
curl http://localhost:8080/api/v1/health
```

### Expected Response
```json
{
  "status": "healthy",
  "message": "Server is running"
}
```

---

## 2. User CRUD Operations

### 2.1 Create User (POST)

#### cURL
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

#### Expected Response (201 Created)
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

#### Validation Errors

**Missing required fields:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe"}'
```

Response (400 Bad Request):
```json
{
  "success": false,
  "message": "Validation failed",
  "error": "Key: 'CreateUserRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"
}
```

**Invalid email format:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "invalid-email"
  }'
```

**Duplicate email:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "john@example.com"
  }'
```

Response (400 Bad Request):
```json
{
  "success": false,
  "message": "Failed to create user",
  "error": "email already exists"
}
```

---

### 2.2 Get All Users (GET with Pagination)

#### Without Pagination Parameters
```bash
curl http://localhost:8080/api/v1/users
```

#### With Pagination
```bash
# Get page 1 with 5 records per page
curl "http://localhost:8080/api/v1/users?page=1&page_size=5"

# Get page 2 with 10 records per page
curl "http://localhost:8080/api/v1/users?page=2&page_size=10"
```

#### Expected Response (200 OK)
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
      },
      {
        "id": 2,
        "name": "Jane Smith",
        "email": "jane@example.com"
      }
    ],
    "page": 1,
    "page_size": 10,
    "total": 2,
    "total_page": 1
  }
}
```

#### Query Parameter Validation

**Invalid page number:**
```bash
curl "http://localhost:8080/api/v1/users?page=0"
```

Response (400 Bad Request):
```json
{
  "success": false,
  "message": "Invalid query parameters",
  "error": "Key: 'PaginationQuery.Page' Error:Field validation for 'Page' failed on the 'min' tag"
}
```

---

### 2.3 Get Specific User (GET with ID)

#### cURL
```bash
curl http://localhost:8080/api/v1/users/1
```

#### Expected Response (200 OK)
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

#### User Not Found (404)
```bash
curl http://localhost:8080/api/v1/users/999
```

Response (404 Not Found):
```json
{
  "success": false,
  "message": "User not found",
  "error": "user not found"
}
```

#### Invalid ID Format
```bash
curl http://localhost:8080/api/v1/users/abc
```

Response (400 Bad Request):
```json
{
  "success": false,
  "message": "Invalid user ID",
  "error": "ID must be a valid number"
}
```

---

### 2.4 Update User (PUT)

#### Update Name
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jonathan Doe"
  }'
```

#### Update Email
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "jonathan@example.com"
  }'
```

#### Update Both
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jonathan Doe",
    "email": "jonathan@example.com"
  }'
```

#### Expected Response (200 OK)
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": 1,
    "name": "Jonathan Doe",
    "email": "jonathan@example.com"
  }
}
```

#### User Not Found
```bash
curl -X PUT http://localhost:8080/api/v1/users/999 \
  -H "Content-Type: application/json" \
  -d '{"name": "New Name"}'
```

Response (400 Bad Request):
```json
{
  "success": false,
  "message": "Failed to update user",
  "error": "user not found"
}
```

---

### 2.5 Delete User (DELETE)

#### cURL
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

#### Expected Response (200 OK)
```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

#### User Not Found
```bash
curl -X DELETE http://localhost:8080/api/v1/users/999
```

Response (400 Bad Request):
```json
{
  "success": false,
  "message": "Failed to delete user",
  "error": "user not found"
}
```

---

## 3. Testing with Postman

### Import Collection

1. Create a new Postman collection named "Incident Report API"
2. Add the following requests:

#### Request 1: Create User
- **Name:** Create User
- **Method:** POST
- **URL:** `{{base_url}}/api/v1/users`
- **Headers:** `Content-Type: application/json`
- **Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

#### Request 2: Get All Users
- **Name:** Get All Users
- **Method:** GET
- **URL:** `{{base_url}}/api/v1/users?page=1&page_size=10`

#### Request 3: Get User by ID
- **Name:** Get User
- **Method:** GET
- **URL:** `{{base_url}}/api/v1/users/1`

#### Request 4: Update User
- **Name:** Update User
- **Method:** PUT
- **URL:** `{{base_url}}/api/v1/users/1`
- **Headers:** `Content-Type: application/json`
- **Body:**
```json
{
  "name": "Jane Doe",
  "email": "jane@example.com"
}
```

#### Request 5: Delete User
- **Name:** Delete User
- **Method:** DELETE
- **URL:** `{{base_url}}/api/v1/users/1`

### Set Environment Variables

Create a Postman environment with:
```
base_url: http://localhost:8080
```

---

## 4. Batch Testing Script

### Bash Script for Complete Testing

```bash
#!/bin/bash

BASE_URL="http://localhost:8080/api/v1"

echo "🧪 Starting Incident Report API Test Suite"
echo "=========================================="

# Test 1: Health Check
echo -e "\n1️⃣ Testing Health Check..."
curl -X GET "$BASE_URL/health" | json_pp

# Test 2: Create User
echo -e "\n2️⃣ Creating User..."
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/users" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com"
  }')
echo "$CREATE_RESPONSE" | json_pp
USER_ID=$(echo "$CREATE_RESPONSE" | grep -o '"id":[0-9]*' | head -1 | grep -o '[0-9]*')

# Test 3: Get All Users
echo -e "\n3️⃣ Getting All Users..."
curl -s -X GET "$BASE_URL/users?page=1&page_size=10" | json_pp

# Test 4: Get Specific User
echo -e "\n4️⃣ Getting User by ID..."
curl -s -X GET "$BASE_URL/users/$USER_ID" | json_pp

# Test 5: Update User
echo -e "\n5️⃣ Updating User..."
curl -s -X PUT "$BASE_URL/users/$USER_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated User",
    "email": "updated@example.com"
  }' | json_pp

# Test 6: Delete User
echo -e "\n6️⃣ Deleting User..."
curl -s -X DELETE "$BASE_URL/users/$USER_ID" | json_pp

echo -e "\n=========================================="
echo "✅ Test Suite Complete!"
```

Save as `test.sh` and run:
```bash
chmod +x test.sh
./test.sh
```

---

## 5. Performance Testing

### Using Apache Bench

```bash
# Create multiple users
ab -n 100 -c 10 -p user.json -T application/json http://localhost:8080/api/v1/users

# Get all users (100 requests, 10 concurrent)
ab -n 100 -c 10 http://localhost:8080/api/v1/users
```

### Using Wrk

```bash
# Test with wrk (needs installation)
wrk -t4 -c100 -d30s http://localhost:8080/api/v1/users
```

---

## 6. Common Error Scenarios

### 1. Database Connection Error
- Ensure MySQL is running
- Verify credentials in `.env`
- Check database exists

### 2. Port Already in Use
- Change `SERVER_PORT` in `.env`
- Or kill existing process

### 3. Missing Dependencies
```bash
go mod tidy
go mod download
```

---

## 7. Response Status Codes Reference

| Status Code | Meaning | Example |
|-------------|---------|---------|
| 200 | OK | Get/Update/Delete success |
| 201 | Created | User created successfully |
| 400 | Bad Request | Validation error, user not found |
| 404 | Not Found | User ID doesn't exist |
| 500 | Internal Server Error | Database error, panic recovery |

---

## Tips for Testing

✅ Always verify `.env` configuration before testing
✅ Use `json_pp` for pretty-printing JSON responses
✅ Test both success and error scenarios
✅ Test pagination with different page sizes
✅ Test validation with invalid data
✅ Check database directly with MySQL client
✅ Monitor server logs for errors

---

## 8. Other API Endpoints

This section lists all remaining API endpoints implemented in the project with short cURL examples and expected response shapes.

- **Base path:** `http://localhost:8080/api/v1`

- **Buildings**
  - Create building (POST):
    ```bash
    curl -X POST http://localhost:8080/api/v1/buildings \
      -H "Content-Type: application/json" \
      -d '{"name":"Building A","address":"123 Main St"}'
    ```
    Expected: 201 Created with created building object

  - Get all buildings (GET):
    ```bash
    curl http://localhost:8080/api/v1/buildings
    ```
    Expected: 200 OK, array of buildings

  - Get building (GET):
    ```bash
    curl http://localhost:8080/api/v1/buildings/1
    ```
    Expected: 200 OK, building object

  - Get floors in building (GET):
    ```bash
    curl http://localhost:8080/api/v1/buildings/1/floors
    ```
    Expected: 200 OK, array of floors for building

  - Update building (PUT):
    ```bash
    curl -X PUT http://localhost:8080/api/v1/buildings/1 \
      -H "Content-Type: application/json" \
      -d '{"name":"New Name"}'
    ```

  - Delete building (DELETE):
    ```bash
    curl -X DELETE http://localhost:8080/api/v1/buildings/1
    ```

- **Floors**
  - Create floor (POST):
    ```bash
    curl -X POST http://localhost:8080/api/v1/floors \
      -H "Content-Type: application/json" \
      -d '{"building_id":1,"name":"1st Floor"}'
    ```

  - Get floor (GET):
    ```bash
    curl http://localhost:8080/api/v1/floors/1
    ```

  - Get rooms on floor (GET):
    ```bash
    curl http://localhost:8080/api/v1/floors/1/rooms
    ```

  - Update floor (PUT):
    ```bash
    curl -X PUT http://localhost:8080/api/v1/floors/1 \
      -H "Content-Type: application/json" \
      -d '{"name":"Ground Floor"}'
    ```

  - Delete floor (DELETE):
    ```bash
    curl -X DELETE http://localhost:8080/api/v1/floors/1
    ```

- **Rooms**
  - Create room (POST):
    ```bash
    curl -X POST http://localhost:8080/api/v1/rooms \
      -H "Content-Type: application/json" \
      -d '{"floor_id":1,"name":"Room 101"}'
    ```

  - Get room (GET):
    ```bash
    curl http://localhost:8080/api/v1/rooms/1
    ```

  - Get components in room (GET):
    ```bash
    curl http://localhost:8080/api/v1/rooms/1/components
    ```

  - Update room (PUT):
    ```bash
    curl -X PUT http://localhost:8080/api/v1/rooms/1 \
      -H "Content-Type: application/json" \
      -d '{"name":"Updated Room"}'
    ```

  - Delete room (DELETE):
    ```bash
    curl -X DELETE http://localhost:8080/api/v1/rooms/1
    ```

- **Component Categories**
  - Create category (POST):
    ```bash
    curl -X POST http://localhost:8080/api/v1/component-categories \
      -H "Content-Type: application/json" \
      -d '{"name":"Electrical"}'
    ```

  - Get all categories (GET):
    ```bash
    curl http://localhost:8080/api/v1/component-categories
    ```

  - Get category (GET):
    ```bash
    curl http://localhost:8080/api/v1/component-categories/1
    ```

  - Get components in category (GET):
    ```bash
    curl http://localhost:8080/api/v1/component-categories/1/components
    ```

  - Update category (PUT):
    ```bash
    curl -X PUT http://localhost:8080/api/v1/component-categories/1 \
      -H "Content-Type: application/json" \
      -d '{"name":"New Category"}'
    ```

  - Delete category (DELETE):
    ```bash
    curl -X DELETE http://localhost:8080/api/v1/component-categories/1
    ```

- **Components**
  - Create component (POST):
    ```bash
    curl -X POST http://localhost:8080/api/v1/components \
      -H "Content-Type: application/json" \
      -d '{"room_id":1,"category_id":1,"name":"Smoke Detector","serial_number":"SN123"}'
    ```

  - Get component (GET):
    ```bash
    curl http://localhost:8080/api/v1/components/1
    ```

  - Update component (PUT):
    ```bash
    curl -X PUT http://localhost:8080/api/v1/components/1 \
      -H "Content-Type: application/json" \
      -d '{"name":"New Component Name"}'
    ```

  - Delete component (DELETE):
    ```bash
    curl -X DELETE http://localhost:8080/api/v1/components/1
    ```

---

**Happy Testing! 🚀**
**Happy Testing! 🚀**
