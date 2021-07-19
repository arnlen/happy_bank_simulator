package configs

import ()

// Types definitions

type general struct {
	StartDate             string  // Simulation start date
	Duration              int     // Simulation duration (in months)
	CreditInterestRate    float64 // Interest rate of the credit part of a loan
	InsuranceInterestRate float64 // Interest rate of the insurance part of a loan
}

type loan struct {
	InitialQuantity      int     // How many loans should exist at the beginning at the simulation
	DefaultAmount        int     // Default amount for a new loan
	DefaultDuration      int     // Default duration for a new loan
	SecurityDepositRate  float64 // For a given loan amout, how much % a borrower must stake
	InsuredQuantityRatio float64 // How many loans are insured, in % of the total
	FailureRate          float64 // How many loans should fail, in % of the total
}

type borrower struct {
	InitialBalance       int     // Initial balance
	BalanceLeverageRatio float64 // Ratio between the balance of the borrower and the amount he can borrow
}

type lender struct {
	InitialBalance   int // Initial balance
	MaxAmountPerLoan int // Maximum amout of money a lender can lend per loan
}

type insurer struct {
	InitialBalance   int // Initial balance
	MaxAmountPerLoan int // Maximum amout of money an insurer can insure per loan
}

// Config intialization

var General = general{
	StartDate:             "07/2022",
	Duration:              60,
	CreditInterestRate:    0.03,
	InsuranceInterestRate: 0.03,
}

var Loan = loan{
	InitialQuantity:      1,
	DefaultAmount:        5000,
	DefaultDuration:      25,
	SecurityDepositRate:  0.1,
	InsuredQuantityRatio: 0.8,
	FailureRate:          1,
}

var Borrower = borrower{
	InitialBalance:       5000,
	BalanceLeverageRatio: 1.0,
}

var Lender = lender{
	InitialBalance:   5000,
	MaxAmountPerLoan: 1000,
}

var Insurer = insurer{
	InitialBalance:   5000,
	MaxAmountPerLoan: 1000,
}
