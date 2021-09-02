package models

import (
	"fmt"
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
	Lenders                     []*Actor `gorm:"many2many:loan_lenders;"`
	Insurers                    []*Actor `gorm:"many2many:loan_insurers;"`
	StartDate                   string
	Duration                    int
	Amount                      float64
	RefundedAmount              float64
	InterrestRate               float64
	InsuranceRate               float64
	SecurityDepositRate         float64
	IsInsured                   bool
	IsActive                    bool
	NumberOfMonthsBeforeFailure int
}

// ---------------------------------------
//
// ------- PUBLIC INSTANCE METHODS -------
//
// ---------------------------------------

// ORM METHODS
// ----------------

func (instance *Loan) Save() {
	global.Db.Save(instance)
	instance.Refresh()
}

func (instance *Loan) Refresh() {
	global.Db.Preload(clause.Associations).First(&instance, instance.ID)
}

// GETTER & SETTERS
// ----------------

func (instance *Loan) GetEndDate() string {
	startDate := helpers.ParseStringToDate(instance.StartDate)
	endDate := helpers.AddMonthsToDate(startDate, instance.Duration)
	return helpers.TimeDateToString(endDate)
}

func (instance *Loan) GetMonthlyAmountToRefund() float64 {
	return gofin.PMT(instance.InterrestRate/12, float64(instance.Duration), float64(-instance.Amount), 0, 0)
}

func (instance *Loan) GetMonthlyAmountToRefundPerLender() float64 {
	return instance.GetMonthlyAmountToRefund() / float64(len(instance.Lenders))
}

func (instance *Loan) GetMonthlyInsuranceCost() float64 {
	return (instance.InsuranceRate * instance.Amount / 100) / 12
}

func (instance *Loan) GetMonthlyInsuranceCostPerInsurer() float64 {
	return instance.GetMonthlyInsuranceCost() / float64(len(instance.Insurers))
}

func (instance *Loan) GetInitialSecurityDepositAmount() float64 {
	return instance.SecurityDepositRate * instance.Amount
}

func (instance *Loan) GetAmountLentPerLender() float64 {
	amountLentPerLender := instance.Amount / float64(len(instance.Lenders))
	fmt.Printf("Every lender lent %1.2f € through this loan.\n",
		amountLentPerLender)
	return amountLentPerLender
}

func (instance *Loan) GetAmountInsuredPerInsurer() float64 {
	amountInsuredPerLender := instance.Amount / float64(len(instance.Insurers))
	fmt.Printf("Every insurer is insuring %1.2f € of this loan.\n",
		amountInsuredPerLender)
	return amountInsuredPerLender
}

func (instance *Loan) SetRandomNumberOfMonthsBeforeFailure() {
	instance.NumberOfMonthsBeforeFailure = instance.generateRandomNumberWithinDuration()
	instance.Save()
}

func (instance *Loan) SetBorrowerMonthlyIncomes() {
	montlyIncomes := instance.requiredMontlyIncomes()
	instance.Borrower.UpdateMontlyIncomes(montlyIncomes)
}

// OTHER PUBLIC METHODS
// ----------------

func (instance *Loan) AddToRefund(amount float64) {
	instance.RefundedAmount += amount
	instance.Save()
}

func (instance *Loan) AssignBorrower(borrower *Actor) {
	instance.Borrower = *borrower
	instance.Save()
}

func (instance *Loan) AssignLender(lender *Actor) {
	instance.AssignLenders([]*Actor{lender})
}

func (instance *Loan) AssignLenders(lenders []*Actor) {
	instance.Lenders = append(instance.Lenders, lenders...)
}

func (instance *Loan) AssignInsurer(insurer *Actor) {
	instance.AssignInsurers([]*Actor{insurer})
}

func (instance *Loan) AssignInsurers(insurers []*Actor) {
	instance.Insurers = append(instance.Insurers, insurers...)
}

func (instance *Loan) Activate() {
	CreateDepositTransaction(instance.Borrower, instance.InitialDeposit)

	for _, lender := range instance.Lenders {
		amount := instance.GetAmountLentPerLender()
		CreateTransaction(*lender, instance.Borrower, amount)
		lender.Refresh()
	}

	instance.IsActive = true
	instance.Save()
	instance.Borrower.Refresh()
}

