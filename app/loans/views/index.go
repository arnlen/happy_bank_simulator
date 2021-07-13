package views

import (
	"fmt"
	"happy_bank_simulator/app/loans"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Initialize controller
var loansController = loans.Controller{}

func RenderIndex() *fyne.Container {
	loansTableData := loansController.GetLoanTableData()

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

	newButton := widget.NewButton("Nouveau prêt", func() {
		fmt.Println("New button")
	})

	return container.NewBorder(newButton, nil, nil, nil, table)
}
