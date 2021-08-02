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

type Loan struct {
	gorm.Model
	Borrower         Actor
	BorrowerID       uint
	Lenders          []*Actor `gorm:"many2many:loan_actors;"`
	Insurers         []*Actor `gorm:"many2many:loan_actors;"`
	StartDate        string
	EndDate          string
	Duration         int
	Amount           float64
	RefundedAmount   float64
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

func (instance *Loan) AddLender(lender *Actor) {
	database.GetDB().Model(&instance).Association("Lenders").Append(lender)
}

func (instance *Loan) AddInsurer(insurer *Actor) {
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
	fmt.Printf("- Duration: %s months, from %s to %s\n", strconv.Itoa(instance.Duration), instance.StartDate, instance.EndDate)
	fmt.Printf("- Amount: %1.2f € \n", instance.Amount)
	fmt.Printf("- Refunded amount: %1.2f €\n", instance.RefundedAmount)
	fmt.Printf("- Credit cost: %1.2f € (%1.2f %%)\n", instance.CreditCost(), instance.CreditRate*100)
	fmt.Printf("- Insurance cost: %1.2f € (%1.2f %%)\n", instance.InsuranceCost(), instance.InsuranceRate*100)
	fmt.Printf("- Loan cost: %1.2f €\n", instance.LoanCost())

	willFailText := "No"
	if instance.WillFailOn != "" {
		willFailText = fmt.Sprintf("Yes, on %s", instance.WillFailOn)
	}
	fmt.Println("- Will fail?", willFailText)

	borrowerName := "None"
	if instance.Borrower.Name != "" {
		name := instance.Borrower.Name
		id := instance.BorrowerID
		borrowerName = fmt.Sprintf("%s (#%s) 💰 %1.2f €", name, strconv.Itoa(int(id)), instance.Borrower.Balance)
	}
	fmt.Printf("- Borrower: %s\n", borrowerName)

	fmt.Printf("- %s lenders\n", strconv.Itoa(len(instance.Lenders)))
	for _, lender := range instance.Lenders {
		fmt.Printf("--- %s (#%s) 💰 %1.2f €\n", lender.Name, strconv.Itoa(int(lender.ID)), lender.Balance)
	}

	fmt.Println("- Is insured?", instance.IsInsured)
	if instance.IsInsured {
		fmt.Printf("- %s insurers\n", strconv.Itoa(len(instance.Insurers)))
		for _, insurer := range instance.Insurers {
			fmt.Printf("--- %s (#%s) 💰 %1.2f €\n", insurer.Name, strconv.Itoa(int(insurer.ID)), insurer.Balance)
		}
	}
	fmt.Printf("\n")
}

func (instance *Loan) Refund(amount float64) {
	instance.RefundedAmount += amount
	instance.Save()
}

func (instance *Loan) CreditCost() float64 {
	return (instance.MonthlyCredit*float64(instance.Duration) - instance.Amount)
}

func (instance *Loan) InsuranceCost() float64 {
	return instance.MonthlyInsurance * float64(instance.Duration)
}

func (instance *Loan) LoanCost() float64 {
	return instance.CreditCost() + instance.InsuranceCost()
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
	return gofin.PMT(interestCreditRate/12, float64(duration), float64(-amount), 0, 0)
}

func CalculateMonthlyInsurancePayment(interestInsuranceRate float64, amount float64) float64 {
	return (float64(interestInsuranceRate) * float64(amount) / 100) / 12
}