func (instance *Loan) Deactivate() {
	instance.IsActive = false
	instance.Save()
}

func (instance *Loan) WillFailOnTime() time.Time {
	startDate := helpers.ParseStringToDate(instance.StartDate)
	return helpers.AddMonthsToDate(startDate, instance.NumberOfMonthsBeforeFailure)
}

func (instance *Loan) WillFailOnString() string {
	failureDate := instance.WillFailOnTime()
	return failureDate.Format("01/2006")
}

func (instance *Loan) SetupLenders() {
	instance.setupActor(configs.Actor.LenderString)
}

func (instance *Loan) SetupInsurers() {
	instance.setupActor(configs.Actor.InsurerString)
}

func (instance *Loan) CreateMontlyTransactionsFromBorrowerToLenders() {
	for _, lender := range instance.Lenders {
		transaction := CreateTransaction(instance.Borrower, *lender, instance.MonthlyAmountToRefund)
		transaction.Print()
		instance.AddToRefund(transaction.Amount)
	}
}

func (instance *Loan) CreateMontlyTransactionsFromLendersToInsurers() {
	if !instance.IsInsured {
		fmt.Printf("Loan #%d isn't insured.", int(instance.ID))
		return
	}

	amoutToPayPerInsurerByEveryLender := instance.MonthlyInsuranceCost / float64(len(instance.Insurers))
	for _, insurer := range instance.Insurers {
		for _, lender := range instance.Lenders {
			transaction := CreateTransaction(*lender, *insurer, amoutToPayPerInsurerByEveryLender)
			transaction.Print()
		}
	}
}

func (instance *Loan) ShouldFailThisMonth(date time.Time) bool {
	return date.Equal(instance.WillFailOnTime())
}

// ----------------------------
//
// ------- PRINT METHOD -------
//
// ----------------------------

func (instance *Loan) Print() {
	fmt.Printf("\n---[ LOAN #%s ]---\n", strconv.Itoa(int(instance.ID)))
	fmt.Printf("- Duration: %s months, from %s to %s\n", strconv.Itoa(instance.Duration), instance.StartDate, instance.GetEndDate())
	fmt.Printf("- Amount: %1.2f € \n", instance.Amount)
	fmt.Printf("- Refunded amount: %1.2f €\n", instance.RefundedAmount)
	fmt.Printf("- Credit cost: %1.2f € (%1.2f %%)\n", instance.totalCreditCost(), instance.InterrestRate*100)
	fmt.Printf("- Insurance cost: %1.2f € (%1.2f %%)\n", instance.totalInsuranceCost(), instance.InsuranceRate*100)
	fmt.Printf("- Loan cost: %1.2f €\n", instance.totalLoanCost())

	willFailText := "No"
	if instance.willFail() {
		willFailText = fmt.Sprintf("Yes, on %s", instance.WillFailOnString())
	}
	fmt.Println("- Will fail?", willFailText)

	borrowerName := "None"
	if instance.Borrower.Name != "" {
		borrowerName = fmt.Sprintf("%s (#%d) 💰 %1.2f €",
			instance.Borrower.Name,
			int(instance.BorrowerID),
			instance.Borrower.Balance,
		)
	}
	fmt.Printf("- Borrower: %s\n", borrowerName)

	fmt.Printf("- %d lenders\n", len(instance.Lenders))
	for _, lender := range instance.Lenders {
		fmt.Printf("--- %s (#%d) 💰 %1.2f €\n",
			lender.Name,
			int(lender.ID),
			lender.Balance,
		)
	}

	fmt.Println("- Is insured?", instance.IsInsured)
	if instance.IsInsured {
		fmt.Printf("- %d insurers\n", len(instance.Insurers))
		for _, insurer := range instance.Insurers {
			fmt.Printf("--- %s (#%d) 💰 %1.2f €\n",
				insurer.Name,
				int(insurer.ID),
				insurer.Balance,
			)
		}
	}

	isActive := "No"
	if instance.IsActive {
		isActive = fmt.Sprintln("Yes")
	}
	fmt.Println("- Is active?", isActive)

	fmt.Printf("\n")
}

// ---------------------------------------
//
// ------- PRIVATE INSTANCE METHODS -------
//
// ---------------------------------------

