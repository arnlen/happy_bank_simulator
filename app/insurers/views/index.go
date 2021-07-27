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
	fmt.Printf("%s", insurerTableData[2])

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
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 150)

	return container.NewBorder(nil, nil, nil, nil, table)
}
