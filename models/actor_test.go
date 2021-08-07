package models_test

import (
	"testing"

	"happy_bank_simulator/factories"
	"happy_bank_simulator/internal/database"

	"github.com/stretchr/testify/assert"
)

func TestActor_GetNetBalance(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrowerWithLoan := factories.NewBorrowerWithLoan()
	lenderWithLoan := factories.NewLenderWithLoan()
	insurerWithLoan := factories.NewInsurerWithLoan()

	assert.Equal(borrowerWithLoan.GetNetBalance(), borrowerWithLoan.Balance-borrowerWithLoan.Loans[0].Amount)
	assert.Equal(lenderWithLoan.GetNetBalance(), lenderWithLoan.Balance)
	assert.Equal(insurerWithLoan.GetNetBalance(), insurerWithLoan.Balance)
}

func TestActor_GetTotalAmountBorrowed(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrowerWithLoan := factories.NewBorrowerWithLoan()

	assert.Equal(borrowerWithLoan.GetTotalAmountBorrowed(), borrowerWithLoan.Loans[0].Amount)
}

func TestActorFactory_UpdateBalance(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	actor := factories.NewInsurer()
	actor.Balance = 0
	actor.UpdateBalance(1000.0)

	assert.Equal(actor.Balance, 1000.0, "Balance of the actor should be increased be 1000")

	actor.UpdateBalance(-500.0)

	assert.Equal(actor.Balance, 500.0, "Balance of the actor should be decreased be 500")
}

func TestActorFactory_UpdateMontlyIncomes(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := factories.NewBorrower()
	borrower.UpdateMontlyIncomes(1000)

	assert.Equal(borrower.MonthlyIncomes, 1000.0, "Borrower's monthly income should be updated")
}

func TestActorFactory_AssignLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := factories.NewBorrower()
	loan := factories.NewLoan()
	borrower.AssignLoan(loan)

	assert.Equal(loan.Borrower.ID, borrower.ID, "Loan's borrower ID match borrower's ID")
	assert.Equal(borrower.Loans[0].ID, loan.ID, "Borrower's loan ID match loan's ID")

	lender := factories.NewLender()
	loan = factories.NewLoan()
	lender.AssignLoan(loan)

	assert.Equal(len(loan.Lenders), 1, "Loan should have one lender")
	assert.Equal(loan.Lenders[0].ID, lender.ID, "Loan's lender ID should be our loan ID")
	assert.Equal(len(lender.Loans), 1, "Lender should have one loan")
	assert.Equal(lender.Loans[0].ID, loan.ID, "Lender's loan ID match loan's ID")

	insurer := factories.NewInsurer()
	loan = factories.NewLoan()
	insurer.AssignLoan(loan)

	assert.Equal(len(loan.Insurers), 1, "Loan should have one insurer")
	assert.Equal(loan.Insurers[0].ID, insurer.ID, "Loan's insurer ID should be our loan ID")
	assert.Equal(len(insurer.Loans), 1, "Insurer should have one loan")
	assert.Equal(insurer.Loans[0].ID, loan.ID, "Insurer's loan ID match loan's ID")
}

func TestActorFactory_CanTakeThisLoan(t *testing.T) {
	// When actor is a borrower
	// Case "can take"
	// Case "cannot take"

	// When actor is a lender
	// Case "can take"
	// Case "cannot take"

	// When actor is an insurer
	// Case "can take"
	// Case "cannot take"
}
