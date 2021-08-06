package models_test

import (
	"testing"

	"happy_bank_simulator/factories"
	"happy_bank_simulator/internal/database"

	"github.com/stretchr/testify/assert"
)

func TestLoan_EndDate(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()
	loan.StartDate = "01/2022"
	loan.Duration = 3

	assert.Equal(loan.EndDate(), "04/2022")
}

func TestLoan_AddLender(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()
	assert.Equal(len(loan.Lenders), 0)

	lender := factories.NewLender()
	loan.AddLender(lender)

	assert.Equal(len(loan.Lenders), 1)
}

func TestLoan_AddInsurer(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()
	assert.Equal(len(loan.Insurers), 0)

	insurer := factories.NewInsurer()
	loan.AddInsurer(insurer)

	assert.Equal(len(loan.Insurers), 1)
}

func TestLoan_Activate(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoanWithBorrowerLendersInsurers()
	lender := loan.Lenders[0]
	insurer := loan.Insurers[0]

	assert.Equal(loan.IsActive, false)
	assert.Equal(loan.Borrower.Balance, loan.Borrower.InitialBalance)
	assert.Equal(lender.Balance, lender.InitialBalance)
	assert.Equal(insurer.Balance, insurer.InitialBalance)

	loan.Activate()

	assert.Equal(loan.IsActive, true)
	assert.Greater(loan.Borrower.Balance, loan.Borrower.InitialBalance)
	assert.Less(lender.Balance, lender.InitialBalance)
	assert.Equal(insurer.Balance, insurer.InitialBalance)
}
