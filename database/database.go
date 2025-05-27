package database

import (
	"fmt"
	"log"

	"github.com/migvas/go-tasks-api/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate all your models
	err = db.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	fmt.Println("Database schema migration completed successfully!")

	return db
}
