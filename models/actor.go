package models

import (
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

func (instance *Actor) GetNetBalance() float64 {
	if instance.IsBorrower() {
		return instance.Balance - instance.GetTotalAmountBorrowed()
	}

	return instance.Balance
}

func (instance *Actor) GetTotalAmountBorrowed() float64 {
	if !instance.IsBorrower() {
		log.Fatal("This actor is not a borrower")
	}

	loans := instance.Loans
	totalAmoutBorrowed := 0.0

	for _, loan := range loans {
		totalAmoutBorrowed += loan.Amount
	}
	return totalAmoutBorrowed
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
	case "insurer":
		loan.AssignInsurer(instance)
	case "lender":
		loan.AssignLender(instance)
	}
}

func (instance *Actor) CanTakeThisLoan(loan Loan) bool {
	// BORROWER
	// => NetBalance > ratio

	// LENDER
	// => NetBalance > loan.amount / qtyOfLenders

	// LENDER
	// => NetBalance > loan.amount / qtyOfInsurers

	return false
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
	global.Db.Preload(clause.Associations).Where("type = ?", actorType).Find(&actors)
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
	actors := ListActors(actorType)
	var availableActorsWithoutLoan []*Actor
	for _, actor := range actors {
		if len(actor.Loans) == 0 {
			availableActorsWithoutLoan = append(availableActorsWithoutLoan, actor)
		}
	}
	return availableActorsWithoutLoan
}

func ListActorsWithLoanOtherThan(actorType string, loan *Loan) []*Actor {
	actors := ListActorsWithoutLoan(actorType)
	var availableActorsWithLoan []*Actor

	for _, actor := range actors {
		if len(actor.Loans) != 0 {
			for _, actorLoan := range actor.Loans {
				if actorLoan.ID != loan.ID && !isActorAlreadyInSlice(*actor, actors) {
					availableActorsWithLoan = append(availableActorsWithLoan, actor)
				}
			}
		}
	}
	return availableActorsWithLoan
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

func isActorAlreadyInSlice(newActor Actor, sliceOfActors []*Actor) bool {
	for _, actor := range sliceOfActors {
		if actor.ID == newActor.ID {
			return true
		}
	}

	return false
}
