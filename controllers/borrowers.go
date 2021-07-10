package controllers

import (
	"happy_bank_simulator/models"
	"log"
)

// Declare conformity with BaseController interface
var _ BaseController = (*Borrowers)(nil)

type Borrowers struct{}

func (c *Borrowers) Create(name string, balance float64) *models.Borrower {
	borrower := &models.Borrower{
		Name:    name,
		Loans:   []models.Loan{},
		Balance: balance,
	}

	result := borrower.Create()

	if borrower.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return borrower
}
