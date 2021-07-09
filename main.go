package main

import (
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"
	// "happy_bank_simulator/ui"

	"gorm.io/gorm"
)

func main() {
	db := database.InitDB()
	db.AutoMigrate(
		&models.Borrower{},
		&models.Loan{},
		&models.Lender{},
		&models.Insurer{},
	)

	// Create test loan
	// startDate := "27/06/2021"
	// endDate := "27/06/2022"
	// duration := 12
	// amount := 10000

	// models.NewLoan(startDate, endDate, int32(duration), float64(amount))

	// dateString := "01/02/2021"
	// date, _ := time.Parse("02/01/2006", dateString)
	// fmt.Println(date.Format("2 Jan. 2006"))

	// ui.InitApp()
}

type MonthlyPayment struct {
	gorm.Model
	Loan     models.Loan
	Borrower models.Borrower
	Amount   float32
}
