package services

import (
	"github.com/drum445/gofin"
)

func CalculateMonthlyCreditPayment(loanInterestCreditRate float64, loanDuration float64, loanAmount float64) float64 {
	return gofin.PMT(loanInterestCreditRate, loanDuration, -loanAmount, 0, 0)
}

func CalculateMonthlyInsurancePayment(loanInterestInsuranceRate float64, loanDuration float64, loanAmount float64) float64 {
	return (loanInterestInsuranceRate * loanAmount / 100) / 12
}
