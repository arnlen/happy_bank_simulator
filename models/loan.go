package models

import (
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"
	"happy_bank_simulator/helpers"
	"log"
	"math/rand"
	"strconv"

	"github.com/drum445/gofin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Declare conformity with Actor interface
var _ ModelBase = (*Loan)(nil)

type Loan struct {
	gorm.Model
	Borrower         Borrower
	BorrowerID       uint
	Lenders          []*Lender  `gorm:"many2many:loan_lenders;"`
	Insurers         []*Insurer `gorm:"many2many:loan_insurers;"`
	StartDate        string
	EndDate          string
	Duration         int
	Amount           int
	InitialDeposit   int
	CreditRate       float64
	InsuranceRate    float64
	MonthlyCredit    float64
	MonthlyInsurance float64
	WillFailOn       string
}

// ------- Instance methods -------

func (instance *Loan) ModelName() string {
	return "emprunt"
}

func (instance *Loan) Refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

func (instance *Loan) Save() {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Loan) AddLender(lender *Lender) {
	database.GetDB().Model(&instance).Association("Lenders").Append(lender)
}

func (instance *Loan) AddInsurer(insurer *Insurer) {
	database.GetDB().Model(&instance).Association("Insurers").Append(insurer)
}

func (instance *Loan) SetRandomFailureDate() {
	startDate := helpers.ParseStringToDate(instance.StartDate)
	numberOfMonthsBeforeFailure := rand.Intn(instance.Duration)
	failureDate := helpers.AddMonthsToDate(startDate, numberOfMonthsBeforeFailure)
	instance.WillFailOn = failureDate.Format("01/2006")

	fmt.Printf("The failure will occure after %s months, on %s\n", strconv.Itoa(numberOfMonthsBeforeFailure), instance.WillFailOn)
	instance.Save()
}

// ------- Package methods -------

func ListLoans() []Loan {
	var loans []Loan
	database.GetDB().Preload(clause.Associations).Find(&loans)
	return loans
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
		StartDate:        configs.General.StartDate,
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
