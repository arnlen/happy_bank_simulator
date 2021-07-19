package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/models"
)

func assignBorrowerToLoan(borrower *models.Borrower, loan *models.Loan) {
	loan.Borrower = *borrower
	loan.Save()
	borrower.Refresh()
	fmt.Printf("Borrower assigned: Loan #%s's borrower = #%s\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(int(loan.BorrowerID)))
	fmt.Printf("Borrower #%s has now %s loans\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(int(len(borrower.Loans))))
}

func assignLendersToLoan(lenders []*models.Lender, loan *models.Loan) {
	fmt.Printf("Assigning %s lenders to loan #%s\n", strconv.Itoa(len(lenders)), strconv.Itoa(int(loan.ID)))
	for _, availableLender := range lenders {
		loan.AddLender(availableLender)
		fmt.Printf("- Lender #%s assigned\n", strconv.Itoa(int(availableLender.ID)))
	}
	fmt.Printf("Loan #%s has now %s lenders\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(len(loan.Lenders)))
}

func assignInsurersToLoan(insurers []*models.Insurer, loan *models.Loan) {
	fmt.Printf("Assigning %s insurers to loan #%s\n", strconv.Itoa(len(insurers)), strconv.Itoa(int(loan.ID)))
	for _, availableInsurer := range insurers {
		loan.AddInsurer(availableInsurer)
		fmt.Printf("- Insurer #%s assigned\n", strconv.Itoa(int(availableInsurer.ID)))
	}
	fmt.Printf("Loan #%s has now %s insurers\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(len(loan.Insurers)))
}
