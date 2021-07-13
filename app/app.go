package app

import (
	borrowerViews "happy_bank_simulator/app/borrowers/views"
	insurerViews "happy_bank_simulator/app/insurers/views"
	lenderViews "happy_bank_simulator/app/lenders/views"
	loanViews "happy_bank_simulator/app/loans/views"
	overview "happy_bank_simulator/app/overview/views"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var (
	app          fyne.App
	masterWindow fyne.Window
)

func InitApp() {
	app = fyneApp.New()

	masterWindow = app.NewWindow("Happy Bank Simulator")
	masterWindow.Resize(fyne.NewSize(1200, 700))

	overviewView := overview.RenderIndex()
	loansView := loanViews.RenderIndex()
	borrowersView := borrowerViews.RenderIndex()
	lendersView := lenderViews.RenderIndex()
	insurersView := insurerViews.RenderIndex()

	tabs := container.NewAppTabs(
		container.NewTabItem("Aperçu", overviewView),
		container.NewTabItem("Crédits", loansView),
		container.NewTabItem("Débiteurs", borrowersView),
		container.NewTabItem("Créanciers", lendersView),
		container.NewTabItem("Assureurs", insurersView),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, tabs)
	masterWindow.SetContent(borderContainer)
	masterWindow.ShowAndRun()
}

func GetApp() fyne.App {
	return app
}

func GetMasterWindow() fyne.Window {
	return masterWindow
}
