package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback for local development if env not set
		dsn = "host=localhost user=postgres password=postgres dbname=taskmanager port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Connected to PostgreSQL database")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	// It will change existing column’s type if its size, precision, nullable changed.
	// It WON’T delete unused columns to protect your data.
	return db.AutoMigrate(
		&models.User{},
		&models.Workspace{},
		&models.Project{},
		&models.Task{},
		&models.TaskHistory{},
	)
}
