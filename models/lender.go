package models

import (
	"happy_bank_simulator/services/database"
	"log"

	"gorm.io/gorm"
)

type Lender struct {
	gorm.Model
	Name    string
	Loans   []Loan
	Balance float64
}

func (instance *Lender) Save() *Lender {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func NewLender(name string, balance float64) *Lender {
	return &Lender{
		Name:    name,
		Loans:   []Loan{},
		Balance: balance,
	}
}

func CreateLender(name string, balance float64) *Lender {
	lender := NewLender(name, balance)
	result := database.GetDB().Create(&lender)

	if lender.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return lender
}
