package app

import (
	"happy_bank_simulator/app/borrowers"
	borrowerViews "happy_bank_simulator/app/borrowers/views"
	"happy_bank_simulator/app/configs"
	appHelpers "happy_bank_simulator/app/helpers"
	"happy_bank_simulator/app/insurers"
	insurerViews "happy_bank_simulator/app/insurers/views"
	"happy_bank_simulator/app/lenders"
	lenderViews "happy_bank_simulator/app/lenders/views"
	"happy_bank_simulator/app/loans"
	loanViews "happy_bank_simulator/app/loans/views"
	overview "happy_bank_simulator/app/overview/views"
	"strings"

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

	configsView := configs.RenderConfigs()
	overviewView := overview.RenderOverview()
	loansView := loanViews.RenderIndex()
	borrowersView := borrowerViews.RenderIndex()
	lendersView := lenderViews.RenderIndex()
	insurersView := insurerViews.RenderIndex()

	tabs := container.NewAppTabs(
		container.NewTabItem("Paramètres", configsView),
		container.NewTabItem("Aperçu", overviewView),
		container.NewTabItem(strings.Title(loansController.GetModelName(true)), loansView),
		container.NewTabItem(strings.Title(borrowersController.GetModelName(true)), borrowersView),
		container.NewTabItem(strings.Title(lendersController.GetModelName(true)), lendersView),
		container.NewTabItem(strings.Title(insurersController.GetModelName(true)), insurersView),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, tabs)
	masterWindow.SetContent(borderContainer)
	masterWindow.ShowAndRun()
}
