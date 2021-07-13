package views

import (
	"fmt"
	"happy_bank_simulator/app/lenders"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Initialize controller
var lendersController = lenders.Controller{}

func RenderIndex() *fyne.Container {
	lenderTableData := lendersController.GetLenderTableData()

	table := widget.NewTable(
		func() (int, int) {
			return len(lenderTableData), len(lenderTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(lenderTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 100)

	newButton := widget.NewButton("Nouveau créancier", func() {
		fmt.Println("New button")
	})

	return container.NewBorder(newButton, nil, nil, nil, table)
}
