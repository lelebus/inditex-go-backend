package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var database *gorm.DB

// InitDatabaseLayer initializes the database layer.
func InitDatabaseLayer() error {
	var err error
	database, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	// setup database
	AutoMigrate(database)

	return nil
}

// GetInstance returns the database instance.
func GetInstance() *gorm.DB {
	return database
}
