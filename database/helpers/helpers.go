package helpers

import (
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"

	"gorm.io/gorm"
)

var modelList = []interface{}{
	&models.Borrower{},
	&models.Insurer{},
	&models.Lender{},
	&models.Loan{},
	&models.Transaction{},
}

func GetModelList() []interface{} {
	return modelList
}

func MigrateDB() {
	modelList := GetModelList()
	database.GetDB().AutoMigrate(modelList...)
}

func DropBD() {
	for _, model := range GetModelList() {
		database.GetDB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model)
	}
}
