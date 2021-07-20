package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/helpers"
	"happy_bank_simulator/models"
)

// TODO
//
// For borrower
// - Place the initial security deposit
//

func createLoans() {
	quantityOfLoansToCreate = configs.Loan.InitialQuantity
	fmt.Println("Quantity of Loans to create:", quantityOfLoansToCreate)

	for i := 0; i < quantityOfLoansToCreate; i++ {
		loan := createEmptyLoan()
		willThisLoanFail := helpers.GetResultForProbability(configs.Loan.FailureRate)
		if willThisLoanFail {
			fmt.Println("This loan will fail 🚨")
			loan.SetRandomFailureDate()
		}

		borrower := createDefaultBorrower()
		if canThisBorrowerTakeThisLoan(borrower, loan) {

			initialDepositAmount := loan.InitialDeposit
			models.NewTransaction(borrower, &depositAccount, initialDepositAmount)
			fmt.Printf("Initial deposit of %s € placed:\n", strconv.Itoa(initialDepositAmount))
			fmt.Printf("- Borrower #%s balance: %s €\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(borrower.Balance))
			fmt.Printf("- DepositAccount balance: %s €\n", strconv.Itoa(depositAccount.Balance))

			assignBorrowerToLoan(borrower, loan)

			setupLendersForLoan(loan)

			isThisLoanInsured := helpers.GetResultForProbability(configs.Loan.InsuredQuantityRatio)
			if isThisLoanInsured {
				fmt.Println("This loan is insured")
				setupInsurersForLoan(loan)
			} else {
				fmt.Println("This loan is NOT insured 🚨")
			}
		}

		printSummaryForLoan(*loan)
	}
}

func createEmptyLoan() *models.Loan {
	var loan = models.NewDefaultLoan()
	loan.Save()
	fmt.Printf("Loan #%s created\n", strconv.Itoa(int(loan.ID)))
	return loan
}

func createDefaultBorrower() *models.Borrower {
	borrower := models.NewDefaultBorrower()
	borrower.Save()
	fmt.Printf("Borrower #%s created\n", strconv.Itoa(int(borrower.ID)))
	return borrower
}

func createMissingLenders(missingQuantity int, availableLenders []*models.Lender) []*models.Lender {
	for i := 0; i < missingQuantity; i++ {
		lender := models.NewDefaultLender()
		lender.Save()
		availableLenders = append(availableLenders, lender)
		fmt.Printf("%s/%s - Lender #%s created\n", strconv.Itoa(i+1), strconv.Itoa(missingQuantity), strconv.Itoa(int(lender.ID)))
	}
	fmt.Printf("%s total lenders now available\n", strconv.Itoa(len(availableLenders)))
	return availableLenders
}

func createMissingInsurers(missingQuantity int, availableInsurers []*models.Insurer) []*models.Insurer {
	for i := 0; i < missingQuantity; i++ {
		insurer := models.NewDefaultInsurer()
		insurer.Save()
		availableInsurers = append(availableInsurers, insurer)
		fmt.Printf("%s/%s - Insurer #%s created\n", strconv.Itoa(i+1), strconv.Itoa(missingQuantity), strconv.Itoa(int(insurer.ID)))
	}
	fmt.Printf("%s total insurers now available\n", strconv.Itoa(len(availableInsurers)))
	return availableInsurers
}
