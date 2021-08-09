package models_test

import (
	"testing"

	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_ListTransactions(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	models.CreateDefaultTransaction()
	models.CreateDefaultTransaction()

	transactions := models.ListTransactions()

	assert.Len(transactions, 2)
	assert.IsType([]*models.Transaction{}, transactions)
}
