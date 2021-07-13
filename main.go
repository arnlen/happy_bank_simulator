package main

import (
	a "happy_bank_simulator/app"
	"happy_bank_simulator/models"
	"happy_bank_simulator/services/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"gorm.io/gorm/clause"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	db := database.GetDB()

	var borrowers []models.Borrower
	var lenders []models.Lender
	var insurers []models.Insurer
	var loans []models.Loan

	db.Preload(clause.Associations).Find(&borrowers)
	db.Preload(clause.Associations).Find(&lenders)
	db.Preload(clause.Associations).Find(&insurers)
	db.Preload(clause.Associations).Find(&loans)

	initApp(borrowers, lenders, insurers, loans)
}

func initApp(borrowers []models.Borrower, lenders []models.Lender, insurers []models.Insurer, loans []models.Loan) {
	app := app.New()
	a.SetApp(&app)

	masterWindow := app.NewWindow("Happy Bank Simulator")
	masterWindow.Resize(fyne.NewSize(1200, 700))
	a.SetMasterWindow(&masterWindow)

	// overviewView := overview.Controller{}
	// loansView := viewLoans.Index(loans)
	// borrowersView := viewBorrowers.Index(borrowers)
	// lendersView := viewLenders.Index(lenders)
	// insurersView := viewInsurers.Index(insurers)

	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Aperçu", overviewView),
	// 	container.NewTabItem("Crédits", loansView),
	// 	container.NewTabItem("Débiteurs", borrowersView),
	// 	container.NewTabItem("Créanciers", lendersView),
	// 	container.NewTabItem("Assureurs", insurersView),
	// )

	// tabs.SetTabLocation(container.TabLocationLeading)

	borderContainer := container.NewBorder(nil, nil, nil, nil, nil)
	a.GetMasterWindow().SetContent(borderContainer)
	masterWindow.ShowAndRun()
}
