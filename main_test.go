package main

import (
	"math/rand"
	"testing"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/charts"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestMain_PrepareSimulation(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	configs.Loan.InitialQuantity = 1
	configs.Loan.InsuredQuantityRatio = 1
	PrepareSimulation()

	assert.Len(models.ListLoans(), 1)
	assert.Len(models.ListActiveLoans(), 1)
	assert.Len(models.ListBorrowers(), 1)
	assert.Len(models.ListLenders(), 5)
	assert.Len(models.ListInsurers(), 5)
	assert.Len(models.ListTransactions(), 6)
}

func TestMain_RunMonthLoop(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	configs.Loan.InitialQuantity = 1
	configs.Loan.InsuredQuantityRatio = 1
	PrepareSimulation()

	loan := models.ListActiveLoans()[0]
	borrower := loan.Borrower
	lenders := loan.Lenders
	lender := lenders[rand.Intn(len(lenders)-1)]
	insurers := loan.Insurers
	insurer := insurers[rand.Intn(len(insurers)-1)]
	chartsManager := charts.ChartsManager{}

	RunMonthLoop(0, &chartsManager)

	borrowerBeforeLoopBalance := borrower.Balance
	lenderBeforeLoopBalance := lender.Balance
	insurerBeforeLoopBalance := insurer.Balance

	borrower.Refresh()
	lender.Refresh()
	insurer.Refresh()

	assert.Less(borrower.Balance, borrowerBeforeLoopBalance)
	assert.Greater(lender.Balance, lenderBeforeLoopBalance)
	assert.Greater(insurer.Balance, insurerBeforeLoopBalance)
}
