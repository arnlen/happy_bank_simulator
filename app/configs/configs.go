package configs

import ()

// ----- GENERAL CONFIGS -----

type general struct {
	StartDate             string  // Simulation start date
	Duration              int     // Simulation duration (in months)
	CreditInterestRate    float64 // Interest rate of the credit part of a loan
	InsuranceInterestRate float64 // Interest rate of the insurance part of a loan
}

var General = general{
	StartDate:             "07/2022",
	Duration:              60,
	CreditInterestRate:    0.03,
	InsuranceInterestRate: 0.03,
}

// ----- LOAN CONFIGS -----

type loan struct {
	InitialQuantity      int     // How many loans should exist at the beginning at the simulation
	DefaultAmount        float64 // Default amount for a new loan
	DefaultDuration      int     // Default duration for a new loan
	SecurityDepositRate  float64 // For a given loan amout, how much % a borrower must stake
	InsuredQuantityRatio float64 // How many loans are insured, in % of the total
	FailureRate          float64 // How many loans should fail, in % of the total
	String               string
}

var Loan = loan{
	InitialQuantity:      3,
	DefaultAmount:        5000,
	DefaultDuration:      24,
	SecurityDepositRate:  0.1,
	InsuredQuantityRatio: 1,
	FailureRate:          1,
	String:               "loan",
}

// ----- ACTOR CONFIGS -----

type actor struct {
	MaxAmountPerLoan             float64 // Maximum amout of money a lender of an insurer can lend/insure per loan
	InitialBalance               float64 // Initial balance
	BorrowerBalanceLeverageRatio float64 // Ratio between the balance of the borrower and the amount he can borrow
	BorrowerString               string
	LenderString                 string
	InsurerString                string
}

var Actor = actor{
	MaxAmountPerLoan:             1000,
	InitialBalance:               5000,
	BorrowerBalanceLeverageRatio: 1.0,
	BorrowerString:               "borrower",
	LenderString:                 "lender",
	InsurerString:                "insurer",
}

// ----- LOAN CONFIGS -----

type transaction struct {
	String string
}

var Transaction = transaction{
	String: "transaction",
}
