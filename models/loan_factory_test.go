package models_test

import (
	"testing"

	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestActorFactory_CreateLoans(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loans := models.CreateLoans(3)

	assert.IsType([]*models.Loan{}, loans)
	assert.Len(loans, 3)
}

func TestActorFactory_CreateLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()

	assert.IsType(models.Loan{}, *loan)
}

func TestActorFactory_CreateLoanWithBorrowerLendersInsurers(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoanWithBorrowerLendersInsurers()
	borrower := loan.Borrower
	lenders := loan.Lenders
	insurers := loan.Insurers

	assert.IsType(models.Loan{}, *loan)
	assert.IsType(models.Actor{}, borrower)
	assert.IsType([]*models.Actor{}, lenders)
	assert.IsType([]*models.Actor{}, insurers)
	assert.Len(lenders, 1)
	assert.Len(insurers, 1)
	assert.NotEqual(0, loan.ID)
	assert.NotEqual(0, borrower.ID)
	assert.NotEqual(0, lenders[0].ID)
	assert.NotEqual(0, insurers[0].ID)
}
