package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/models"
)

func assignBorrowerToLoan(borrower *models.Actor, loan *models.Loan) {
	loan.Borrower = *borrower
	loan.Save()
	borrower.Refresh()
	fmt.Printf("Borrower assigned: Loan #%s's borrower = #%s\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(int(loan.BorrowerID)))
	fmt.Printf("Borrower #%s has now %s loans\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(int(len(borrower.Loans))))
}

func assignLendersToLoan(lenders []*models.Actor, loan *models.Loan) {
	totalAmountLent := 0.0
	amountToLend := calculateAmountToLendForLender(loan.Amount)

	fmt.Printf("Assigning %s lenders to loan #%s\n", strconv.Itoa(len(lenders)), strconv.Itoa(int(loan.ID)))
	for _, availableLender := range lenders {
		totalAmountLent += amountToLend
		models.CreateTransaction(*availableLender, loan.Borrower, amountToLend).Print()
		loan.AddLender(availableLender)
		fmt.Printf("- Lender #%s assigned. InitialBalance: %1.2f | Balance: %1.2f\n", strconv.Itoa(int(availableLender.ID)), availableLender.InitialBalance, availableLender.Balance)
	}
	fmt.Printf("Loan #%s has now %s lenders. Total amout lent: %1.2f € (%1.2f/lender) \n",
		strconv.Itoa(int(loan.ID)),
		strconv.Itoa(len(loan.Lenders)),
		totalAmountLent,
		amountToLend,
	)
}

func assignInsurersToLoan(insurers []*models.Actor, loan *models.Loan) {
	fmt.Printf("Assigning %s insurers to loan #%s\n", strconv.Itoa(len(insurers)), strconv.Itoa(int(loan.ID)))
	for _, availableInsurer := range insurers {
		loan.AddInsurer(availableInsurer)
		fmt.Printf("- Insurer #%s assigned\n", strconv.Itoa(int(availableInsurer.ID)))
	}
	fmt.Printf("Loan #%s has now %s insurers\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(len(loan.Insurers)))
}
