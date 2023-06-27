package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/bookmanage?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{}) // Update Local Password
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