func (instance *Loan) getLenderQuantityRequired() int {
	quantity := int(math.Ceil(instance.Amount / configs.Actor.MaxAmountPerLoan))
	fmt.Printf("%s lenders are required\n", strconv.Itoa(quantity))
	return quantity
}

func (instance *Loan) getInsurerQuantityRequired() int {
	quantity := int(math.Ceil(instance.Amount / configs.Actor.MaxAmountPerLoan))
	fmt.Printf("%s insurers are required\n", strconv.Itoa(quantity))
	return quantity
}

func (instance *Loan) totalCreditCost() float64 {
	return (instance.MonthlyAmountToRefund*float64(instance.Duration) - instance.Amount)
}

func (instance *Loan) totalInsuranceCost() float64 {
	return instance.MonthlyInsuranceCost * float64(instance.Duration)
}

func (instance *Loan) totalLoanCost() float64 {
	return instance.totalCreditCost() + instance.totalInsuranceCost()
}

func (instance *Loan) generateRandomNumberWithinDuration() int {
	return rand.Intn(instance.Duration)
}

func (instance *Loan) willFail() bool {
	return instance.NumberOfMonthsBeforeFailure > 0
}

func (instance *Loan) assignActors(actors []*Actor) {
	for _, actor := range actors {
		switch actor.Type {
		case configs.Actor.BorrowerString:
			instance.AssignBorrower(actor)
		case configs.Actor.LenderString:
			instance.AssignLender(actor)
		case configs.Actor.InsurerString:
			instance.AssignInsurer(actor)
		}
	}
}

func (instance *Loan) requiredMontlyIncomes() float64 {
	if instance.willFail() {
		amountToRefundUntilFailure := instance.GetTotalMonthlyPayment() * float64(instance.NumberOfMonthsBeforeFailure)
		totalIncomesUntilFailure := amountToRefundUntilFailure - instance.Borrower.Balance
		return totalIncomesUntilFailure / float64(instance.NumberOfMonthsBeforeFailure)
	}

	return instance.GetTotalMonthlyPayment()
}

func (instance *Loan) setupActor(actorType string) {
	fmt.Printf("Setup %ss for Loan #%s:\n",
		actorType,
		strconv.Itoa(int(instance.ID)),
	)

	var actorThatCanTakeThisLoan []*Actor

	for _, actor := range ListActors(actorType) {
		if actor.CanTakeThisLoan(*instance) {
			actorThatCanTakeThisLoan = append(actorThatCanTakeThisLoan, actor)
		}
	}
	fmt.Printf("- %d %s can take this loan\n",
		len(actorThatCanTakeThisLoan),
		actorType,
	)

	var actorQuantityRequired int
	switch actorType {
	case configs.Actor.LenderString:
		actorQuantityRequired = instance.getLenderQuantityRequired()
	case configs.Actor.InsurerString:
		actorQuantityRequired = instance.getInsurerQuantityRequired()
	}

	missingActorsQuantity := actorQuantityRequired - len(actorThatCanTakeThisLoan)
	if missingActorsQuantity > 0 {
		fmt.Printf("- Not enough available %s: missing %d\n",
			actorType,
			missingActorsQuantity,
		)
		newActors := CreateActors(actorType, missingActorsQuantity)
		fmt.Printf("- %d new %s created\n",
			len(newActors),
			actorType,
		)
		actorThatCanTakeThisLoan = append(actorThatCanTakeThisLoan, newActors...)
	}

	instance.assignActors(actorThatCanTakeThisLoan)
	fmt.Printf("- %d %s assigned to loan\n",
		len(actorThatCanTakeThisLoan),
		actorType,
	)

	var loanActors []*Actor
	switch actorType {
	case configs.Actor.LenderString:
		loanActors = instance.Lenders
	case configs.Actor.InsurerString:
		loanActors = instance.Insurers
	}
	fmt.Printf("=> Loan #%d has now %d %s assigned: ",
		int(instance.ID),
		len(loanActors),
		actorType,
	)

	for _, actor := range loanActors {
		fmt.Printf("#%s, ", strconv.Itoa(int(actor.ID)))
	}
	fmt.Printf("\n")
}

// --------------------------------
//
// ------- PACKAGE METHODS -------
//
// --------------------------------

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
