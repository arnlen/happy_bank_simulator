package views

import (
	"fmt"
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
		widget.NewLabelWithData(loansCounterBindingStrings),
		widget.NewLabelWithData(borrowersCounterBindingStrings),
		widget.NewLabelWithData(lendersCounterBindingStrings),
		widget.NewLabelWithData(insurersCounterBindingStrings),
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

	hbox := container.NewHBox(
		populateDatabaseButton,
		wipeDatabaseButton,
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
