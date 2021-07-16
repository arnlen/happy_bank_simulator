package models

import (
	"log"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"syreclabs.com/go/faker"
)

type Borrower struct {
	gorm.Model
	Name       string
	Loans      []Loan
	Balance    int
	WillFailOn string
}

func (instance *Borrower) ModelName() string {
	return "emprunteur"
}

func (instance *Borrower) Refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

func (instance *Borrower) Save() {
	db := database.GetDB()
	result := db.Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

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

func NewBorrower(name string, balance int) *Borrower {
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

func CreateBorrower(name string, balance int) *Borrower {
	borrower := NewBorrower(name, balance)
	result := database.GetDB().Create(&borrower)

	if borrower.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return borrower
}
