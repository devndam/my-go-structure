package migrations

import (
	"github.com/devndam/go-starter/app/models"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) error {
	// Create the table for the User model
	return db.AutoMigrate(&models.User{})
}

func Down(db *gorm.DB) error {
	// Drop the table for the User model
	return db.Migrator().DropTable(&models.User{})
}
