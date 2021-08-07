package factories_test

import (
	"testing"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/factories"
	"happy_bank_simulator/internal/database"

	"github.com/stretchr/testify/assert"
)

func TestActorFactory_NewBorrower(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := factories.NewBorrower()

	assert.True(borrower.IsBorrower())
	assert.Equal(len(borrower.Loans), 0)
}

func TestActorFactory_NewBorrowerWithLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := factories.NewBorrowerWithLoan()

	assert.True(borrower.IsBorrower())
	assert.Equal(len(borrower.Loans), 1)
	assert.NotEqual(borrower.Loans[0].ID, 0)
}

func TestActorFactory_NewLender(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	lender := factories.NewLender()

	assert.True(lender.IsLender())
	assert.Equal(len(lender.Loans), 0)
}

func TestActorFactory_NewLenderWithLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	lender := factories.NewLenderWithLoan()

	assert.True(lender.IsLender())
	assert.Equal(len(lender.Loans), 1)
	assert.NotEqual(lender.Loans[0].ID, 0)
	assert.Equal(lender.Loans[0].Amount, configs.Loan.DefaultAmount)
}

func TestActorFactory_NewInsurer(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	insurer := factories.NewInsurer()

	assert.True(insurer.IsInsurer())
	assert.Equal(len(insurer.Loans), 0)
}

func TestActorFactory_NewInsurerWithLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	insurer := factories.NewInsurerWithLoan()

	assert.True(insurer.IsInsurer())
	assert.Equal(len(insurer.Loans), 1)
	assert.NotEqual(insurer.Loans[0].ID, 0)
	assert.Equal(insurer.Loans[0].Amount, configs.Loan.DefaultAmount)
}

func TestActorFactory_NewLoanWithBorrowerLendersInsurers(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoanWithBorrowerLendersInsurers()

	assert.NotNil(loan.Borrower)
	assert.Equal(len(loan.Lenders), 1)
	assert.Equal(len(loan.Insurers), 1)
	assert.NotEqual(loan.Lenders[0].ID, 0)
	assert.NotEqual(loan.Insurers[0].ID, 0)
}

func TestActorFactory_NewTransaction(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	transaction := factories.NewTransaction()

	assert.NotEqual(transaction.ID, 0)
	assert.NotEqual(transaction.SenderID, 0)
	assert.NotEqual(transaction.ReceiverID, 0)
	assert.Greater(transaction.Amount, 0.0)
}
