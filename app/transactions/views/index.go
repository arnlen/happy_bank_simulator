package views

import (
	"fmt"
	"happy_bank_simulator/app/transactions"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Initialize controller
var transactionsController = transactions.Controller{}

func RenderIndex() *fyne.Container {
	transactionTableData := transactionsController.GetTransactionTableData()

	table := widget.NewTable(
		func() (int, int) {
			return len(transactionTableData), len(transactionTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(transactionTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 150)

	newButtonString := strings.Title(transactionsController.GetModelName(false))
	newButton := widget.NewButtonWithIcon(newButtonString, theme.ContentAddIcon(), func() {
		fmt.Println("newButton", newButtonString)
		RenderNew()
	})

	return container.NewBorder(nil, newButton, nil, nil, table)
}

func RenderNew() {
}
