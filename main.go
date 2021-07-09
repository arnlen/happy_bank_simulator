package main

import (
	"fmt"
	"happy_bank_simulator/database"
	// "happy_bank_simulator/factories"
	"happy_bank_simulator/models"

	// "happy_bank_simulator/ui"

	"gorm.io/gorm"
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

	var borrower models.Borrower
	db.First(&borrower)
	fmt.Println(borrower)

	var borrowers []models.Borrower
	db.Find(&borrowers)
	fmt.Println(len(borrowers))

	// ui.InitApp()
}

type MonthlyPayment struct {
	gorm.Model
	Loan     models.Loan
	Borrower models.Borrower
	Amount   float32
}
