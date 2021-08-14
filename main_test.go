package main

import (
	"testing"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestMain_prepareSimulation(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	configs.Loan.InitialQuantity = 1
	configs.Loan.InsuredQuantityRatio = 1
	prepareSimulation()

	assert.Len(models.ListLoans(), 1)
	assert.Len(models.ListActiveLoans(), 1)
	assert.Len(models.ListBorrowers(), 1)
	assert.Len(models.ListLenders(), 5)
	assert.Len(models.ListInsurers(), 5)
	assert.Len(models.ListTransactions(), 6)
}
