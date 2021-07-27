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

		loan.Print()
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
		fmt.Printf("\n--------- Start of Month #%s - %s 📅 ---------\n", strconv.Itoa(monthIndex+1), helpers.TimeDateToString(currentDate))
		loans := models.ListActiveLoans()

		if len(loans) == 0 {
			fmt.Println("No more active loan... 💤 ")
			break
		}
		fmt.Println(len(loans), "active loans")

		for _, loan := range loans {
			loanEndDate := helpers.ParseStringToDate(loan.EndDate)
			borrower := loan.Borrower
			lenders := loan.Lenders
			quantityOfLenders := len(lenders)
			insurers := loan.Insurers
			quantityOfInsurers := len(insurers)

			loan.Print()

			if currentDate.After(loanEndDate) {
				fmt.Printf("Loan #%s is over. ✅\n", strconv.Itoa(int(loan.ID)))
				loan.IsActive = false
				loan.Save()
				continue
			}

			if loan.WillFailOn != "" {
				failureDate := helpers.ParseStringToDate(loan.WillFailOn)

				if currentDate.After(failureDate) {
					fmt.Printf("Loan #%s just fails this month. ❌\n", strconv.Itoa(int(loan.ID)))
					loan.IsActive = false
					loan.Save()

					borrower.UpdateBalance(0)
					fmt.Printf("- Borrower #%s's balance: 0 €.\n", strconv.Itoa(int(borrower.ID)))

					if loan.IsInsured {
						fmt.Printf("- Loan #%s is insured by %s insurers. 🆘\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(quantityOfInsurers))

						amountLeftToRefund := loan.Amount - loan.RefundedAmount
						amountToRefundByLender := amountLeftToRefund / float64(quantityOfLenders)

						for _, insurer := range loan.Insurers {
							fmt.Printf("--- Insurer #%s will refund %s lenders.\n", strconv.Itoa(int(insurer.ID)), strconv.Itoa(quantityOfLenders))

							for _, lender := range lenders {
								models.CreateTransaction(insurer, lender, amountToRefundByLender).Print()
							}
						}
						continue

					} else {
						fmt.Printf("- Loan #%s isn't insured. 🕳\n", strconv.Itoa(int(loan.ID)))
						continue
					}
				}
			}

			fmt.Printf("Loan #%s has %s lenders, Borrower #%s will pay %1.2f € to each of them. 🏦\n",
				strconv.Itoa(int(loan.ID)),
				strconv.Itoa(quantityOfLenders),
				strconv.Itoa(int(borrower.ID)),
				loan.MonthlyCredit)
			for _, lender := range loan.Lenders {
				models.CreateTransaction(&borrower, lender, loan.MonthlyCredit).Print()
				loan.Refund(loan.MonthlyCredit)
			}

			if loan.IsInsured {
				fmt.Printf("Loan #%s has %s insurers, Borrower #%s will pay %1.2f € to each of them. 🏥\n",
					strconv.Itoa(int(loan.ID)),
					strconv.Itoa(quantityOfInsurers),
					strconv.Itoa(int(borrower.ID)),
					loan.MonthlyInsurance)
				for _, insurer := range loan.Insurers {
					models.CreateTransaction(&borrower, insurer, loan.MonthlyInsurance).Print()
				}
			}
		}

		fmt.Printf("\n--------- End of Month #%s - %s ---------\n", strconv.Itoa(monthIndex+1), helpers.TimeDateToString(currentDate))
	}
}
