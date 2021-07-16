package configs

import ()

// Types definitions

type general struct {
	StartDate             string
	Duration              int
	CreditInterestRate    float64
	InsuranceInterestRate float64
}

type loan struct {
	DefaultAmount       int
	DefaultDuration     int
	SecurityDepositRate float64
	InitialQuantity     int
	InsuredQuantityRate float64
}

type borrower struct {
	InitialBalance int
	FailureRate    float64
}

type lender struct {
	InitialBalance   int
	MaxAmountPerLoan int
}

type insurer struct {
	InitialBalance   int
	MaxAmountPerLoan int
}

// Config intialization

var General = general{
	StartDate:             "1/01/21",
	Duration:              60,
	CreditInterestRate:    0.03,
	InsuranceInterestRate: 0.03,
}

var Loan = loan{
	DefaultAmount:       5000,
	DefaultDuration:     12,
	SecurityDepositRate: 0.1,
	InitialQuantity:     1,
	InsuredQuantityRate: 0.8,
}

var Borrower = borrower{
	InitialBalance: 5000,
	FailureRate:    0.2,
}

var Lender = lender{
	InitialBalance:   5000,
	MaxAmountPerLoan: 1000,
}

var Insurer = insurer{
	InitialBalance:   5000,
	MaxAmountPerLoan: 1000,
}
