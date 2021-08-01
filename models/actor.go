package models

import (
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"
	"log"

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
	Type           string
}

// ------- Instance methods -------

func (instance *Actor) Refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

func (instance *Actor) Save() {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Actor) GetNetBalance() float64 {
	netBalance := 0.0

	switch instance.Type {
	case configs.Actor.Borrower:
		netBalance = instance.Balance - instance.GetTotalAmountBorrowed()
	default:
		netBalance = instance.Balance
	}

	return netBalance
}

// TODO: refactor using double return, including error. See Go time package for example
func (instance *Actor) GetTotalAmountBorrowed() float64 {
	if instance.Type != configs.Actor.Borrower {
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

func (instance *Actor) GetID() uint {
	return instance.ID
}

// ------- Package methods -------

func FindActor(actorType string, id int) *Actor {
	var actor Actor
	database.GetDB().Preload(clause.Associations).First(&actor, id)
	return &actor
}

func ListActors(actorType string) []*Actor {
	var actors []*Actor
	database.GetDB().Preload(clause.Associations).Find(&actors)
	return actors
}

// TODO: Next here: convert this copy/pasted methods from Actor to Actor

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

// Duplicate with Insurer: same method
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

func NewActor(actorType string, name string, balance float64) *Actor {
	return &Actor{
		Name:    name,
		Loans:   []*Loan{},
		Balance: balance,
		Type:    actorType,
	}
}

func NewDefaultActor(actorType string) *Actor {
	return &Actor{
		Name:           faker.Name().Name(),
		Loans:          []*Loan{},
		InitialBalance: configs.Actor.InitialBalance,
		Balance:        configs.Actor.InitialBalance,
		Type:           actorType,
	}
}

func CreateDefaultActor(actorType string) *Actor {
	actor := NewDefaultActor(actorType)
	actor.Save()
	return actor
}

func CreateActor(actorType string, name string, balance float64) *Actor {
	actor := NewActor(actorType, name, balance)
	result := database.GetDB().Create(&actor)

	if actor.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

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
