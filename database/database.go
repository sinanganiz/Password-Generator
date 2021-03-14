package database

import (
	"github.com/sinanganizz/password-generator/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	err    error
)

func InitDatabase() {
	DBConn, err = gorm.Open(sqlite.Open("passwords.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DBConn.AutoMigrate(&models.Password{})
}

func SavePassword(password string, passwordTitle string) (result bool) {
	DBConn.Create(&models.Password{Password: password, PasswordTitle: passwordTitle})
	return true
}

func GetPasswords() (passwords []models.Password) {
	var password_records []models.Password
	DBConn.Find(&password_records)

	return password_records
}
