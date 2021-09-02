package models_test

import (
	"testing"
	"time"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestLoan_GetEndDate(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	loan.StartDate = "01/2022"
	loan.Duration = 3

	assert.Equal(loan.GetEndDate(), "04/2022")
}

func TestLoan_GetTotalMonthlyPayment(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	loan.MonthlyAmountToRefund = 100
	loan.MonthlyInsuranceCost = 200

	assert.Equal(300.0, loan.GetTotalMonthlyPayment())
}

func TestLoan_GetAmountLentPerLender(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	loan.Amount = 1000
	loan.Save()
	lenders := models.CreateLenders(4)
	loan.AssignLenders(lenders)

	assert.Equal(250.0, loan.GetAmountLentPerLender())
}

func TestLoan_GetAmountInsuredPerInsurer(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	loan.Amount = 1000
	loan.Save()
	insurers := models.CreateInsurers(4)
	loan.AssignInsurers(insurers)

	assert.Equal(250.0, loan.GetAmountInsuredPerInsurer())
}

func TestLoan_SetRandomNumberOfMonthsBeforeFailure(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	loan.NumberOfMonthsBeforeFailure = 0
	loan.Save()
	loan.SetRandomNumberOfMonthsBeforeFailure()

	assert.Greater(loan.NumberOfMonthsBeforeFailure, 0)
}

func TestLoan_SetBorrowerMonthlyIncomes(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoanWithBorrowerLendersInsurers()
	assert.Equal(0.0, loan.Borrower.MonthlyIncomes)
	loan.SetBorrowerMonthlyIncomes()

	assert.Greater(loan.Borrower.MonthlyIncomes, 0.0)
}

func TestLoan_AddToRefund(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	assert.Equal(0.0, loan.RefundedAmount)
	loan.AddToRefund(1200)
	loan.Refresh()

	assert.Equal(1200.0, loan.RefundedAmount)
}

func TestLoan_AssignBorrower(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	borrower := models.CreateBorrower()
	loan.AssignBorrower(borrower)

	assert.Equal(borrower.ID, loan.Borrower.ID)
}

func TestLoan_AssignLender(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	assert.Len(loan.Lenders, 0)

	lender := models.CreateLender()
	loan.AssignLender(lender)

	assert.Len(loan.Lenders, 1)
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

func TestLoan_AssignInsurer(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	assert.Len(loan.Insurers, 0)

	insurer := models.CreateInsurer()
	loan.AssignInsurer(insurer)

	assert.Len(loan.Insurers, 1)
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

func TestLoan_Deactivate(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	loan.Activate()

	assert.Equal(loan.IsActive, true)

	loan.Deactivate()

	assert.Equal(loan.IsActive, false)
}

func TestLoan_WillFailOnTimeAndString(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loanWillFail := models.CreateLoan()
	loanWillFail.SetRandomNumberOfMonthsBeforeFailure()

	assert.IsType(time.Time{}, loanWillFail.WillFailOnTime())
	assert.IsType(*new(string), loanWillFail.WillFailOnString())
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

func TestLoan_CreateMontlyTransactionsFromBorrowerToLenders(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoanWithBorrowerLendersInsurers()
	lenders := models.CreateLenders(4)
	loan.AssignLenders(lenders)
	borrower := loan.Borrower
	initialBorrowerBalance := borrower.Balance

	assert.Len(models.ListTransactions(), 0)

	loan.CreateMontlyTransactionsFromBorrowerToLenders()
	borrower.Refresh()

	assert.Len(models.ListTransactions(), 5)
	assert.Less(borrower.Balance, initialBorrowerBalance)
}

func TestLoan_CreateMontlyTransactionsFromLendersToInsurers(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	// When loan isn't insured
	unisuredLoan := models.CreateLoan()
	unisuredLoan.IsInsured = false

	unisuredLoan.CreateMontlyTransactionsFromLendersToInsurers()

	assert.Len(models.ListTransactions(), 0)

	// When loan is insured
	insuredLoan := models.CreateLoan()
	insurers := models.CreateInsurers(5)
	insuredLoan.AssignInsurers(insurers)
	insuredLoan.IsInsured = true

	insuredLoan.CreateMontlyTransactionsFromLendersToInsurers()

	assert.Len(models.ListTransactions(), 5)
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

func TestLoan_ShouldFailThisMonth(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()
	currentDate := loan.WillFailOnTime()

	assert.True(loan.ShouldFailThisMonth(currentDate))
}
