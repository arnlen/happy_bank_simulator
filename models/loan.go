package models

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/helpers"
	"happy_bank_simulator/internal/global"

	"github.com/drum445/gofin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Loan struct {
	gorm.Model
	Borrower                    Actor
	BorrowerID                  uint
	Lenders                     []*Actor `gorm:"many2many:loan_actors;"`
	Insurers                    []*Actor `gorm:"many2many:loan_actors;"`
	StartDate                   string
	Duration                    int
	Amount                      float64
	RefundedAmount              float64
	InitialDeposit              float64
	CreditRate                  float64
	InsuranceRate               float64
	MonthlyCredit               float64
	MonthlyInsurance            float64
	IsInsured                   bool
	IsActive                    bool
	NumberOfMonthsBeforeFailure int
}

// ------- Instance methods -------

func (instance *Loan) EndDate() string {
	startDate := helpers.ParseStringToDate(instance.StartDate)
	endDate := helpers.AddMonthsToDate(startDate, instance.Duration)
	return helpers.TimeDateToString(endDate)
}

func (instance *Loan) Save() {
	result := global.Db.Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Loan) Refresh() {
	global.Db.Preload(clause.Associations).First(&instance, instance.ID)
}

func (instance *Loan) AddLender(lender *Actor) {
	global.Db.Model(&instance).Association("Lenders").Append(lender)
}

func (instance *Loan) AddInsurer(insurer *Actor) {
	global.Db.Model(&instance).Association("Insurers").Append(insurer)
}

func (instance *Loan) Activate() {
	CreateDepositTransaction(instance.Borrower, instance.InitialDeposit)

	for _, lender := range instance.Lenders {
		amount := instance.AmountPerLender()
		CreateTransaction(*lender, instance.Borrower, amount)
		lender.Refresh()
	}

	instance.IsActive = true
	instance.Save()
	instance.Borrower.Refresh()
}

func (instance *Loan) SetRandomNumberOfMonthsBeforeFailure() {
	instance.NumberOfMonthsBeforeFailure = instance.generateRandomNumberWithinDuration()
	instance.Save()
}

func (instance *Loan) WillFail() bool {
	return instance.NumberOfMonthsBeforeFailure > 0
}

func (instance *Loan) WillFailOnTime() time.Time {
	startDate := helpers.ParseStringToDate(instance.StartDate)
	return helpers.AddMonthsToDate(startDate, instance.NumberOfMonthsBeforeFailure)
}

func (instance *Loan) WillFailOnString() string {
	failureDate := instance.WillFailOnTime()
	return failureDate.Format("01/2006")
}

func (instance *Loan) SetBorrowerMonthlyIncomes() {
	montlyIncomes := instance.RequiredMontlyIncomes()
	instance.Borrower.UpdateMontlyIncomes(montlyIncomes)
}

func (instance *Loan) AssignBorrower(borrower *Actor) {
	instance.Borrower = *borrower
}

func (instance *Loan) AssignLender(lender *Actor) {
	instance.Lenders = append(instance.Lenders, lender)
}

func (instance *Loan) AssignInsurer(insurer *Actor) {
	instance.Insurers = append(instance.Insurers, insurer)
}

func (instance *Loan) RequiredMontlyIncomes() float64 {
	if instance.WillFail() {
		amountToRefundUntilFailure := instance.monthlyPaymentDue() * float64(instance.NumberOfMonthsBeforeFailure)
		totalIncomesUntilFailure := amountToRefundUntilFailure - instance.Borrower.Balance
		return totalIncomesUntilFailure / float64(instance.NumberOfMonthsBeforeFailure)
	}

	return instance.monthlyPaymentDue()
}

func (instance *Loan) lenderQuantityRequired() int {
	quantity := int(math.Ceil(instance.Amount / configs.Actor.MaxAmountPerLoan))
	fmt.Printf("%s lenders are required\n", strconv.Itoa(quantity))
	return quantity
}

