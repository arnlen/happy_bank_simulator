package models

import (
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"
	"log"

	"github.com/drum445/gofin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Loan struct {
	gorm.Model
	Borrower         Borrower
	BorrowerID       uint
	Lender           Lender
	LenderID         uint
	Insurer          Insurer
	InsurerID        uint
	StartDate        string
	EndDate          string
	Duration         int
	Amount           int
	InitialDeposit   int
	CreditRate       float64
	InsuranceRate    float64
	MonthlyCredit    float64
	MonthlyInsurance float64
}

func (instance *Loan) ModelName() string {
	return "emprunt"
}

func ListLoans() []Loan {
	var loans []Loan
	database.GetDB().Preload(clause.Associations).Find(&loans)
	return loans
}

func (instance *Loan) Save() *Loan {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func NewDefaultLoan() *Loan {
	amount := configs.Loan.DefaultAmount
	duration := configs.Loan.DefaultDuration
	creditRate := configs.General.CreditInterestRate
	insuranceRate := configs.General.InsuranceInterestRate

	initialDeposit := configs.Loan.SecurityDepositRate * float64(amount)
	monthlyCredit := CalculateMonthlyCreditPayment(creditRate, duration, amount)
	monthlyInsurance := CalculateMonthlyInsurancePayment(insuranceRate, amount)

	return &Loan{
		Duration:         configs.Loan.DefaultDuration,
		Amount:           configs.Loan.DefaultAmount,
		InitialDeposit:   int(initialDeposit),
		CreditRate:       creditRate,
		InsuranceRate:    insuranceRate,
		MonthlyCredit:    monthlyCredit,
		MonthlyInsurance: monthlyInsurance,
	}
}

func CalculateMonthlyCreditPayment(interestCreditRate float64, duration int, amount int) float64 {
	return gofin.PMT(interestCreditRate, float64(duration), float64(-amount), 0, 0)
}

func CalculateMonthlyInsurancePayment(interestInsuranceRate float64, amount int) float64 {
	return (float64(interestInsuranceRate) * float64(amount) / 100) / 12
}
