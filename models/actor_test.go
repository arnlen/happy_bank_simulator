package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActor_GetNetBalance(t *testing.T) {
	assert := assert.New(t)

	borrowerWithoutLoan := &Actor{
		Name:           "Test Actor",
		Loans:          []*Loan{},
		InitialBalance: 1000,
		Balance:        500,
		Type:           "borrower",
	}

	insurer := &Actor{
		Name:           "Test Actor",
		Loans:          []*Loan{},
		InitialBalance: 1000,
		Balance:        500,
		Type:           "insurer",
	}

	lender := &Actor{
		Name:           "Test Actor",
		Loans:          []*Loan{},
		InitialBalance: 1000,
		Balance:        500,
		Type:           "lender",
	}

	var tests = []struct {
		input    Actor
		expected float64
	}{
		{*borrowerWithoutLoan, 500},
		{*insurer, 500},
		{*lender, 500},
	}

	for _, test := range tests {
		assert.Equal(test.input.GetNetBalance(), test.expected)
	}
}
