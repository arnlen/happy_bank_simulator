package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/helpers"
	"happy_bank_simulator/models"
)

func createInitialLoans() []*models.Loan {
	var loans []*models.Loan
	quantityOfLoansToCreate := configs.Loan.InitialQuantity
	fmt.Println("Quantity of Loans to create:", quantityOfLoansToCreate)

	for i := 0; i < quantityOfLoansToCreate; i++ {
		loan := models.CreateEmptyLoan()
		fmt.Printf("Loan #%s created\n", strconv.Itoa(int(loan.ID)))

		isThisLoanInsured := helpers.GetResultForProbability(configs.Loan.InsuredQuantityRatio)
		if isThisLoanInsured {
			fmt.Println("This loan is insured")
			loan.IsInsured = true
		} else {
			fmt.Println("This loan is NOT insured 🚨")
			loan.IsInsured = false
		}

		willThisLoanFail := helpers.GetResultForProbability(configs.Loan.FailureRate)
		if willThisLoanFail {
			fmt.Println("This loan will fail 🚨")
			numberOfMonthsBeforeFailure := loan.SetRandomFailureDate()
			fmt.Printf("The failure will occure after %s months, on %s\n", strconv.Itoa(numberOfMonthsBeforeFailure), loan.WillFailOn)
		}

		loans = append(loans, loan)
	}
	return loans
}

func createBorrowersForLoans(loans []*models.Loan) []*models.Borrower {
	var borrowers []*models.Borrower

	for _, loan := range loans {
		borrower := models.CreateDefaultBorrower()
		fmt.Printf("Borrower #%s created\n", strconv.Itoa(int(borrower.ID)))

		if canThisBorrowerTakeThisLoan(borrower, loan) {
			initialDepositAmount := loan.InitialDeposit
			models.CreateDepositTransaction(*borrower, initialDepositAmount)
			borrower.Refresh()
			fmt.Printf("Initial deposit of %s € placed:\n", strconv.Itoa(initialDepositAmount))
			fmt.Printf("- Borrower #%s balance: %s €\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(borrower.Balance))
		}

		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func createMissingLenders(missingQuantity int, availableLenders []*models.Lender) []*models.Lender {
	for i := 0; i < missingQuantity; i++ {
		lender := models.CreateDefaultLender()
		availableLenders = append(availableLenders, lender)
		fmt.Printf("%s/%s - Lender #%s created\n", strconv.Itoa(i+1), strconv.Itoa(missingQuantity), strconv.Itoa(int(lender.ID)))
	}
	fmt.Printf("%s total lenders now available\n", strconv.Itoa(len(availableLenders)))
	return availableLenders
}

func createMissingInsurers(missingQuantity int, availableInsurers []*models.Insurer) []*models.Insurer {
	for i := 0; i < missingQuantity; i++ {
		insurer := models.CreateDefaultInsurer()
		availableInsurers = append(availableInsurers, insurer)
		fmt.Printf("%s/%s - Insurer #%s created\n", strconv.Itoa(i+1), strconv.Itoa(missingQuantity), strconv.Itoa(int(insurer.ID)))
	}
	fmt.Printf("%s total insurers now available\n", strconv.Itoa(len(availableInsurers)))
	return availableInsurers
}
