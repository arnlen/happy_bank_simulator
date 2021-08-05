package factories_test

import (
	"testing"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/factories"
	"happy_bank_simulator/internal/database"

	"github.com/stretchr/testify/assert"
)

func TestActorFactory_NewBorrowerWithLoan(t *testing.T) {
	database.SetupDB()

	borrower := factories.NewBorrowerWithLoan()

	assert.Equal(t, borrower.Type, "borrower", "Actor Type is borrower")
	assert.Equal(t, len(borrower.Loans), 1, "Borrower has one loan")
}

func TestActorFactory_NewLenderWithLoan(t *testing.T) {
	lender := factories.NewLenderWithLoan()

	assert.Equal(t, lender.Type, "lender")
	assert.Equal(t, len(lender.Loans), 1)
	assert.Equal(t, lender.Loans[0].Amount, configs.Loan.DefaultAmount)
}

func TestActorFactory_NewInsurerWithLoan(t *testing.T) {
	insurer := factories.NewInsurerWithLoan()

	assert.Equal(t, insurer.Type, "insurer")
	assert.Equal(t, len(insurer.Loans), 1)
	assert.Equal(t, insurer.Loans[0].Amount, configs.Loan.DefaultAmount)
}
