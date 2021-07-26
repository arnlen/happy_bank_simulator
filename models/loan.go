package models

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"
	"happy_bank_simulator/helpers"

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
	Amount           float64
	InitialDeposit   float64
	CreditRate       float64
	InsuranceRate    float64
	MonthlyCredit    float64
	MonthlyInsurance float64
	IsInsured        bool
	IsActive         bool
	WillFailOn       string
}

// ------- Instance methods -------

func (instance *Loan) ModelName() string {
	return "loan"
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

func (instance *Loan) SetRandomFailureDate() int {
	startDate := helpers.ParseStringToDate(instance.StartDate)
	numberOfMonthsBeforeFailure := rand.Intn(instance.Duration)
	failureDate := helpers.AddMonthsToDate(startDate, numberOfMonthsBeforeFailure)
	instance.WillFailOn = failureDate.Format("01/2006")
	instance.Save()

	return numberOfMonthsBeforeFailure
}

func (instance *Loan) Print() {
	fmt.Printf("\n---[ LOAN #%s ]---\n", strconv.Itoa(int(instance.ID)))
	fmt.Println("- Start date:", instance.StartDate)
	fmt.Println("- End date:", instance.EndDate)
	fmt.Println("- Amount:", instance.Amount, "€")

	willFailText := "No"
	if instance.WillFailOn != "" {
		willFailText = fmt.Sprintf("Yes, on %s", instance.WillFailOn)
	}
	fmt.Println("- Will fail?", willFailText)

	borrowerName := "None"
	if instance.Borrower.Name != "" {
		name := instance.Borrower.Name
		id := instance.BorrowerID
		borrowerName = fmt.Sprintf("%s (#%s)", name, strconv.Itoa(int(id)))
	}
	fmt.Printf("- Borrower: %s\n", borrowerName)

	fmt.Printf("- %s lenders\n", strconv.Itoa(len(instance.Lenders)))
	for _, lender := range instance.Lenders {
		fmt.Printf("--- %s (#%s)\n", lender.Name, strconv.Itoa(int(lender.ID)))
	}

	fmt.Println("- Is insured?", instance.IsInsured)
	if instance.IsInsured {
		fmt.Printf("- %s insurers\n", strconv.Itoa(len(instance.Insurers)))
		for _, insurer := range instance.Insurers {
			fmt.Printf("--- %s (#%s)\n", insurer.Name, strconv.Itoa(int(insurer.ID)))
		}
	}
	fmt.Printf("\n")
}

// ------- Package methods -------

func ListLoans() []Loan {
	var loans []Loan
	database.GetDB().Preload(clause.Associations).Find(&loans)
	return loans
}

func ListActiveLoans() []Loan {
	var loans []Loan
	database.GetDB().Preload(clause.Associations).Where("is_active = ?", true).Find(&loans)
	return loans
}

func NewDefaultLoan() *Loan {
	amount := configs.Loan.DefaultAmount
	startDate := configs.General.StartDate
	duration := configs.Loan.DefaultDuration
	creditRate := configs.General.CreditInterestRate
	insuranceRate := configs.General.InsuranceInterestRate

	initialDeposit := configs.Loan.SecurityDepositRate * float64(amount)
	monthlyCredit := CalculateMonthlyCreditPayment(creditRate, duration, amount)
	monthlyInsurance := CalculateMonthlyInsurancePayment(insuranceRate, amount)
	endDate := helpers.TimeDateToString(helpers.AddMonthsToDate(helpers.ParseStringToDate(startDate), duration))

	return &Loan{
		StartDate:        startDate,
		Duration:         configs.Loan.DefaultDuration,
		EndDate:          endDate,
		Amount:           configs.Loan.DefaultAmount,
		InitialDeposit:   initialDeposit,
		CreditRate:       creditRate,
		InsuranceRate:    insuranceRate,
		MonthlyCredit:    monthlyCredit,
		MonthlyInsurance: monthlyInsurance,
		IsActive:         true,
	}
}

func CreateEmptyLoan() *Loan {
	var loan = NewDefaultLoan()
	loan.Save()
	return loan
}

func CalculateMonthlyCreditPayment(interestCreditRate float64, duration int, amount float64) float64 {
	return gofin.PMT(interestCreditRate, float64(duration), float64(-amount), 0, 0)
}

func CalculateMonthlyInsurancePayment(interestInsuranceRate float64, amount float64) float64 {
	return (float64(interestInsuranceRate) * float64(amount) / 100) / 12
}
