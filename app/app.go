package app

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

	configEditView := configs.RenderEdit()

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
	databaseHelpers.DropBD()
	fmt.Println("Database dropped")
	database.InitDB()
	fmt.Println("Database initialized")
	databaseHelpers.MigrateDB()
	fmt.Println("Database migrated")
}
