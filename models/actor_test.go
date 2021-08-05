package models_test

import (
	"testing"

	"happy_bank_simulator/factories"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestActor_GetNetBalance(t *testing.T) {
	database.SetupDB()

	assert := assert.New(t)

	borrower := factories.NewBorrower()
	borrowerWithLoan := factories.NewBorrowerWithLoan()
	lender := factories.NewLender()
	lenderWithLoan := factories.NewLenderWithLoan()
	insurer := factories.NewInsurer()
	insurerWithLoan := factories.NewInsurerWithLoan()

	var tests = []struct {
		name     string
		input    *models.Actor
		expected float64
	}{
		{"Borrower", borrower, borrower.Balance},
		{"Borrower with loan", borrowerWithLoan, borrowerWithLoan.Balance},
		{"Lender", lender, lender.Balance},
		{"Lender with loan", lenderWithLoan, lenderWithLoan.Balance},
		{"Insurer", insurer, insurer.Balance},
		{"Insurer with loan", insurerWithLoan, insurerWithLoan.Balance},
	}

	for _, test := range tests {
		assert.Equal(test.input.GetNetBalance(), test.expected, test.name)
	}
}

func TestActorFactory_AssignLoan(t *testing.T) {
	database.SetupDB()

	borrower := factories.NewBorrower()
	loan := factories.NewLoan()
	borrower.AssignLoan(loan)

	assert.Equal(t, loan.Borrower.ID, borrower.ID, "Loan's borrower ID match borrower's ID")
	assert.Equal(t, borrower.Loans[0].ID, loan.ID, "Borrower's loan ID match loan's ID")
}
