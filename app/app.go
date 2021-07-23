package app

import (
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
	overview "happy_bank_simulator/app/overview/views"
	"happy_bank_simulator/app/transactions"
	transactionViews "happy_bank_simulator/app/transactions/views"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func InitApp() {
	app := fyneApp.New()

	masterWindow := app.NewWindow("Happy Bank Simulator")
	masterWindow.Resize(fyne.NewSize(1200, 700))
	appHelpers.SetMasterWindow(&masterWindow)

	var borrowersController = borrowers.Controller{}
	var insurersController = insurers.Controller{}
	var lendersController = lenders.Controller{}
	var loansController = loans.Controller{}
	var transactionsController = transactions.Controller{}

	configEditView := configViews.RenderEdit()
	overviewView := overview.RenderOverview()
	loanIndexView := loanViews.RenderIndex()
	borrowerIndexView := borrowerViews.RenderIndex()
	lenderIndexView := lenderViews.RenderIndex()
	insurerIndexView := insurerViews.RenderIndex()
	transactionIndexView := transactionViews.RenderIndex()

	tabs := container.NewAppTabs(
		container.NewTabItem("Paramètres", configEditView),
		container.NewTabItem("Aperçu", overviewView),
		container.NewTabItem(strings.Title(loansController.GetModelName(true)), loanIndexView),
		container.NewTabItem(strings.Title(borrowersController.GetModelName(true)), borrowerIndexView),
		container.NewTabItem(strings.Title(lendersController.GetModelName(true)), lenderIndexView),
		container.NewTabItem(strings.Title(insurersController.GetModelName(true)), insurerIndexView),
		container.NewTabItem(strings.Title(transactionsController.GetModelName(true)), transactionIndexView),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, tabs)
	masterWindow.SetContent(borderContainer)
	masterWindow.ShowAndRun()
}
