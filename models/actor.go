package models

import (
	"fmt"
	"log"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/internal/global"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"syreclabs.com/go/faker"
)

type Actor struct {
	gorm.Model
	Name           string
	Loans          []*Loan `gorm:"many2many:loan_actors;"`
	InitialBalance float64
	Balance        float64
	MonthlyIncomes float64
	Type           string
}

// ------- Instance methods -------

func (instance *Actor) Save() {
	result := global.Db.Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Actor) Refresh() {
	global.Db.Preload(clause.Associations).Find(&instance)
}

func (instance *Actor) NetBalance() float64 {
	if instance.IsBorrower() || instance.IsInsurer() {
		return instance.Balance - instance.TotalAmountAssigned()
	}

	return instance.Balance
}

func (instance *Actor) TotalAmountAssigned() float64 {
	loans := instance.Loans
	totalAmoutAssigned := 0.0

	for _, loan := range loans {
		totalAmoutAssigned += loan.Amount
	}
	return totalAmoutAssigned
}

func (instance *Actor) UpdateBalance(amount float64) {
	instance.Balance += amount
	instance.Save()
}

func (instance *Actor) UpdateMontlyIncomes(amount float64) {
	if !instance.IsBorrower() {
		log.Fatalf("Tried to call UpdateMontlyIncomes on %s #%s (not a borrower)\n",
			instance.Type, strconv.Itoa(int(instance.ID)))
	}

	instance.MonthlyIncomes = amount
	instance.Save()
}

func (instance *Actor) AssignLoan(loan *Loan) {
	instance.Loans = append(instance.Loans, loan)
	instance.Save()

	switch instance.Type {
	case "borrower":
		loan.AssignBorrower(instance)
	case "lender":
		loan.AssignLender(instance)
	case "insurer":
		loan.AssignInsurer(instance)
	}
}

func (instance *Actor) CanTakeThisLoan(loan Loan) bool {
	if instance.IsBorrower() {
		loans := instance.Loans
		totalAmountBorrowed := instance.TotalAmountAssigned()
		fmt.Printf("Borrower #%s has %s loans for a total of %1.2f €\n",
			strconv.Itoa(int(instance.ID)), strconv.Itoa(len(loans)), totalAmountBorrowed)

		borrowerNetBalance := instance.NetBalance()
		fmt.Printf("Borrower #%s net balance is %1.2f €\n", strconv.Itoa(int(instance.ID)), borrowerNetBalance)

		ratio := float64(borrowerNetBalance / loan.Amount)
		if ratio >= configs.Actor.BorrowerBalanceLeverageRatio {
			fmt.Printf("Borrower #%s can take the loan #%s\n",
				strconv.Itoa(int(instance.ID)), strconv.Itoa(int(loan.ID)))
			return true
		} else {
			fmt.Printf("Borrower #%s cannot take the loan #%s: net balance to low\n",
				strconv.Itoa(int(instance.ID)), strconv.Itoa(int(loan.ID)))
			return false
		}
	}

	if instance.IsLender() {
		amountToLend := loan.AmountPerLender()
		return instance.NetBalance() >= amountToLend
	}

	amountToInsure := loan.AmountPerInsurer()
	return instance.NetBalance() >= amountToInsure
}

func (instance *Actor) IsBorrower() bool {
	return instance.Type == configs.Actor.BorrowerString
}

func (instance *Actor) IsLender() bool {
	return instance.Type == configs.Actor.LenderString
}

func (instance *Actor) IsInsurer() bool {
	return instance.Type == configs.Actor.InsurerString
}

// ------- Package methods -------

func ListActors(actorType string) []*Actor {
	var actors []*Actor
	global.Db.Preload(clause.Associations).Where("Type = ?", actorType).Find(&actors)
	return actors
}

func ListActorsWithPositiveBalance(actorType string) []*Actor {
	actors := ListActors(actorType)
	var actorsWithPositiveBalance []*Actor
	for _, actor := range actors {
		if actor.Balance > 0 {
			actorsWithPositiveBalance = append(actorsWithPositiveBalance, actor)
		}
	}
	return actorsWithPositiveBalance
}

func ListActorsWithoutLoan(actorType string) []*Actor {
	var actorsWithoutLoan []*Actor
	for _, actor := range ListActors(actorType) {
		if len(actor.Loans) == 0 {
			actorsWithoutLoan = append(actorsWithoutLoan, actor)
		}
	}
	return actorsWithoutLoan
}

func ListActorsWithLoan(actorType string) []*Actor {
	var actorsWithLoan []*Actor
	for _, actor := range ListActors(actorType) {
		if len(actor.Loans) != 0 {
			actorsWithLoan = append(actorsWithLoan, actor)
		}
	}
	return actorsWithLoan
}

func ListActorsWithLoanOtherThanTarget(actorType string, loan *Loan) []*Actor {
	var actorsWithLoanOtherThan []*Actor

	for _, actor := range ListActorsWithLoan(actorType) {
		targetLoanDetected := false

		for _, actorLoan := range actor.Loans {
			if actorLoan.ID == loan.ID {
				targetLoanDetected = true
			}
		}

		if !targetLoanDetected && !actorAlreadyInSlice(*actor, actorsWithLoanOtherThan) {
			fmt.Println("[Actor] ✅ Adding actor to slice:", actor.ID)
			actorsWithLoanOtherThan = append(actorsWithLoanOtherThan, actor)
		}
	}
	return actorsWithLoanOtherThan
}

func newActor(actorType string, name string, balance float64) *Actor {
	return &Actor{
		Name:    name,
		Loans:   []*Loan{},
		Balance: balance,
		Type:    actorType,
	}
}

func CreateActor(actorType string, name string, balance float64) *Actor {
	actor := newActor(actorType, name, balance)
	result := global.Db.Create(&actor)

	if actor.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return actor
}

func newDefaultActor(actorType string) *Actor {
	return &Actor{
		Name:           faker.Name().Name(),
		Loans:          []*Loan{},
		InitialBalance: configs.Actor.InitialBalance,
		Balance:        configs.Actor.InitialBalance,
		Type:           actorType,
	}
}

func CreateDefaultActor(actorType string) *Actor {
	actor := newDefaultActor(actorType)
	actor.Save()
	return actor
}

func actorAlreadyInSlice(newActor Actor, sliceOfActors []*Actor) bool {
	for _, actor := range sliceOfActors {
		if actor.ID == newActor.ID {
			return true
		}
	}

	return false
}
