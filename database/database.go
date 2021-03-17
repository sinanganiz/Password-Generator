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
func GetPassword(id int) (password models.Password) {
	var password_record models.Password
	password_record.ID = uint(id)

	DBConn.Find(&password_record)

	return password_record
}
func UpdatePassword(id int, newPass string, newTitle string) {
	var password_record models.Password
	password_record.ID = uint(id)

	DBConn.Find(&password_record)

	password_record.Password = newPass
	password_record.PasswordTitle = newTitle
	DBConn.Save(&password_record)

}

func DeletePassword(id int) (result bool) {
	var password models.Password
	DBConn.Where("ID = ?", id).Delete(&password)
	return true
}
