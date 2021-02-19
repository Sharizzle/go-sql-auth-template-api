package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"fmt"
)

//database global
var DB *gorm.DB

func SetupDB() *gorm.DB {

	//db config vars
	var dbHost string = os.Getenv("DB_HOST")
	var dbName string = os.Getenv("DB_NAME")
	var dbUser string = os.Getenv("DB_USERNAME")
	var dbPassword string = os.Getenv("DB_PASSWORD")

	fmt.Println(dbHost,dbName, dbUser, dbPassword )

	//connect to db
	db, dbError := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if dbError != nil {
		panic("Failed to connect to database")
	}

	return db
}
