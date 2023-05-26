package database

import (
	"github.com/jinzhu/gorm"
	"github.com/lelebus/inditext-go-backend/internal/domain"
)

// AutoMigrate runs the auto migration tool for the database.
func AutoMigrate(db *gorm.DB) {
	// Here you should add all models...
	db.AutoMigrate(&domain.Brand{})
	db.AutoMigrate(&domain.Price{})
}
