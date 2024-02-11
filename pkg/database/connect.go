package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"twitterClone/pkg/config"
)

func Connect() {
	co := config.Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", co.DB.Username, co.DB.Password, co.DB.Host, co.DB.Port, co.DB.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot Connect Database: ", err)
	}

	DB = db
}
