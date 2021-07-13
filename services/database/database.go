package database

import (
	"happy_bank_simulator/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	db, err = gorm.Open(sqlite.Open("tmp/happy_dev.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func MigrateDB() {
	modelList := models.GetModelList()
	db.AutoMigrate(modelList...)
}

func GetDB() *gorm.DB {
	return db
}

func SetDB(newDb *gorm.DB) {
	db = newDb
}

func DropBD() {
	for _, model := range models.GetModelList() {
		db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model)
	}
}
