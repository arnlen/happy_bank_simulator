package database

import (
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	modelList []interface{}
)

func GetDB() *gorm.DB {
	return db
}

func SetDB(newDb *gorm.DB) {
	db = newDb
}

func DropBD() {
	for _, model := range modelList {
		db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model)
	}
}

func SetModelList(newModelList []interface{}) {
	modelList = newModelList
}

func GetModelList() []interface{} {
	return modelList
}
