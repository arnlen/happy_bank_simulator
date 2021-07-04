package main

import (
	"fmt"
	"log"
	"time"

	"github.com/drum445/gofin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// Gorm init
	db, err := gorm.Open(sqlite.Open("happy_dev.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&Loan{},
		&Borrower{},
	)

	// Create test loan
	testLoan := Loan{
		Model:            gorm.Model{},
		StartDate:        "27/06/2021",
		EndDate:          "27/06/2022",
		Duration:         12,
		Amount:           1000,
		InitialDeposit:   100,
		MonthlyCredit:    1.2,
		MonthlyInsurance: 12.4,
	}

	fmt.Printf("%+v\n", testLoan)

	dateString := "01/02/2021"
	date, _ := time.Parse("02/01/2006", dateString)
	fmt.Println(date.Format("2 Jan. 2006"))
}

func CalculateMonthlyCreditPayment(loanInterestCreditRate float64, loanDuration float64, loanAmount float64) float64 {
	return gofin.PMT(loanInterestCreditRate, loanDuration, -loanAmount, 0, 0)
}

func CalculateMonthlyInsurancePayment(loanInterestInsuranceRate float64, loanDuration float64, loanAmount float64) float64 {
	return (loanInterestInsuranceRate * loanAmount / 100) / 12
}

type Borrower struct {
	gorm.Model
	Loans []Loan
}

type Loan struct {
	gorm.Model
	BorrowerID       uint
	StartDate        string
	EndDate          string
	Duration         int32
	Amount           float32
	InitialDeposit   float32
	MonthlyCredit    float32
	MonthlyInsurance float32
}

type MonthlyPayment struct {
	gorm.Model
	Loan     Loan
	Borrower Borrower
	Amount   float32
}
