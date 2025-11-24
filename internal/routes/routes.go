package routes

import (
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/controllers"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/services"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(repo repository.Repository) *gin.Engine {
	r := gin.Default()

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize services
	workspaceService := services.NewWorkspaceService(repo)
	projectService := services.NewProjectService(repo)
	taskService := services.NewTaskService(repo)

	// Initialize controllers
	authController := controllers.NewAuthController(repo)
	managerController := controllers.NewManagerController(workspaceService, projectService, taskService)
	devController := controllers.NewDevController(taskService, projectService)

	// Public routes
	public := r.Group("/api")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		public.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}

	// Protected routes (require authentication)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", authController.GetProfile)
	}

	// Manager and Admin routes
	manager := r.Group("/api/manager")
	manager.Use(middleware.AuthMiddleware())
	manager.Use(middleware.RoleMiddleware(models.RoleAdmin, models.RoleManager))
	{
		// Workspace management
		manager.POST("/workspaces", managerController.CreateWorkspace)
		manager.GET("/workspaces/:workspace_id/projects", managerController.ListWorkspaceProjects)

		// Project management
		manager.POST("/projects", managerController.CreateProject)

		// Task assignment
		manager.PUT("/tasks/:id/assign", managerController.AssignTask)
	}

	// Developer routes (all authenticated users can access)
	dev := r.Group("/api/dev")
	dev.Use(middleware.AuthMiddleware())
	{
		// Project viewing (must come before tasks routes to avoid conflict)
		dev.GET("/projects/:id", devController.GetProject)
		dev.GET("/projects/:id/tasks", devController.ListProjectTasks)

		// Task operations
		dev.POST("/tasks", devController.CreateTask)
		dev.GET("/tasks/:id", devController.GetTask)
		dev.PUT("/tasks/:id", devController.UpdateTask)
	}

	// Admin only routes
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RoleMiddleware(models.RoleAdmin))
	{
		// Add admin-only endpoints here
		admin.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "admin endpoint"})
		})
	}

	return r
}
