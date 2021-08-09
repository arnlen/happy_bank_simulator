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
	amount := configs.Loan.DefaultAmount
	startDate := configs.General.StartDate
	duration := configs.Loan.DefaultDuration
	creditRate := configs.General.CreditInterestRate
	insuranceRate := configs.General.InsuranceInterestRate

	initialDeposit := configs.Loan.SecurityDepositRate * float64(amount)
	monthlyCredit := calculateMonthlyCreditPayment(creditRate, duration, amount)
	monthlyInsurance := calculateMonthlyInsurancePayment(insuranceRate, amount)

	return &Loan{
		StartDate:        startDate,
		Duration:         configs.Loan.DefaultDuration,
		Amount:           configs.Loan.DefaultAmount,
		InitialDeposit:   initialDeposit,
		CreditRate:       creditRate,
		InsuranceRate:    insuranceRate,
		MonthlyCredit:    monthlyCredit,
		MonthlyInsurance: monthlyInsurance,
		IsActive:         false,
	}
}

func createDefaultLoan() *Loan {
	loan := newDefaultLoan()
	loan.Save()
	return loan
}
