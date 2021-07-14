package views

import (
	"fmt"
	"happy_bank_simulator/app/charts"
	"happy_bank_simulator/app/overview"

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

// Initialize controller
var overviewController = overview.Controller{}

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
		overviewController.PopulateDatabase()
		updateCounters(overviewController.GetCounters())
	})

	wipeDatabaseButton := widget.NewButton("Vider la base de données", func() {
		fmt.Println("Wipe button tapped")
		overviewController.WipeDatabase()
		updateCounters(overviewController.GetCounters())
	})

	generateChartsButton := widget.NewButton("Générer les graphes", func() {
		fmt.Println("Generate chartes button tapped")
		charts.RenderChart()
	})

	hbox := container.NewHBox(
		populateDatabaseButton,
		wipeDatabaseButton,
		generateChartsButton,
	)

	updateCounters(overviewController.GetCounters())

	return container.NewBorder(nil, hbox, nil, nil, vbox)
}

func updateCounters(counters []int) {
	borrowersCounterBinding.Set(counters[0])
	lendersCounterBinding.Set(counters[1])
	insurersCounterBinding.Set(counters[2])
	loansCounterBinding.Set(counters[3])
	fmt.Println("Overview counters updated")
}
