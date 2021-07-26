package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/app/configs"
	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/helpers"
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
// - Create transaction for every loan assignation
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
		transaction.Print()
	}
}

func Run() {
	fmt.Println("\nRunning a new simulation! 🚀")
	simulationStartDate := helpers.ParseStringToDate(configs.General.StartDate)
	simulationDuration := configs.General.Duration

	for monthIndex := 0; monthIndex < simulationDuration-1; monthIndex++ {
		currentDate := helpers.AddMonthsToDate(simulationStartDate, monthIndex)
		fmt.Printf("Month %s - %s 📅\n", strconv.Itoa(monthIndex), helpers.TimeDateToString(currentDate))
		loans := models.ListLoans()

		for _, loan := range loans {
			loan.UpdateActiveStatus(currentDate)
			if loan.IsActive {
				fmt.Printf("Loan #%s is active.\n", strconv.Itoa(int(loan.ID)))
				borrower := loan.Borrower

				if helpers.ParseStringToDate(loan.WillFailOn) == currentDate {
					fmt.Printf("Loan #%s is failed.\n", strconv.Itoa(int(loan.ID)))
					borrower.UpdateBalance(0)
					fmt.Printf("Borrower #%s updated to 0.\n", strconv.Itoa(int(borrower.ID)))
					continue
				}

				for _, lender := range loan.Lenders {
					models.CreateTransaction(&borrower, lender, int(loan.MonthlyCredit)).Print()
				}

				if loan.IsInsured {
					for _, insurer := range loan.Insurers {
						models.CreateTransaction(&borrower, insurer, int(loan.MonthlyInsurance)).Print()
					}
				}
			} else {
				fmt.Printf("Loan #%s is inactive.\n", strconv.Itoa(int(loan.ID)))
			}
		}
	}
}
