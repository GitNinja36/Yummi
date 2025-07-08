package config

import (
	"log"
	"os"

	"github.com/GitNinja36/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Favorite{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
}
