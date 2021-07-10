package main

import (
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"

	"happy_bank_simulator/ui"
)

func main() {
	db := database.InitDB()

	db.AutoMigrate(
		&models.Borrower{},
		&models.Lender{},
		&models.Insurer{},
		&models.Loan{},
	)

	var borrowers []models.Borrower
	var lenders []models.Lender
	var insurers []models.Insurer
	var loans []models.Loan

	db.Find(&borrowers)
	db.Find(&lenders)
	db.Find(&insurers)
	db.Find(&loans)

	ui.InitApp(borrowers, lenders, insurers, loans)
}
