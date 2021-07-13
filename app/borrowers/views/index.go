package views

import (
	"fmt"
	"happy_bank_simulator/app/borrowers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Initialize controller
var borrowersController = borrowers.Controller{}

func RenderIndex() *fyne.Container {
	borrowerTableData := borrowersController.GetBorrowerTableData()

	table := widget.NewTable(
		func() (int, int) {
			return len(borrowerTableData), len(borrowerTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(borrowerTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 100)

	newButton := widget.NewButton("Nouveau débiteur", func() {
		fmt.Println("New button")
	})

	return container.NewBorder(newButton, nil, nil, nil, table)
}
