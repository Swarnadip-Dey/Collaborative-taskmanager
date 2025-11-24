package routes

import (
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/controllers"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(repo repository.Repository) *gin.Engine {
	r := gin.Default()

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize controllers
	authController := controllers.NewAuthController(repo)

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

	// Manager and Admin routes
	managerAdmin := r.Group("/api/manager")
	managerAdmin.Use(middleware.AuthMiddleware())
	managerAdmin.Use(middleware.RoleMiddleware(models.RoleAdmin, models.RoleManager))
	{
		// Add manager/admin endpoints here
		managerAdmin.GET("/projects", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "manager/admin endpoint"})
		})
	}

	return r
}
