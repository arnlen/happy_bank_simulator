package main

import (
	"testing"
)

func TestCalculateMonthlyCreditPayment(t *testing.T) {
	loanInterestCreditRate := 3
	loanDuration := 60
	loanAmount := 10000

	expected := 179.6869066406344
	result := CalculateMonthlyCreditPayment(float64(loanInterestCreditRate), float64(loanDuration), float64(loanAmount))

	if result != expected {
		t.Fatalf(`CalculateMonthlyCreditPayment() = %v expected %v`, result, expected)
	}
}

func TestCalculateMonthlyInsurancePayment(t *testing.T) {
	loanInterestInsuranceRate := 3
	loanDuration := 60
	loanAmount := 10000

	expected := float64(25)
	result := CalculateMonthlyInsurancePayment(float64(loanInterestInsuranceRate), float64(loanDuration), float64(loanAmount))

	if result != expected {
		t.Fatalf(`CalculateMonthlyInsurancePayment() = %v expected %v`, result, expected)
	}
}
