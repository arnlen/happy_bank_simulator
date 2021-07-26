package models

import (
	"log"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"syreclabs.com/go/faker"
)

// Declare conformity with Actor interface
var _ ModelBase = (*Borrower)(nil)
var _ Actor = (*Borrower)(nil)

type Borrower struct {
	gorm.Model
	Name    string
	Loans   []Loan
	Balance float64
}

// ------- Instance methods -------

func (instance *Borrower) ModelName() string {
	return "borrower"
}

func (instance *Borrower) Refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

func (instance *Borrower) Save() {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Borrower) GetNetBalance() float64 {
	netBalance := instance.Balance - instance.GetTotalAmountBorrowed()
	return netBalance
}

func (instance *Borrower) GetTotalAmountBorrowed() float64 {
	loans := instance.Loans
	totalAmoutBorrowed := 0.0

	for _, loan := range loans {
		totalAmoutBorrowed += loan.Amount
	}
	return totalAmoutBorrowed
}

func (instance *Borrower) UpdateBalance(amount float64) {
	instance.Balance += amount
	instance.Save()
}

func (instance *Borrower) GetID() uint {
	return instance.ID
}

// ------- Package methods -------

func FindBorrower(id int) *Borrower {
	var borrower Borrower
	database.GetDB().Preload(clause.Associations).First(&borrower, id)
	return &borrower
}

func ListBorrowers() []Borrower {
	var borrowers []Borrower
	database.GetDB().Preload(clause.Associations).Find(&borrowers)
	return borrowers
}

func NewBorrower(name string, balance float64) *Borrower {
	return &Borrower{
		Name:    name,
		Loans:   []Loan{},
		Balance: balance,
	}
}

func NewDefaultBorrower() *Borrower {
	return &Borrower{
		Name:    faker.Name().Name(),
		Loans:   []Loan{},
		Balance: configs.Borrower.InitialBalance,
	}
}

func CreateDefaultBorrower() *Borrower {
	borrower := NewDefaultBorrower()
	borrower.Save()
	return borrower
}

func CreateBorrower(name string, balance float64) *Borrower {
	borrower := NewBorrower(name, balance)
	result := database.GetDB().Create(&borrower)

	if borrower.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return borrower
}
