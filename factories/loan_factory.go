package factories

import (
	"happy_bank_simulator/models"
)

func NewLoans(quantity int) []*models.Loan {
	var loans []*models.Loan
	for i := 0; i < quantity; i++ {
		loan := models.CreateDefaultLoan()
		loans = append(loans, loan)
	}
	return loans
}

func NewLoan() *models.Loan {
	return NewLoans(1)[0]
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
