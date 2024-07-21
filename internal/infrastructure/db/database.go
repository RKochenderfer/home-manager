package db

import (
	"home-manager/server/internal/core/shared"
	"home-manager/server/internal/infrastructure/db/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
}

func Init() Database {
	connString := os.Getenv(shared.EnvConnectionString)
	env := os.Getenv(shared.Environment)
	db, err := gorm.Open(sqlite.Open(connString))
	if err != nil {
		panic("Failed to connect to database")
	}

	if env == shared.Development {
		performMigrations(db)
	}

	return Database{db}
}

func performMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Room{}); err != nil {
		println("Failed to migrate Room")
	}
	if err := db.AutoMigrate(&models.Chore{}); err != nil {
		println("Failed to migrate Chore")
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		println("Failed to migrate User")
	}
	if err := db.AutoMigrate(&models.Assignment{}); err != nil {
		println("Failed to migrate Assignment")
	}
	if err := db.AutoMigrate(&models.Redemption{}); err != nil {
		println("Failed to migrate redemption")
	}
	if err := db.AutoMigrate(&models.Reward{}); err != nil {
		println("Failed to migrate Reward")
	}
}
