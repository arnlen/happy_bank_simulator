package models_test

import (
	"testing"
	"time"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestLoan_EndDate(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	loan.StartDate = "01/2022"
	loan.Duration = 3

	assert.Equal(loan.EndDate(), "04/2022")
}

func TestLoan_AssignBorrower(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	borrower := models.CreateBorrower()
	loan.AssignBorrower(borrower)

	assert.Equal(borrower.ID, loan.Borrower.ID)
}

func TestLoan_AssignLenders(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	assert.Len(loan.Lenders, 0)

	lenders := models.CreateLenders(3)
	loan.AssignLenders(lenders)

	assert.Len(loan.Lenders, 3)
}

func TestLoan_AssignInsurers(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	assert.Len(loan.Insurers, 0)

	insurers := models.CreateInsurers(3)
	loan.AssignInsurers(insurers)

	assert.Len(loan.Insurers, 3)
}

func TestLoan_AssignInsurer(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	assert.Len(loan.Insurers, 0)

	insurer := models.CreateInsurer()
	loan.AssignInsurer(insurer)

	assert.Len(loan.Insurers, 1)
}

func TestLoan_Activate(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoanWithBorrowerLendersInsurers()
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

	loan := models.CreateLoan()
	assert.Equal(0.0, loan.RefundedAmount)
	loan.Refund(1200)
	loan.Refresh()

	assert.Equal(1200.0, loan.RefundedAmount)
}

func TestLoan_SetRandomNumberOfMonthsBeforeFailure(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	loan.NumberOfMonthsBeforeFailure = 0
	loan.Save()
	loan.SetRandomNumberOfMonthsBeforeFailure()

	assert.Less(0, loan.NumberOfMonthsBeforeFailure)
}

func TestLoan_WillFail(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loanWontFail := models.CreateLoan()
	loanWillFail := models.CreateLoan()
	loanWillFail.SetRandomNumberOfMonthsBeforeFailure()

	assert.False(loanWontFail.WillFail())
	assert.True(loanWillFail.WillFail())
}

func TestLoan_WillFailOn(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loanWillFail := models.CreateLoan()
	loanWillFail.SetRandomNumberOfMonthsBeforeFailure()

	assert.IsType(time.Time{}, loanWillFail.WillFailOnTime())
	assert.IsType(*new(string), loanWillFail.WillFailOnString())
}

func TestLoan_SetBorrowerMonthlyIncomes(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoanWithBorrowerLendersInsurers()
	assert.Equal(0.0, loan.Borrower.MonthlyIncomes)
	loan.SetBorrowerMonthlyIncomes()

	assert.Less(0.0, loan.Borrower.MonthlyIncomes)
}

func TestLoan_SetupLenders(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	// Case 1: when no lender in database
	loan1 := models.CreateLoan()
	assert.Len(loan1.Lenders, 0)
	assert.Len(models.ListActors(configs.Actor.LenderString), 0)

	loan1.SetupLenders()
	assert.Len(loan1.Lenders, 5)
	assert.Len(models.ListActors(configs.Actor.LenderString), 5)

	// Case 2: when already 5 lenders in database
	loan2 := models.CreateLoan()
	loan2.SetupLenders()
	assert.Len(loan2.Lenders, 5)
	assert.Len(models.ListActors(configs.Actor.LenderString), 5)
}

func TestLoan_SetupInsurers(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	// Case 1: when no lender in database
	loan1 := models.CreateLoan()
	assert.Len(loan1.Insurers, 0)
	assert.Len(models.ListActors(configs.Actor.InsurerString), 0)

	loan1.SetupInsurers()
	assert.Len(loan1.Insurers, 5)
	assert.Len(models.ListActors(configs.Actor.InsurerString), 5)

	// Case 2: when already 5 lenders in database
	loan2 := models.CreateLoan()
	loan2.SetupInsurers()
	assert.Len(loan2.Insurers, 5)
	assert.Len(models.ListActors(configs.Actor.InsurerString), 5)
}

func TestLoan_RequiredMontlyIncomes(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	// Case loan will fail
	loanWillFail := models.CreateLoanWithBorrowerLendersInsurers()
	loanWillFail.SetRandomNumberOfMonthsBeforeFailure()
	monthlyPaymentDue := loanWillFail.MonthlyPayment()
	assert.Less(loanWillFail.RequiredMontlyIncomes(), monthlyPaymentDue)

	// Case loan wont fail
	loanWontFail := models.CreateLoanWithBorrowerLendersInsurers()
	monthlyPaymentDue = loanWontFail.MonthlyPayment()
	assert.Equal(loanWontFail.RequiredMontlyIncomes(), monthlyPaymentDue)
}

func TestLoan_ListLoans(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	models.CreateLoans(3)
	assert.Len(models.ListLoans(), 3)
}

func TestLoan_ListActiveLoans(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loans := models.CreateLoans(3)
	loans[0].Activate()
	loans[2].Activate()

	assert.Len(models.ListActiveLoans(), 2)
}
