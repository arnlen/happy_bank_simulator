package models

import (
	"happy_bank_simulator/database"
	"happy_bank_simulator/services"
	"log"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	BorrowerID       uint
	LenderID         uint
	InsurerID        uint
	StartDate        string
	EndDate          string
	Duration         int32
	Amount           float64
	InitialDeposit   float64
	CreditRate       float64
	InsuranceRate    float64
	MonthlyCredit    float64
	MonthlyInsurance float64
}

func (instance *Loan) Save() *Loan {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func NewLoan(startDate string, endDate string, duration int32, amount float64, lender Lender, insurer Insurer) *Loan {
	initialDeposit := amount / 10
	creditRate := 0.3
	insuranceRate := 0.3
	monthlyCredit := services.CalculateMonthlyCreditPayment(creditRate, float64(duration), float64(amount))
	monthlyInsurance := services.CalculateMonthlyInsurancePayment(insuranceRate, float64(duration), float64(amount))

	return &Loan{
		InsurerID:        insurer.ID,
		LenderID:         lender.ID,
		StartDate:        startDate,
		EndDate:          endDate,
		Duration:         duration,
		Amount:           amount,
		InitialDeposit:   initialDeposit,
		CreditRate:       creditRate,
		InsuranceRate:    insuranceRate,
		MonthlyCredit:    monthlyCredit,
		MonthlyInsurance: monthlyInsurance,
	}
}

func CreateLoan(startDate string, endDate string, duration int32, amount float64, lender Lender, insurer Insurer) *Loan {
	loan := NewLoan(startDate, endDate, duration, amount, lender, insurer)
	result := database.GetDB().Create(&loan)

	if loan.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return loan
}
