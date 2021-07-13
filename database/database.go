package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() *gorm.DB {
	db, err = gorm.Open(sqlite.Open("tmp/happy_dev.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
