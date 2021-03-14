package database

import (
	"github.com/sinanganizz/password-generator/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	DBConn, err := gorm.Open(sqlite.Open("passwords.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DBConn.AutoMigrate(&models.Password{})
}
