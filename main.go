package main

import (
	"fmt"
	"strings"

	"happy_bank_simulator/app/borrowers"
	"happy_bank_simulator/app/configs"
	appHelpers "happy_bank_simulator/app/helpers"
	"happy_bank_simulator/app/insurers"
	"happy_bank_simulator/app/lenders"
	"happy_bank_simulator/app/loans"
	"happy_bank_simulator/app/transactions"
	"happy_bank_simulator/charts"
	"happy_bank_simulator/helpers"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	database.SetupDB()

	prepareSimulation()
	// app.InitApp()
	runSimulation()
}

func prepareSimulation() {
	loans := createInitialLoans()
	borrowers := models.CreateBorrowers(len(loans))

	for index, loan := range loans {
		borrower := borrowers[index]
		loan.AssignBorrower(borrower)

		loan.SetBorrowerMonthlyIncomes()
		borrower.Refresh()
		fmt.Printf("🪙  Borrower #%d's monthly incomes set to %1.2f €.\n",
			int(borrower.ID),
			borrower.MonthlyIncomes,
		)

		loan.SetupLenders()

		if loan.IsInsured {
			loan.SetupInsurers()
		}

		loan.Activate()
		loan.Print()
	}

	transactions := models.ListTransactions()
	fmt.Println(len(transactions), "transaction(s) in database")

	for _, transaction := range transactions {
		transaction.Print()
	}
}

func createInitialLoans() []*models.Loan {
	loans := models.CreateLoans(configs.Loan.InitialQuantity)
	fmt.Printf("Initial loans created: %d loan(s)\n", len(loans))

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
			fmt.Printf("- The failure will occure after %d month(s), on %s\n",
				loan.NumberOfMonthsBeforeFailure,
				loan.WillFailOnString(),
			)
		}
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

			chartsManager.UpdateChartFor([]*models.Actor{&borrower}, monthString)
			chartsManager.UpdateChartFor(lenders, monthString)
			chartsManager.UpdateChartFor(insurers, monthString)
			payBorrower(&borrower)

			if currentDate.Equal(loanEndDate) {
				fmt.Printf("Loan #%d is over. ✅\n", int(loan.ID))
				loan.IsActive = false
				loan.Save()
				continue
			}

			fmt.Printf("Loan #%d has %d lenders, Borrower #%d will pay %1.2f € to each of them. 🏦\n",
				int(loan.ID),
				quantityOfLenders,
				int(borrower.ID),
				loan.MonthlyCredit)
			loan.MakeLendersMonthlyPayments()

			if loan.IsInsured {
				fmt.Printf("Loan #%d has %d insurers, Borrower #%d will pay %1.2f € to each of them. 🏥\n",
					int(loan.ID),
					quantityOfInsurers,
					int(borrower.ID),
					loan.MonthlyInsurance)
				loan.MakeInsurersMonthlyPayments()
			}

			// if loan.WillFail() {
			// 	failureDate := loan.WillFailOnTime()

			// 	if currentDate.After(failureDate) {
			// 		fmt.Printf("Loan #%d just fails this month. ❌\n", int(loan.ID))
			// 		loan.IsActive = false
			// 		loan.Save()

			// 		fmt.Printf("- Borrower #%d's balance: %1.2f €.\n", int(borrower.ID), borrower.Balance)

			// 		if loan.IsInsured {
			// 			fmt.Printf("- Loan #%d is insured by %d insurers. 🆘\n", int(loan.ID), quantityOfInsurers)

			// 			amountLeftToRefund := loan.Amount - loan.RefundedAmount
			// 			amountToRefundByLender := amountLeftToRefund / float64(quantityOfLenders)

			// 			for _, insurer := range loan.Insurers {
			// 				fmt.Printf("--- Insurer #%d will refund %d lenders.\n", int(insurer.ID), quantityOfLenders)

			// 				for _, lender := range lenders {
			// 					models.CreateTransaction(*insurer, *lender, amountToRefundByLender).Print()
			// 				}
			// 			}
			// 			continue

			// 		} else {
			// 			fmt.Printf("- Loan #%d isn't insured. 🕳\n", int(loan.ID))
			// 			continue
			// 		}
			// 	}
			// }
		}

		fmt.Printf("\n--------- End of Month #%d - %s ---------\n", monthIndex+1, helpers.TimeDateToString(currentDate))
	}

	chartsManager.DrawChartsFromList()
}

func InitApp() {
	app := fyneApp.New()

	masterWindow := app.NewWindow("Happy Bank Simulator")
	masterWindow.Resize(fyne.NewSize(1200, 700))
	appHelpers.SetMasterWindow(&masterWindow)

	configEditView := configs.RenderEdit()

	runButton := widget.NewButtonWithIcon("Run simulation", theme.ContentAddIcon(), func() {
		prepareSimulation()
		runSimulation()
		renderSimulationResultsWindow()
	})

	wipeDatabaseButton := widget.NewButton("Vider la base de données", func() {
		fmt.Println("Wipe button tapped")
		wipeDatabase()
	})

	hbox := container.NewHBox(
		wipeDatabaseButton,
		runButton,
	)

	masterBorderContainer := container.NewBorder(nil, hbox, nil, nil, configEditView)

	masterWindow.SetContent(masterBorderContainer)
	masterWindow.ShowAndRun()
}

func renderSimulationResultsWindow() {

	// overviewView := overview.RenderOverview()
	loanIndexView := loans.RenderIndex()
	borrowerIndexView := borrowers.RenderIndex()
	lenderIndexView := lenders.RenderIndex()
	insurerIndexView := insurers.RenderIndex()
	transactionIndexView := transactions.RenderIndex()

	tabs := container.NewAppTabs(
		container.NewTabItem(strings.Title(configs.Loan.String), loanIndexView),
		container.NewTabItem(strings.Title(configs.Actor.BorrowerString), borrowerIndexView),
		container.NewTabItem(strings.Title(configs.Actor.LenderString), lenderIndexView),
		container.NewTabItem(strings.Title(configs.Actor.InsurerString), insurerIndexView),
		container.NewTabItem(strings.Title(configs.Transaction.String), transactionIndexView),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, tabs)

	dialog := dialog.NewCustom("Simulation results", "Fermer", borderContainer, appHelpers.GetMasterWindow())
	dialog.Resize(fyne.NewSize(1200, 700))
	dialog.Show()
}

func wipeDatabase() {
	database.ResetDB()
}

func payBorrower(borrower *models.Actor) {
	models.CreateIncomeTransaction(*borrower, borrower.MonthlyIncomes).Print()

	fmt.Printf("🤑 Borrower #%d got paid %1.2f €!\n",
		int(borrower.ID), borrower.MonthlyIncomes)
}

// func updateChartsWith(chartsManager charts.ChartsManager, borrower *models.Actor, lenders []*models.Actor, insurers []*models.Actor, monthString string) {
// 	chartsManager.UpdateChartFor([]*models.Actor{borrower}, monthString)
// 	chartsManager.UpdateChartFor(lenders, monthString)
// 	chartsManager.UpdateChartFor(insurers, monthString)
// }
