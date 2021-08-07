package database

import (
	"fmt"
	"happy_bank_simulator/internal/global"
	"happy_bank_simulator/internal/initializers"
	"happy_bank_simulator/models"

	"gorm.io/gorm"
)

var modelList = []interface{}{
	&models.Actor{},
	&models.Loan{},
	&models.Transaction{},
}

func SetupDB() {
	global.Db = initializers.InitDB()
	MigrateDB()
}

func ResetDB() {
	SetupDB()
	DropBD()
	SetupDB()
}

func MigrateDB() {
	global.Db.AutoMigrate(modelList...)
	fmt.Println("Database migrated")
}

func DropBD() {
	for _, model := range modelList {
		global.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model)
	}
	fmt.Println("Database dropped")
}
