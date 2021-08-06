package factories

import (
	"happy_bank_simulator/models"
)

func NewLoan() *models.Loan {
	return models.CreateDefaultLoan()
}

func NewLoanWithBorrowerLendersInsurers() *models.Loan {
	loan := NewLoan()
	borrower := NewBorrower()
	lender := NewLender()
	insurer := NewInsurer()

	loan.AssignBorrower(borrower)
	loan.AssignLender(lender)
	loan.AssignInsurer(insurer)

	return loan
}
