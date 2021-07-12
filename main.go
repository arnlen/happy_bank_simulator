package main

import (
	"happy_bank_simulator/database"
	"happy_bank_simulator/initializers"
	"happy_bank_simulator/models"
	"happy_bank_simulator/views"

	"gorm.io/gorm/clause"
)

func main() {
	initializers.InitDB()
	db := database.GetDB()

	var borrowers []models.Borrower
	var lenders []models.Lender
	var insurers []models.Insurer
	var loans []models.Loan

	db.Preload(clause.Associations).Find(&borrowers)
	db.Preload(clause.Associations).Find(&lenders)
	db.Preload(clause.Associations).Find(&insurers)
	db.Preload(clause.Associations).Find(&loans)

	views.InitApp(borrowers, lenders, insurers, loans)
}
