package views

import (
	"happy_bank_simulator/models"
	viewBorrowers "happy_bank_simulator/views/borrowers"
	viewInsurers "happy_bank_simulator/views/insurers"
	viewLenders "happy_bank_simulator/views/lenders"
	viewLoans "happy_bank_simulator/views/loans"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func InitApp(borrowers []models.Borrower, lenders []models.Lender, insurers []models.Insurer, loans []models.Loan) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Happy Bank Simulator")
	myWindow.Resize(fyne.NewSize(1200, 700))

	overviewView := Overview(len(loans), len(borrowers), len(lenders), len(insurers))
	loansView := viewLoans.Index(loans)
	borrowersView := viewBorrowers.Index(borrowers)
	lendersView := viewLenders.Index(lenders)
	insurersView := viewInsurers.Index(insurers)

	tabs := container.NewAppTabs(
		container.NewTabItem("Aperçu", overviewView),
		container.NewTabItem("Crédits", loansView),
		container.NewTabItem("Débiteurs", borrowersView),
		container.NewTabItem("Créanciers", lendersView),
		container.NewTabItem("Assureurs", insurersView),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, tabs)
	myWindow.SetContent(borderContainer)
	myWindow.ShowAndRun()
}
