package models

import "gorm.io/gorm"

type Password struct {
	gorm.Model
	Password      string `json:"password"`
	PasswordTitle string `json:"passwordTitle"`
}
