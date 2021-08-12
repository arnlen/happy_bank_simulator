package main

import (
	// app "happy_bank_simulator/app"
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/charts"
	"happy_bank_simulator/helpers"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"
)

func main() {
	database.SetupDB()

	prepareSimulation()
	// app.InitApp()
	runSimulation()
}

func prepareSimulation() {
	database.DropBD()
	database.MigrateDB()

	loans := createInitialLoans()
	borrowers := models.CreateBorrowers(len(loans))

	for index, loan := range loans {
		loan.AssignBorrower(borrowers[index])

		loan.SetBorrowerMonthlyIncomes()
		fmt.Printf("🪙 Borrower %d's monthly incomes set to %1.2f.\n",
			int(loan.BorrowerID),
			loan.Borrower.MonthlyIncomes,
		)

		loan.SetupLenders()

		if loan.IsInsured {
			loan.SetupInsurers()
		}

		loan.Print()
	}

	transactions := models.ListTransactions()
	fmt.Println(len(transactions), "transactions in database")

	for _, transaction := range transactions {
		transaction.Print()
	}
}

func createInitialLoans() []*models.Loan {
	loans := models.CreateLoans(configs.Loan.InitialQuantity)
	fmt.Printf("Initial loans created: %d loans\n", len(loans))

	for _, loan := range loans {
		fmt.Printf("Loan #%d setup:\n", int(loan.ID))

		isThisLoanInsured := helpers.GetResultForProbability(configs.Loan.InsuredQuantityRatio)
		if isThisLoanInsured {
			fmt.Println("- This loan is insured")
			loan.IsInsured = true
		} else {
			fmt.Println("- This loan is NOT insured 🚨")
			loan.IsInsured = false
		}

		willThisLoanFail := helpers.GetResultForProbability(configs.Loan.FailureRate)
		if willThisLoanFail {
			fmt.Println("- This loan will fail 🚨")
			loan.SetRandomNumberOfMonthsBeforeFailure()
			fmt.Printf("- The failure will occure after %d months, on %s\n",
				loan.NumberOfMonthsBeforeFailure, loan.WillFailOnString())
		}

		loans = append(loans, loan)
	}
	return loans
}

func runSimulation() {
	fmt.Println("\nRunning a new simulation! 🚀")
	simulationStartDate := helpers.ParseStringToDate(configs.General.StartDate)
	simulationDuration := configs.General.Duration
	chartsManager := charts.ChartsManager{}

	for monthIndex := 0; monthIndex < simulationDuration-1; monthIndex++ {
		currentDate := helpers.AddMonthsToDate(simulationStartDate, monthIndex)

		fmt.Printf("\n--------- Start of Month #%d - %s 📅 ---------\n", monthIndex+1, helpers.TimeDateToString(currentDate))
		loans := models.ListActiveLoans()

		if len(loans) == 0 {
			fmt.Println("No more active loan... 💤 ")
			break
		}
		fmt.Println(len(loans), "active loans")

		for _, loan := range loans {
			loan.Print()

			loanEndDate := helpers.ParseStringToDate(loan.EndDate())
			borrower := loan.Borrower
			lenders := loan.Lenders
			quantityOfLenders := len(lenders)
			insurers := loan.Insurers
			quantityOfInsurers := len(insurers)
			monthString := helpers.TimeDateToString(currentDate)

			// ------- Updates charts for actors -------
			chartsManager.UpdateChartFor([]*models.Actor{&borrower}, monthString)
			chartsManager.UpdateChartFor(lenders, monthString)
			chartsManager.UpdateChartFor(insurers, monthString)

			// ------- Borrower monthly incomes -------
			fmt.Printf("🤑 Borrower #%d got paid %1.2f €!\n",
				int(loan.BorrowerID), borrower.MonthlyIncomes)
			models.CreateIncomeTransaction(borrower, borrower.MonthlyIncomes).Print()

			if currentDate.After(loanEndDate) {
				fmt.Printf("Loan #%d is over. ✅\n", int(loan.ID))
				loan.IsActive = false
				loan.Save()
				continue
			}

			if loan.WillFail() {
				failureDate := loan.WillFailOnTime()

				if currentDate.After(failureDate) {
					fmt.Printf("Loan #%d just fails this month. ❌\n", int(loan.ID))
					loan.IsActive = false
					loan.Save()

					borrower.UpdateBalance(0)
					fmt.Printf("- Borrower #%d's balance: %1.2f €.\n", int(borrower.ID), borrower.Balance)

					if loan.IsInsured {
						fmt.Printf("- Loan #%d is insured by %d insurers. 🆘\n", int(loan.ID), quantityOfInsurers)

						amountLeftToRefund := loan.Amount - loan.RefundedAmount
						amountToRefundByLender := amountLeftToRefund / float64(quantityOfLenders)

						for _, insurer := range loan.Insurers {
							fmt.Printf("--- Insurer #%d will refund %d lenders.\n", int(insurer.ID), quantityOfLenders)

							for _, lender := range lenders {
								models.CreateTransaction(*insurer, *lender, amountToRefundByLender).Print()
							}
						}
						continue

					} else {
						fmt.Printf("- Loan #%d isn't insured. 🕳\n", int(loan.ID))
						continue
					}
				}
			}

			fmt.Printf("Loan #%d has %d lenders, Borrower #%d will pay %1.2f € to each of them. 🏦\n",
				int(loan.ID),
				quantityOfLenders,
				int(borrower.ID),
				loan.MonthlyCredit)
			for _, lender := range loan.Lenders {
				models.CreateTransaction(borrower, *lender, loan.MonthlyCredit).Print()
				loan.Refund(loan.MonthlyCredit)
			}

			if loan.IsInsured {
				fmt.Printf("Loan #%d has %d insurers, Borrower #%d will pay %1.2f € to each of them. 🏥\n",
					int(loan.ID),
					quantityOfInsurers,
					int(borrower.ID),
					loan.MonthlyInsurance)
				for _, insurer := range loan.Insurers {
					models.CreateTransaction(borrower, *insurer, loan.MonthlyInsurance).Print()
				}
			}
		}

		fmt.Printf("\n--------- End of Month #%d - %s ---------\n", monthIndex+1, helpers.TimeDateToString(currentDate))
	}

	chartsManager.DrawChartsFromList()
}
