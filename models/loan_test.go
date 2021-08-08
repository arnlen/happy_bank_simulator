package models_test

import (
	"testing"
	"time"

	"happy_bank_simulator/factories"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

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

func TestLoan_AssignBorrower(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()
	borrower := factories.NewBorrower()
	loan.AssignBorrower(borrower)

	assert.Equal(borrower.ID, loan.Borrower.ID)
}

func TestLoan_AssignLender(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()
	assert.Len(loan.Lenders, 0)

	lenders := factories.NewLenders(3)
	loan.AssignLender(lenders[0])
	loan.AssignLender(lenders[1])
	loan.AssignLender(lenders[2])

	assert.Len(loan.Lenders, 3)
}

func TestLoan_AssignInsurer(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()
	assert.Len(loan.Insurers, 0)

	insurers := factories.NewInsurers(3)
	loan.AssignInsurer(insurers[0])
	loan.AssignInsurer(insurers[1])
	loan.AssignInsurer(insurers[2])

	assert.Len(loan.Insurers, 3)
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

func TestLoan_Refund(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()
	assert.Equal(0.0, loan.RefundedAmount)
	loan.Refund(1200)
	loan.Refresh()

	assert.Equal(1200.0, loan.RefundedAmount)
}

func TestLoan_SetRandomNumberOfMonthsBeforeFailure(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()
	loan.NumberOfMonthsBeforeFailure = 0
	loan.Save()
	loan.SetRandomNumberOfMonthsBeforeFailure()

	assert.Less(0, loan.NumberOfMonthsBeforeFailure)
}

func TestLoan_WillFail(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loanWontFail := factories.NewLoan()
	loanWillFail := factories.NewLoan()
	loanWillFail.SetRandomNumberOfMonthsBeforeFailure()

	assert.False(loanWontFail.WillFail())
	assert.True(loanWillFail.WillFail())
}

func TestLoan_WillFailOn(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loanWillFail := factories.NewLoan()
	loanWillFail.SetRandomNumberOfMonthsBeforeFailure()

	assert.IsType(time.Time{}, loanWillFail.WillFailOnTime())
	assert.IsType(*new(string), loanWillFail.WillFailOnString())
}

func TestLoan_SetBorrowerMonthlyIncomes(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoanWithBorrowerLendersInsurers()
	assert.Equal(0.0, loan.Borrower.MonthlyIncomes)
	loan.SetBorrowerMonthlyIncomes()

	assert.Less(0.0, loan.Borrower.MonthlyIncomes)
}

func TestLoan_RequiredMontlyIncomes(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	// Case loan will fail
	loanWillFail := factories.NewLoanWithBorrowerLendersInsurers()
	loanWillFail.SetRandomNumberOfMonthsBeforeFailure()
	monthlyPaymentDue := loanWillFail.MonthlyPayment()
	assert.Less(loanWillFail.RequiredMontlyIncomes(), monthlyPaymentDue)

	// Case loan wont fail
	loanWontFail := factories.NewLoanWithBorrowerLendersInsurers()
	monthlyPaymentDue = loanWontFail.MonthlyPayment()
	assert.Equal(loanWontFail.RequiredMontlyIncomes(), monthlyPaymentDue)
}

func TestLoan_ListLoans(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	factories.NewLoans(3)
	assert.Len(models.ListLoans(), 3)
}

func TestLoan_ListActiveLoans(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loans := factories.NewLoans(3)
	loans[0].Activate()
	loans[2].Activate()

	assert.Len(models.ListActiveLoans(), 2)
}

func TestActorFactory_CreateDefaultLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateDefaultLoan()

	assert.IsType(models.Loan{}, *loan)
}
