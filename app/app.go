package app

import (
	"fmt"
	"strings"

	"happy_bank_simulator/app/borrowers"
	borrowerViews "happy_bank_simulator/app/borrowers/views"
	configViews "happy_bank_simulator/app/configs/views"
	appHelpers "happy_bank_simulator/app/helpers"
	"happy_bank_simulator/app/insurers"
	insurerViews "happy_bank_simulator/app/insurers/views"
	"happy_bank_simulator/app/lenders"
	lenderViews "happy_bank_simulator/app/lenders/views"
	"happy_bank_simulator/app/loans"
	loanViews "happy_bank_simulator/app/loans/views"
	"happy_bank_simulator/app/transactions"
	transactionViews "happy_bank_simulator/app/transactions/views"
	"happy_bank_simulator/database"
	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/simulation"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func InitApp() {
	app := fyneApp.New()

	masterWindow := app.NewWindow("Happy Bank Simulator")
	masterWindow.Resize(fyne.NewSize(1200, 700))
	appHelpers.SetMasterWindow(&masterWindow)

	configEditView := configViews.RenderEdit()

	runButton := widget.NewButtonWithIcon("Run simulation", theme.ContentAddIcon(), func() {
		simulation.Prepare()
		simulation.Run()
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
	loansController := loans.Controller{}
	borrowersController := borrowers.Controller{}
	lendersController := lenders.Controller{}
	insurersController := insurers.Controller{}
	transactionsController := transactions.Controller{}

	// overviewView := overview.RenderOverview()
	loanIndexView := loanViews.RenderIndex()
	borrowerIndexView := borrowerViews.RenderIndex()
	lenderIndexView := lenderViews.RenderIndex()
	insurerIndexView := insurerViews.RenderIndex()
	transactionIndexView := transactionViews.RenderIndex()

	tabs := container.NewAppTabs(
		// container.NewTabItem("Aperçu", overviewView),
		container.NewTabItem(strings.Title(loansController.GetModelName(true)), loanIndexView),
		container.NewTabItem(strings.Title(borrowersController.GetModelName(true)), borrowerIndexView),
		container.NewTabItem(strings.Title(lendersController.GetModelName(true)), lenderIndexView),
		container.NewTabItem(strings.Title(insurersController.GetModelName(true)), insurerIndexView),
		container.NewTabItem(strings.Title(transactionsController.GetModelName(true)), transactionIndexView),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, tabs)

	dialog := dialog.NewCustom("Simulation results", "Fermer", borderContainer, appHelpers.GetMasterWindow())
	dialog.Resize(fyne.NewSize(1200, 700))
	dialog.Show()
}

func wipeDatabase() {
	databaseHelpers.DropBD()
	fmt.Println("Database dropped")
	database.InitDB()
	fmt.Println("Database initialized")
	databaseHelpers.MigrateDB()
	fmt.Println("Database migrated")
}
