package simulation

import (
	"fmt"
	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/models"
	"strconv"
)

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
