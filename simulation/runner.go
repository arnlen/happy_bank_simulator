package simulation

import (
	"fmt"
	"strconv"

	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/models"
)

// TODO NEXT STEPS
//
// 1. Create method #runSimulation(numberOfMonths)
// 2. For every month, take every borrower, and check if it should fail && this month?
// 3a. For those who fail, implement the failure (balance == 0 + log).
// 3b. For those who don't fail, create the transaction borrower => lender + log
// 4. Print a simulation summary including:
// 		- Total loans at the beginning vs total loans at the end (diff = failures)
// 		- Total money in the economy at the beginning vs at the end
//
// - Bind "borrowers", "insurers", "lenders", "loans" and "transactions" to tabs to see ouput of simulation

func Prepare() {
	databaseHelpers.DropBD()
	databaseHelpers.MigrateDB()

	loans := createInitialLoans()
	borrowers := createBorrowersForLoans(loans)

	for index, loan := range loans {
		assignBorrowerToLoan(borrowers[index], loan)
		setupLendersForLoan(loan)

		if loan.IsInsured {
			setupInsurersForLoan(loan)
		}

		printSummaryForLoan(*loan)
	}

	transactions := models.ListTransactions()
	fmt.Println(len(transactions), "transactions in database")

	for _, transaction := range transactions {
		sender := fmt.Sprintf("%s #%s", transaction.SenderType, strconv.Itoa(int(transaction.SenderID)))
		receiver := fmt.Sprintf("%s #%s", transaction.ReceiverType, strconv.Itoa(int(transaction.ReceiverID)))
		fmt.Printf("Transaction #%s from %s to %s of %s €\n", strconv.Itoa(int(transaction.ID)), sender, receiver, strconv.Itoa(transaction.Amount))
	}
}
