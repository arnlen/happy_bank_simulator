package factories

import (
	"happy_bank_simulator/models"
)

func NewBorrowers(quantity int) []*models.Actor {
	var borrowers []*models.Actor
	for i := 0; i < quantity; i++ {
		borrower := models.CreateDefaultActor("borrower")
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func NewBorrower() *models.Actor {
	return NewBorrowers(1)[0]
}

func NewBorrowersWithLoan(quantity int) []*models.Actor {
	var borrowers []*models.Actor
	for i := 0; i < quantity; i++ {
		borrower := NewBorrower()
		loan := NewLoan()
		borrower.AssignLoan(loan)
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func NewBorrowerWithLoan() *models.Actor {
	return NewBorrowersWithLoan(1)[0]
}

func NewLenders(quantity int) []*models.Actor {
	var lenders []*models.Actor
	for i := 0; i < quantity; i++ {
		lender := models.CreateDefaultActor("lender")
		lenders = append(lenders, lender)
	}
	return lenders
}

func NewLender() *models.Actor {
	return NewLenders(1)[0]
}

func NewLendersWithLoan(quantity int) []*models.Actor {
	var lenders []*models.Actor
	for i := 0; i < quantity; i++ {
		lender := NewLender()
		loan := NewLoan()
		lender.AssignLoan(loan)
		lenders = append(lenders, lender)
	}
	return lenders
}

func NewLenderWithLoan() *models.Actor {
	return NewLendersWithLoan(1)[0]
}

func NewInsurers(quantity int) []*models.Actor {
	var insurers []*models.Actor
	for i := 0; i < quantity; i++ {
		insurer := models.CreateDefaultActor("insurer")
		insurers = append(insurers, insurer)
	}
	return insurers
}

func NewInsurer() *models.Actor {
	return NewInsurers(1)[0]
}

func NewInsurersWithLoan(quantity int) []*models.Actor {
	var insurers []*models.Actor
	for i := 0; i < quantity; i++ {
		insurer := NewInsurer()
		loan := NewLoan()
		insurer.AssignLoan(loan)
		insurers = append(insurers, insurer)
	}
	return insurers
}

func NewInsurerWithLoan() *models.Actor {
	return NewInsurersWithLoan(1)[0]
}
