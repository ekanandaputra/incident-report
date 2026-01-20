package routes

import (
	"incident-report/controllers"
	"incident-report/middleware"
	"incident-report/services"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all API routes for the application
// It organizes routes using versioning (/api/v1) for better API management
func RegisterRoutes(router *gin.Engine) {
	// Apply global middleware
	router.Use(middleware.ErrorHandlerMiddleware())

	// Serve Swagger UI - main endpoint
	router.GET("/swagger", func(c *gin.Context) {
		c.File("templates/swagger.html")
	})

	// Serve Swagger UI HTML page with specific path
	router.GET("/swagger/index.html", func(c *gin.Context) {
		c.File("templates/swagger.html")
	})

	// Serve swagger.yaml file for Swagger UI to load
	router.StaticFile("/swagger.yaml", "swagger.yaml")

	// Create a service layer instance (dependency injection)
	userService := services.NewUserService()

	// Create a controller instance with the service
	userController := controllers.NewUserController(userService)

	// Create building management controllers
	buildingController := controllers.NewBuildingController()
	floorController := controllers.NewFloorController()
	roomController := controllers.NewRoomController()
	componentCategoryController := controllers.NewComponentCategoryController()
	componentController := controllers.NewComponentController()

	// Create report management controller
	reportService := services.NewReportService()
	reportController := controllers.NewReportController(reportService)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Health check endpoint
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "healthy",
				"message": "Server is running",
			})
		})

		// User routes with RESTful conventions
		// POST   /api/v1/users           - Create a new user
		// GET    /api/v1/users           - Get all users (with pagination)
		// GET    /api/v1/users/:id       - Get a specific user
		// PUT    /api/v1/users/:id       - Update a specific user
		// DELETE /api/v1/users/:id       - Delete a specific user
		users := v1.Group("/users")
		{
			// Create user - POST request
			users.POST("", userController.CreateUser)

			// Get all users - GET request with pagination support
			users.GET("", userController.GetAllUsers)

			// Get specific user - GET request with ID parameter
			users.GET("/:id", userController.GetUser)

			// Update user - PUT request with ID parameter
			users.PUT("/:id", userController.UpdateUser)

			// Delete user - DELETE request with ID parameter
			users.DELETE("/:id", userController.DeleteUser)
		}

		// Building routes
		// POST   /api/v1/buildings           - Create a new building
		// GET    /api/v1/buildings           - Get all buildings
		// GET    /api/v1/buildings/:id       - Get a specific building
		// PUT    /api/v1/buildings/:id       - Update a specific building
		// DELETE /api/v1/buildings/:id       - Delete a specific building
		buildings := v1.Group("/buildings")
		{
			buildings.POST("", buildingController.CreateBuilding)
			buildings.GET("", buildingController.GetAllBuildings)

			// Floors within a building - Register nested routes BEFORE wildcard routes
			// GET    /api/v1/buildings/:id/floors           - Get all floors in a building
			buildings.GET("/:id/floors", floorController.GetFloorsByBuilding)

			buildings.GET("/:id", buildingController.GetBuilding)
			buildings.PUT("/:id", buildingController.UpdateBuilding)
			buildings.DELETE("/:id", buildingController.DeleteBuilding)
		}

		// Floor routes
		// POST   /api/v1/floors           - Create a new floor
		// GET    /api/v1/floors           - Get all floors (with pagination)
		// GET    /api/v1/floors/:id       - Get a specific floor
		// PUT    /api/v1/floors/:id       - Update a specific floor
		// DELETE /api/v1/floors/:id       - Delete a specific floor
		floors := v1.Group("/floors")
		{
			floors.POST("", floorController.CreateFloor)
			floors.GET("", floorController.GetAllFloors)

			// Rooms within a floor - Register nested routes BEFORE wildcard routes
			// GET    /api/v1/floors/:id/rooms           - Get all rooms on a floor
			floors.GET("/:id/rooms", roomController.GetRoomsByFloor)

			floors.GET("/:id", floorController.GetFloor)
			floors.PUT("/:id", floorController.UpdateFloor)
			floors.DELETE("/:id", floorController.DeleteFloor)
		}

		// Room routes
		// POST   /api/v1/rooms           - Create a new room
		// GET    /api/v1/rooms           - Get all rooms (with pagination)
		// GET    /api/v1/rooms/:id       - Get a specific room
		// PUT    /api/v1/rooms/:id       - Update a specific room
		// DELETE /api/v1/rooms/:id       - Delete a specific room
		rooms := v1.Group("/rooms")
		{
			rooms.POST("", roomController.CreateRoom)
			rooms.GET("", roomController.GetAllRooms)

			// Components within a room - Register nested routes BEFORE wildcard routes
			// GET    /api/v1/rooms/:id/components           - Get all components in a room
			rooms.GET("/:id/components", componentController.GetComponentsByRoom)

			rooms.GET("/:id", roomController.GetRoom)
			rooms.PUT("/:id", roomController.UpdateRoom)
			rooms.DELETE("/:id", roomController.DeleteRoom)
		}

		// Component Category routes
		// POST   /api/v1/component-categories           - Create a new component category
		// GET    /api/v1/component-categories           - Get all component categories
		// GET    /api/v1/component-categories/:id       - Get a specific component category
		// PUT    /api/v1/component-categories/:id       - Update a specific component category
		// DELETE /api/v1/component-categories/:id       - Delete a specific component category
		categories := v1.Group("/component-categories")
		{
			categories.POST("", componentCategoryController.CreateComponentCategory)
			categories.GET("", componentCategoryController.GetAllComponentCategories)

			// Components within a category - Register nested routes BEFORE wildcard routes
			// GET    /api/v1/component-categories/:id/components           - Get all components in a category
			categories.GET("/:id/components", componentController.GetComponentsByCategory)

			categories.GET("/:id", componentCategoryController.GetComponentCategory)
			categories.PUT("/:id", componentCategoryController.UpdateComponentCategory)
			categories.DELETE("/:id", componentCategoryController.DeleteComponentCategory)
		}

		// Component routes
		// POST   /api/v1/components           - Create a new component
		// GET    /api/v1/components/:id       - Get a specific component
		// PUT    /api/v1/components/:id       - Update a specific component
		// DELETE /api/v1/components/:id       - Delete a specific component
		components := v1.Group("/components")
		{
			components.POST("", componentController.CreateComponent)
			components.GET("/:id", componentController.GetComponent)
			components.PUT("/:id", componentController.UpdateComponent)
			components.DELETE("/:id", componentController.DeleteComponent)
		}

		// Report routes
		// POST   /api/v1/reports           - Create a new report
		// GET    /api/v1/reports           - Get all reports (with pagination)
		// GET    /api/v1/reports/:id       - Get a specific report
		// PUT    /api/v1/reports/:id       - Update a specific report
		// DELETE /api/v1/reports/:id       - Delete a specific report
		reports := v1.Group("/reports")
		{
			reports.POST("", reportController.CreateReport)
			reports.GET("", reportController.GetAllReports)
			reports.GET("/:id", reportController.GetReport)
			reports.PUT("/:id", reportController.UpdateReport)
			reports.DELETE("/:id", reportController.DeleteReport)
			reports.PUT("/:id/assign-user", reportController.AssignUserToReport)
		}
	}
}
