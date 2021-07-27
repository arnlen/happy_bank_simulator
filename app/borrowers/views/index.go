package views

import (
	"fmt"
	"happy_bank_simulator/app/borrowers"
	"happy_bank_simulator/app/helpers"
	"happy_bank_simulator/app/loans"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Initialize controller
var borrowersController = borrowers.Controller{}

func RenderIndex() *fyne.Container {
	borrowerTableData := updateTableData()

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
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 150)

	newButtonString := strings.Title(borrowersController.GetModelName(false))
	newButton := widget.NewButtonWithIcon(newButtonString, theme.ContentAddIcon(), func() {
		fmt.Println("newButton", newButtonString)
		RenderNew()
	})

	return container.NewBorder(nil, newButton, nil, nil, table)
}

func updateTableData() [][]string {
	return borrowersController.GetBorrowerTableData()
}

func RenderNew() {
	loansController := loans.Controller{}
	loanList := loansController.GetLoanStringList()

	nameEntry := widget.NewEntry()
	balanceEntry := widget.NewEntry()
	selectMenu := widget.Select{
		Options: loanList,
	}

	formItems := []*widget.FormItem{
		{Text: "Nom", Widget: nameEntry},
		{Text: "Balance", Widget: balanceEntry},
		{Text: "Prêt associé", Widget: &selectMenu},
	}

	dialog.ShowForm(strings.Title(borrowersController.GetModelName(false)), "Créer", "Annuler", formItems, func(bool) {
		fmt.Println("Nom :", nameEntry.Text)
		fmt.Println("Balance :", balanceEntry.Text)
	}, helpers.GetMasterWindow())
}
