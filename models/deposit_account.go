package models

import ()

type DepositAccount struct {
	Balance int
}

// ------- Instance methods -------

func (instance *DepositAccount) UpdateBalance(amount int) {
	instance.Balance += amount
}
