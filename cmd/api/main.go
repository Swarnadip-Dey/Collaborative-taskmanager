package main

import (
	"log"
	"net/http"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository/postgres"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/pkg/db"
	"github.com/gin-gonic/gin"
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
	_ = repo // Prevent unused variable error for now

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
