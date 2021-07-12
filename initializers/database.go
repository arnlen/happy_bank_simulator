package initializers

import (
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open("database/happy_dev.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	modelList := []interface{}{&models.Borrower{}, &models.Lender{}, &models.Insurer{}, &models.Loan{}}
	database.SetModelList(modelList)
	db.AutoMigrate(modelList...)

	database.SetDB(db)
}
