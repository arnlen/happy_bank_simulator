package factories

import (
	"happy_bank_simulator/models"
)

func NewBorrower() *models.Actor {
	return models.CreateDefaultActor("borrower")
}

func NewBorrowerWithLoan() *models.Actor {
	borrower := NewBorrower()
	loan := NewLoan()
	borrower.AssignLoan(loan)
	return borrower
}

func NewLender() *models.Actor {
	return models.CreateDefaultActor("lender")
}

func NewLenderWithLoan() *models.Actor {
	lender := NewLender()
	loan := NewLoan()
	lender.AssignLoan(loan)
	return lender
}

func NewInsurer() *models.Actor {
	return models.CreateDefaultActor("insurer")
}

func NewInsurerWithLoan() *models.Actor {
	insurer := NewInsurer()
	loan := NewLoan()
	insurer.AssignLoan(loan)
	return insurer
}
