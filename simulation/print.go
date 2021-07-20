package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/models"
)

func printSummaryForLoan(loan models.Loan) {
	fmt.Printf("Summary for Loan #%s:\n", strconv.Itoa(int(loan.ID)))
	fmt.Printf("- 1 borrower: %s (#%s)\n", loan.Borrower.Name, strconv.Itoa(int(loan.Borrower.ID)))
	fmt.Printf("- %s lenders\n", strconv.Itoa(len(loan.Lenders)))
	for _, lender := range loan.Lenders {
		fmt.Printf("--- %s (#%s)\n", lender.Name, strconv.Itoa(int(lender.ID)))
	}
	fmt.Printf("- %s insurers\n", strconv.Itoa(len(loan.Insurers)))
	for _, insurer := range loan.Insurers {
		fmt.Printf("--- %s (#%s)\n", insurer.Name, strconv.Itoa(int(insurer.ID)))
	}
}
