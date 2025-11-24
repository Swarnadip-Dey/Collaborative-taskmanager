package main

import (
	"log"
	"time"

	_ "github.com/Swarnadip-Dey/Collaborative-taskmanager/docs"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository/postgres"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/routes"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/services"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/pkg/db"
	"github.com/joho/godotenv"
)

// @title Collaborative Task Manager API
// @version 1.0
// @description API for managing collaborative tasks with role-based access control
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@taskmanager.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Database connection
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Start health monitor (ping DB every minute)
	services.StartHealthMonitor(database, time.Minute)

	// Auto Migration
	if err := db.Migrate(database); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Repository initialization
	repo := postgres.NewRepository(database)

	// Setup routes
	r := routes.SetupRouter(repo)

	// Start server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