func (instance *Loan) insurerQuantityRequired() int {
	quantity := int(math.Ceil(instance.Amount / configs.Actor.MaxAmountPerLoan))
	fmt.Printf("%s insurers are required\n", strconv.Itoa(quantity))
	return quantity
}

func (instance *Loan) AmountPerLender() float64 {
	amount := instance.Amount / float64(instance.lenderQuantityRequired())
	fmt.Printf("%1.2f €/lender.\n", amount)
	return amount
}

func (instance *Loan) AmountPerInsurer() float64 {
	return instance.Amount / float64(instance.insurerQuantityRequired())
}

func (instance *Loan) Print() {
	fmt.Printf("\n---[ LOAN #%s ]---\n", strconv.Itoa(int(instance.ID)))
	fmt.Printf("- Duration: %s months, from %s to %s\n", strconv.Itoa(instance.Duration), instance.StartDate, instance.EndDate())
	fmt.Printf("- Amount: %1.2f € \n", instance.Amount)
	fmt.Printf("- Refunded amount: %1.2f €\n", instance.RefundedAmount)
	fmt.Printf("- Credit cost: %1.2f € (%1.2f %%)\n", instance.totalCreditCost(), instance.CreditRate*100)
	fmt.Printf("- Insurance cost: %1.2f € (%1.2f %%)\n", instance.totalInsuranceCost(), instance.InsuranceRate*100)
	fmt.Printf("- Loan cost: %1.2f €\n", instance.totalLoanCost())

	willFailText := "No"
	if instance.WillFail() {
		willFailText = fmt.Sprintf("Yes, on %s", instance.WillFailOnString())
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

func (instance *Loan) totalCreditCost() float64 {
	return (instance.MonthlyCredit*float64(instance.Duration) - instance.Amount)
}

func (instance *Loan) totalInsuranceCost() float64 {
	return instance.MonthlyInsurance * float64(instance.Duration)
}

func (instance *Loan) monthlyPaymentDue() float64 {
	return instance.MonthlyCredit + instance.MonthlyInsurance
}

func (instance *Loan) totalLoanCost() float64 {
	return instance.totalCreditCost() + instance.totalInsuranceCost()
}

func (instance *Loan) generateRandomNumberWithinDuration() int {
	return rand.Intn(instance.Duration)
}

// ------- Package methods -------

func ListLoans() []Loan {
	var loans []Loan
	global.Db.Preload(clause.Associations).Find(&loans)
	return loans
}

func ListActiveLoans() []Loan {
	var loans []Loan
	global.Db.Preload(clause.Associations).Where("is_active = ?", true).Find(&loans)
	return loans
}

func newDefaultLoan() *Loan {
	amount := configs.Loan.DefaultAmount
	startDate := configs.General.StartDate
	duration := configs.Loan.DefaultDuration
	creditRate := configs.General.CreditInterestRate
	insuranceRate := configs.General.InsuranceInterestRate

	initialDeposit := configs.Loan.SecurityDepositRate * float64(amount)
	monthlyCredit := calculateMonthlyCreditPayment(creditRate, duration, amount)
	monthlyInsurance := calculateMonthlyInsurancePayment(insuranceRate, amount)

	return &Loan{
		StartDate:        startDate,
		Duration:         configs.Loan.DefaultDuration,
		Amount:           configs.Loan.DefaultAmount,
		InitialDeposit:   initialDeposit,
		CreditRate:       creditRate,
		InsuranceRate:    insuranceRate,
		MonthlyCredit:    monthlyCredit,
		MonthlyInsurance: monthlyInsurance,
		IsActive:         false,
	}
}

func CreateDefaultLoan() *Loan {
	var loan = newDefaultLoan()
	loan.Save()
	return loan
}

func calculateMonthlyCreditPayment(interestCreditRate float64, duration int, amount float64) float64 {
	return gofin.PMT(interestCreditRate/12, float64(duration), float64(-amount), 0, 0)
}

func calculateMonthlyInsurancePayment(interestInsuranceRate float64, amount float64) float64 {
	return (float64(interestInsuranceRate) * float64(amount) / 100) / 12
}
