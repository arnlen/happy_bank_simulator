package factories

import (
	"happy_bank_simulator/models"
)

func NewTransaction() *models.Transaction {
	sender := NewInsurer()
	receiver := NewLender()
	amount := 1000.0

	return models.CreateTransaction(*sender, *receiver, amount)
}
