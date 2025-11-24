package main

import (
	"log"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository/postgres"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/routes"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/pkg/db"
	"github.com/joho/godotenv"
)

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
