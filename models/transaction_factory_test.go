package models_test

import (
	"testing"

	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_CreateTransaction(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	sender := models.CreateInsurer()
	receiver := models.CreateLender()
	sender.Balance = 1000.0
	receiver.Balance = 1000.0
	transactionAmount := 1000.0

	transaction := models.CreateTransaction(*sender, *receiver, transactionAmount)
	sender.Refresh()
	receiver.Refresh()

	assert.Equal(sender.Balance, 0.0)
	assert.Equal(receiver.Balance, 2000.0)
	assert.Equal(transaction.Amount, transactionAmount)
	assert.Equal(transaction.SenderID, int(sender.ID))
	assert.Equal(transaction.ReceiverID, int(receiver.ID))
}

func TestTransaction_CreateDepositTransaction(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := models.CreateBorrower()
	borrower.Balance = 1000.0
	initialDeposit := 500.0
	transaction := models.CreateDepositTransaction(*borrower, initialDeposit)
	borrower.Refresh()

	assert.Equal(borrower.Balance, 500.0)
	assert.Equal(transaction.Amount, initialDeposit)
	assert.Equal(transaction.SenderID, int(borrower.ID))
	assert.Equal(transaction.ReceiverID, 0)
}

func TestTransaction_CreateIncomeTransaction(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := models.CreateBorrower()
	borrower.Balance = 1000.0
	income := 500.00

	transaction := models.CreateIncomeTransaction(*borrower, income)
	borrower.Refresh()

	assert.Equal(borrower.Balance, 1500.0)
	assert.Equal(transaction.Amount, income)
	assert.Equal(transaction.SenderID, 0)
	assert.Equal(transaction.ReceiverID, int(borrower.ID))
}
