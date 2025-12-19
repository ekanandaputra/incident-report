package config

// This file contains GORM query examples for reference
// Uncomment and modify as needed for your use cases

/*
BASIC QUERY EXAMPLES
====================

// Find by primary key
var user models.User
DB.First(&user, 1) // Find record with ID = 1

// Find with conditions
DB.Where("email = ?", "user@example.com").First(&user)

// Find all
var users []models.User
DB.Find(&users)

// Find with pagination
DB.Offset(0).Limit(10).Find(&users)

// Count records
var count int64
DB.Model(&models.User{}).Count(&count)

// Select specific columns
DB.Select("id", "name").Find(&users)

// Order results
DB.Order("created_at desc").Find(&users)

// Group and aggregate
DB.Model(&models.User{}).Select("COUNT(*) as count").Group("name")

WRITE OPERATIONS
================

// Create single record
user := models.User{Name: "John", Email: "john@example.com"}
DB.Create(&user)

// Create multiple records
users := []models.User{
    {Name: "User1", Email: "user1@example.com"},
    {Name: "User2", Email: "user2@example.com"},
}
DB.CreateInBatches(users, 100)

// Update single field
DB.Model(&models.User{}).Where("id = ?", 1).Update("name", "New Name")

// Update multiple fields
DB.Model(&models.User{}).Where("id = ?", 1).Updates(models.User{Name: "New Name", Email: "new@example.com"})

// Update with map
DB.Model(&models.User{}).Where("id = ?", 1).Updates(map[string]interface{}{"name": "New Name"})

// Delete record (hard delete - permanent)
DB.Delete(&user, 1)

// Delete with conditions (hard delete)
DB.Where("email = ?", "old@example.com").Delete(&models.User{})

// Soft delete (marks deleted_at timestamp)
DB.Delete(&user)

ADVANCED QUERIES
================

// Join tables (when you add relationships)
DB.Joins("JOIN other_table ON ...").Find(&users)

// Subquery
subquery := DB.Select("name").Where("age > ?", 18)
DB.Where("name IN ?", subquery).Find(&users)

// Raw SQL
DB.Raw("SELECT * FROM users WHERE id = ?", 1).Scan(&user)

// With transaction
tx := DB.BeginTx(ctx, nil)
tx.Create(&user)
tx.Commit()

HOOKS (Lifecycle callbacks)
==========================

// Model can have hooks like:
func (u *User) BeforeCreate(tx *gorm.DB) error {
    // Validation before creating
    return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
    // Action after creating
    return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
    // Validation before updating
    return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
    // Action before deleting
    return nil
}

ERROR HANDLING
==============

// Check if record not found
if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    // Handle not found
}

// Check other errors
if result.Error != nil {
    // Handle error
    log.Printf("Error: %v", result.Error)
}

// Get row affected count
affected := result.RowsAffected

PERFORMANCE TIPS
================

1. Use indexes on frequently queried columns
   - Email is indexed in User model
   - Add more indexes as needed

2. Use pagination to limit result sets
   DB.Offset(offset).Limit(pageSize).Find(&users)

3. Select only needed columns
   DB.Select("id", "name").Find(&users)

4. Use batching for bulk operations
   DB.CreateInBatches(users, 100)

5. Use prepared statements (automatic with GORM)

6. Monitor slow queries with logging

7. Use database indexes for filtering conditions

RELATIONSHIPS (for future models)
=================================

// One-to-Many relationship example:
type Company struct {
    ID    uint
    Name  string
    Users []User // Has many
}

type User struct {
    ID        uint
    Name      string
    CompanyID uint // Foreign key
    Company   Company
}

// Queries:
var company Company
DB.Preload("Users").First(&company, 1)

// Many-to-Many:
type User struct {
    ID    uint
    Name  string
    Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
    ID   uint
    Name string
}

// Query:
DB.Preload("Roles").First(&user, 1)

SOFT DELETE QUERIES
===================

// Query all (excludes soft deleted)
DB.Find(&users)

// Include soft deleted records
DB.Unscoped().Find(&users)

// Find only soft deleted
DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)

// Restore soft deleted record
DB.Model(&user).Update("deleted_at", nil)

// Permanently delete soft deleted record
DB.Unscoped().Delete(&user)
*/
