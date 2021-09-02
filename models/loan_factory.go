package models

import "happy_bank_simulator/app/configs"

func CreateLoans(quantity int) []*Loan {
	var loans []*Loan
	for i := 0; i < quantity; i++ {
		loan := createDefaultLoan()
		loans = append(loans, loan)
	}
	return loans
}

func CreateLoan() *Loan {
	return CreateLoans(1)[0]
}

func CreateLoanWithBorrowerLendersInsurers() *Loan {
	loan := CreateLoan()
	borrower := CreateBorrower()
	lender := CreateLender()
	insurer := CreateInsurer()

	loan.AssignBorrower(borrower)
	loan.AssignLender(lender)
	loan.AssignInsurer(insurer)

	return loan
}

// ----- PRIVATE METHODS -----

func newDefaultLoan() *Loan {
	return &Loan{
		StartDate:           configs.General.StartDate,
		Duration:            configs.Loan.DefaultDuration,
		Amount:              configs.Loan.DefaultAmount,
		SecurityDepositRate: configs.Loan.SecurityDepositRate,
		InterrestRate:       configs.General.CreditInterestRate,
		InsuranceRate:       configs.General.InsuranceInterestRate,
		IsActive:            false,
	}
}

func createDefaultLoan() *Loan {
	loan := newDefaultLoan()
	loan.Save()
	return loan
}
