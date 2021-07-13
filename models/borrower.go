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

func (instance *Borrower) ModelName() string {
	return "emprunteur"
}

func (instance *Borrower) Save() *Borrower {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func (instance *Borrower) Create() *gorm.DB {
	return database.GetDB().Create(instance)
}
