package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(191)"`
	Email    string `gorm:"type:varchar(191);not null;unique"`
	Password string `gorm:"type:varchar(191);not null"`
}
