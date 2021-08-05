package factories

import (
	"happy_bank_simulator/models"
)

func NewLoan() *models.Loan {
	return models.NewDefaultLoan()
}
