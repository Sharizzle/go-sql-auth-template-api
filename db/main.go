package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//database global
var DB *gorm.DB

func SetupDB() *gorm.DB {

	//connect to db
	db, dbError := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if dbError != nil {
		panic("Failed to connect to database")
	}

	return db
}
