package models_test

import (
	"testing"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestActorFactory_CreateBorrowers(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrowers := models.CreateBorrowers(3)

	assert.Len(borrowers, 3)
	assert.True(borrowers[0].IsBorrower())
	assert.Len(borrowers[0].Loans, 0)
}

func TestActorFactory_CreateBorrowerWithLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := models.CreateBorrowerWithLoan()

	assert.True(borrower.IsBorrower())
	assert.Len(borrower.Loans, 1)
	assert.NotEqual(borrower.Loans[0].ID, 0)
}

func TestActorFactory_CreateLender(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	lender := models.CreateLender()

	assert.True(lender.IsLender())
	assert.Len(lender.Loans, 0)
}

func TestActorFactory_CreateLenderWithLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	lender := models.CreateLenderWithLoan()

	assert.True(lender.IsLender())
	assert.Len(lender.Loans, 1)
	assert.NotEqual(lender.Loans[0].ID, 0)
	assert.Equal(lender.Loans[0].Amount, configs.Loan.DefaultAmount)
}

func TestActorFactory_CreateInsurer(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	insurer := models.CreateInsurer()

	assert.True(insurer.IsInsurer())
	assert.Len(insurer.Loans, 0)
}

func TestActorFactory_CreateInsurerWithLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	insurer := models.CreateInsurerWithLoan()

	assert.True(insurer.IsInsurer())
	assert.Len(insurer.Loans, 1)
	assert.NotEqual(insurer.Loans[0].ID, 0)
	assert.Equal(insurer.Loans[0].Amount, configs.Loan.DefaultAmount)
}
