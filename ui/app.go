package ui

import (
	"fmt"
	"happy_bank_simulator/models"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var borrowers []models.Borrower
var lenders []models.Lender
var insurers []models.Insurer
var loans []models.Loan

func InitApp(initialBorrowers []models.Borrower, initialLenders []models.Lender, initialInsurers []models.Insurer, initialLoans []models.Loan) {
	borrowers = initialBorrowers
	lenders = initialLenders
	insurers = initialInsurers
	loans = initialLoans

	myApp := app.New()
	myWindow := myApp.NewWindow("Happy Bank Simulator")
	myWindow.Resize(fyne.NewSize(1024, 768))

	overviewTabContent := createOverviewTabContent()
	borrowersTabContent := createBorrowersTabContent()
	lendersTabContent := createLendersTabContent()
	insurersTabContent := createInsurersTabContent()
	loansTabContent := createLoansTabContent()

	tabs := container.NewAppTabs(
		container.NewTabItem("Aperçu", overviewTabContent),
		container.NewTabItem("Crédits", loansTabContent),
		container.NewTabItem("Débiteurs", borrowersTabContent),
		container.NewTabItem("Créanciers", lendersTabContent),
		container.NewTabItem("Assureurs", insurersTabContent),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, tabs)
	myWindow.SetContent(borderContainer)
	myWindow.ShowAndRun()
}

func createOverviewTabContent() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Nombre de crédits : %s", strconv.Itoa(len(loans)))),
		widget.NewLabel(fmt.Sprintf("Nombre de débiteurs : %s", strconv.Itoa(len(borrowers)))),
		widget.NewLabel(fmt.Sprintf("Nombre de créanciers : %s", strconv.Itoa(len(lenders)))),
		widget.NewLabel(fmt.Sprintf("Nombre d'assureurs : %s", strconv.Itoa(len(insurers)))),
	)
}

func createBorrowersTabContent() *fyne.Container {
	vbox := container.NewVBox()

	for _, borrower := range borrowers {
		borrowerString := fmt.Sprintf("%s - %9.2f - Loans: %v", borrower.Name, borrower.Balance, borrower.Loans)
		label := widget.NewLabel(borrowerString)
		deleteButton := widget.NewButton("Suppr", func() {
			log.Println(borrowerString)
		})

		borderContainer := container.NewBorder(nil, nil, label, deleteButton)
		vbox.Add(borderContainer)
	}

	// data := binding.BindStringList(
	// 	&[]string{},
	// )

	// list := widget.NewListWithData(data,
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("template")
	// 	},
	// 	func(i binding.DataItem, o fyne.CanvasObject) {
	// 		o.(*widget.Label).Bind(i.(binding.String))
	// 	})

	return vbox
}

func createLendersTabContent() *fyne.Container {
	vbox := container.NewVBox()

	for _, lender := range lenders {
		lenderString := fmt.Sprintf("%s - %9.2f - Loans: %v", lender.Name, lender.Balance, lender.Loans)
		label := widget.NewLabel(lenderString)
		deleteButton := widget.NewButton("Suppr", func() {
			log.Println(lenderString)
		})

		borderContainer := container.NewBorder(nil, nil, label, deleteButton)
		vbox.Add(borderContainer)
	}

	return vbox
}

func createInsurersTabContent() *fyne.Container {
	vbox := container.NewVBox()

	for _, insurer := range insurers {
		insurerString := fmt.Sprintf("%s - %9.2f - Loans: %v", insurer.Name, insurer.Balance, insurer.Loans)
		label := widget.NewLabel(insurerString)
		deleteButton := widget.NewButton("Suppr", func() {
			log.Println(insurerString)
		})

		borderContainer := container.NewBorder(nil, nil, label, deleteButton)
		vbox.Add(borderContainer)
	}

	return vbox
}

func createLoansTabContent() *fyne.Container {
	vbox := container.NewVBox()

	for _, loan := range loans {
		loanString := fmt.Sprintf("%i - %i - %i", loan.BorrowerID, loan.LenderID, loan.InsurerID)
		label := widget.NewLabel(loanString)
		deleteButton := widget.NewButton("Suppr", func() {
			log.Println(loanString)
		})

		borderContainer := container.NewBorder(nil, nil, label, deleteButton)
		vbox.Add(borderContainer)
	}

	return vbox
}
