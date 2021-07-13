package views

import (
	"fmt"
	"happy_bank_simulator/app/insurers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Initialize controller
var insurersController = insurers.Controller{}

func RenderIndex() *fyne.Container {
	insurerTableData := insurersController.GetInsurerTableData()

	table := widget.NewTable(
		func() (int, int) {
			return len(insurerTableData), len(insurerTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(insurerTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 100)

	newButton := widget.NewButton("Nouvel assureur", func() {
		fmt.Println("New button")
	})

	return container.NewBorder(newButton, nil, nil, nil, table)
}
