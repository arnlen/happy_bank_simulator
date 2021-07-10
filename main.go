package main

import (
	"happy_bank_simulator/database"
	// "happy_bank_simulator/factories"
	"happy_bank_simulator/models"
	"happy_bank_simulator/ui"

	"gorm.io/gorm/clause"
)

func main() {
	db := database.InitDB()

	db.AutoMigrate(
		&models.Borrower{},
		&models.Lender{},
		&models.Insurer{},
		&models.Loan{},
	)

	// factories.CreateSeedState()

	var borrowers []models.Borrower
	var lenders []models.Lender
	var insurers []models.Insurer
	var loans []models.Loan

	db.Preload(clause.Associations).Find(&borrowers)
	db.Preload(clause.Associations).Find(&lenders)
	db.Preload(clause.Associations).Find(&insurers)
	db.Preload(clause.Associations).Find(&loans)

	ui.InitApp(borrowers, lenders, insurers, loans)
}
