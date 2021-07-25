package views

import (
	"fmt"
	"happy_bank_simulator/app/helpers"
	"happy_bank_simulator/app/insurers"
	"happy_bank_simulator/app/loans"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
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

	refreshButton := widget.NewButton("Refraichir", func() {
		fmt.Println("Refresh not yet implemented!")
	})

	newButtonString := strings.Title(insurersController.GetModelName(false))
	newButton := widget.NewButtonWithIcon(newButtonString, theme.ContentAddIcon(), func() {
		fmt.Println("newButton", newButtonString)
		RenderNew()
	})

	return container.NewBorder(refreshButton, newButton, nil, nil, table)
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

	dialog.ShowForm(strings.Title(insurersController.GetModelName(false)), "Créer", "Annuler", formItems, func(bool) {
		fmt.Println("Nom :", nameEntry.Text)
		fmt.Println("Balance :", balanceEntry.Text)
	}, helpers.GetMasterWindow())
}
