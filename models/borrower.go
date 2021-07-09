package models

import (
	"happy_bank_simulator/database"
	"log"

	"gorm.io/gorm"
)

type Borrower struct {
	gorm.Model
	Name    string
	Loans   []Loan
	Balance float64
}

func (instance *Borrower) Save() *Borrower {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func NewBorrower(name string, balance float64) *Borrower {
	return &Borrower{
		Name:    name,
		Loans:   []Loan{},
		Balance: balance,
	}
}

func CreateBorrower(name string, balance float64) *Borrower {
	borrower := NewBorrower(name, balance)
	result := database.GetDB().Create(&borrower)

	if borrower.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return borrower
}
