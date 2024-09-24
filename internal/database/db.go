package database

import (
	"go-node/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=diary_user password=diary_password dbname=diary port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	

	// Миграция схемы
	err = db.AutoMigrate(&models.Node{})
	if err != nil {
		return nil, err
	}

	return db, nil
}