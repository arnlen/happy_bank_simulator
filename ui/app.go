package ui

import (
	"fmt"
	"happy_bank_simulator/models"
	"log"
	"strconv"

	// "strconv"

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

	// overviewTabContent := createOverviewTabContent()
	loansTabContent := createLoansTabContent()

	borrowersTabContent := createBorrowersTabContent()
	// lendersTabContent := createLendersTabContent()
	// insurersTabContent := createInsurersTabContent()

	tabs := container.NewAppTabs(
		// container.NewTabItem("Aperçu", overviewTabContent),
		container.NewTabItem("Crédits", loansTabContent),
		container.NewTabItem("Débiteurs", borrowersTabContent),
		// container.NewTabItem("Créanciers", lendersTabContent),
		// container.NewTabItem("Assureurs", insurersTabContent),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, tabs)
	myWindow.SetContent(borderContainer)
	myWindow.ShowAndRun()
}

// func createOverviewTabContent() *fyne.Container {
// 		return container.NewVBox(
// 			widget.NewLabel(fmt.Sprintf("Nombre de crédits : %s", strconv.Itoa(len(loans)))),
// 			widget.NewLabel(fmt.Sprintf("Nombre de débiteurs : %s", strconv.Itoa(len(borrowers)))),
// 			widget.NewLabel(fmt.Sprintf("Nombre de créanciers : %s", strconv.Itoa(len(lenders)))),
// 			widget.NewLabel(fmt.Sprintf("Nombre d'assureurs : %s", strconv.Itoa(len(insurers)))),
// 		)
// }

func createBorrowersTabContent() *fyne.Container {
	vbox := container.NewVBox()

	for _, borrower := range borrowers {
		borrowerString := fmt.Sprintf("%s - %.0f €", borrower.Name, borrower.Balance)
		label := widget.NewLabel(borrowerString)
		deleteButton := widget.NewButton("Suppr", func() {
			log.Println(borrowerString)
		})

		borderContainer := container.NewBorder(nil, nil, label, deleteButton)
		vbox.Add(borderContainer)
	}

	return vbox
}

// func createLendersTabContent() *fyne.Container {
// 	vbox := container.NewVBox()

// 	for _, lender := range lenders {
// 		lenderString := fmt.Sprintf("%s - %9.2f - Loans: %v", lender.Name, lender.Balance, lender.Loans)
// 		loanString := fmt.Sprintf("Loan #%v", lender.Loans[0].ID)

// 		label := widget.NewLabel(lenderString + " " + loanString)
// 		deleteButton := widget.NewButton("Suppr", func() {
// 			log.Println(lenderString)
// 		})

// 		borderContainer := container.NewBorder(nil, nil, label, deleteButton)
// 		vbox.Add(borderContainer)
// 	}

// 	return vbox
// }

// func createInsurersTabContent() *fyne.Container {
// 	vbox := container.NewVBox()

// 	for _, insurer := range insurers {
// 		insurerString := fmt.Sprintf("%s - %9.2f - Loans: %v", insurer.Name, insurer.Balance, insurer.Loans)
// 		label := widget.NewLabel(insurerString)
// 		deleteButton := widget.NewButton("Suppr", func() {
// 			log.Println(insurerString)
// 		})

// 		borderContainer := container.NewBorder(nil, nil, label, deleteButton)
// 		vbox.Add(borderContainer)
// 	}

// 	return vbox
// }

func createLoansTabContent() *fyne.Container {
	var loansTableData = [][]string{
		{"ID", "Débiteur", "Créancier", "Assureur", "Montant", "Durée"},
		{"1", "Arnaud Lenglet", "Wendy Lenglet", "Adrien Lenglet", "100 €", "12 mois"}}

	// vbox := container.NewVBox()

	for _, loan := range loans {
		loanRow := []string{
			strconv.Itoa(int(loan.ID)),
			loan.Borrower.Name,
			loan.Lender.Name,
			loan.Insurer.Name,
			fmt.Sprintf("%8.2f €", loan.Amount),
			strconv.Itoa(int(loan.Duration)),
		}

		loansTableData = append(loansTableData, loanRow)
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(loansTableData), len(loansTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(loansTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 250)
	table.SetColumnWidth(3, 250)
	table.SetColumnWidth(4, 100)

	return container.NewBorder(nil, nil, nil, nil, table)
}
