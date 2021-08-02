package overview

import (
	"fmt"

	"happy_bank_simulator/database"
	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/models"

	"gorm.io/gorm/clause"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Declare bindings
var loansCounterBinding = binding.NewInt()
var borrowersCounterBinding = binding.NewInt()
var lendersCounterBinding = binding.NewInt()
var insurersCounterBinding = binding.NewInt()

func RenderOverview() *fyne.Container {
	loansCounterBindingStrings := binding.IntToString(loansCounterBinding)
	borrowersCounterBindingStrings := binding.IntToString(borrowersCounterBinding)
	lendersCounterBindingStrings := binding.IntToString(lendersCounterBinding)
	insurersCounterBindingStrings := binding.IntToString(insurersCounterBinding)

	vbox := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Nombres de prêts :"),
			widget.NewLabelWithData(loansCounterBindingStrings),
		),
		container.NewHBox(
			widget.NewLabel("Nombres d'emprunteurs :"),
			widget.NewLabelWithData(borrowersCounterBindingStrings),
		),
		container.NewHBox(
			widget.NewLabel("Nombres de prêteurs :"),
			widget.NewLabelWithData(lendersCounterBindingStrings),
		),
		container.NewHBox(
			widget.NewLabel("Nombres d'assureurs :"),
			widget.NewLabelWithData(insurersCounterBindingStrings),
		),
	)

	populateDatabaseButton := widget.NewButton("Remplir la base", func() {
		fmt.Println("Populate button tapped")
		populateDatabase()
		updateCounters(getCounters())
	})

	wipeDatabaseButton := widget.NewButton("Vider la base de données", func() {
		fmt.Println("Wipe button tapped")
		wipeDatabase()
		updateCounters(getCounters())
	})

	generateChartsButton := widget.NewButton("Générer les graphes", func() {
		fmt.Println("Not implemented")
	})

	hbox := container.NewHBox(
		populateDatabaseButton,
		wipeDatabaseButton,
		generateChartsButton,
	)

	refreshButton := widget.NewButton("Refraichir", func() {
		updateCounters(getCounters())
		fmt.Println("Refreshed!")
	})

	updateCounters(getCounters())

	return container.NewBorder(refreshButton, hbox, nil, nil, vbox)
}

func updateCounters(counters []int) {
	borrowersCounterBinding.Set(counters[0])
	lendersCounterBinding.Set(counters[1])
	insurersCounterBinding.Set(counters[2])
	loansCounterBinding.Set(counters[3])
	fmt.Println("Overview counters updated")
}

func getCounters() []int {
	db := database.GetDB()

	var borrowers []models.Actor
	var lenders []models.Actor
	var insurers []models.Actor
	var loans []models.Loan

	db.Preload(clause.Associations).Find(&borrowers)
	db.Preload(clause.Associations).Find(&lenders)
	db.Preload(clause.Associations).Find(&insurers)
	db.Preload(clause.Associations).Find(&loans)

	return []int{len(borrowers), len(lenders), len(insurers), len(loans)}
}

func populateDatabase() {
	fmt.Println("Deactivated")
}

func wipeDatabase() {
	databaseHelpers.DropBD()
	fmt.Println("Database dropped")
	database.InitDB()
	fmt.Println("Database initialized")
	databaseHelpers.MigrateDB()
	fmt.Println("Database migrated")
}
