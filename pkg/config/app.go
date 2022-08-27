package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

const username = "root"
const password = ""
const database_name = "bookstrore"

func Connect() {
	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/" + database_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
