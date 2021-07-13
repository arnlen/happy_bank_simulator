package models

import (
	"happy_bank_simulator/services/database"
	"log"

	"gorm.io/gorm"
)

type Insurer struct {
	gorm.Model
	Name    string
	Loans   []Loan
	Balance float64
}

func (instance *Insurer) Save() *Insurer {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func NewInsurer(name string, balance float64) *Insurer {
	return &Insurer{
		Name:    name,
		Loans:   []Loan{},
		Balance: balance,
	}
}

func CreateInsurer(name string, balance float64) *Insurer {
	insurer := NewInsurer(name, balance)
	result := database.GetDB().Create(&insurer)

	if insurer.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return insurer
}
