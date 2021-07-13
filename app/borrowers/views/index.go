package views

import (
	"fmt"
	"happy_bank_simulator/app/borrowers"
	appHelpers "happy_bank_simulator/app/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
		RenderNew()
	})

	return container.NewBorder(newButton, nil, nil, nil, table)
}

func RenderNew() {
	loanList := borrowersController.GetLoanStringList()

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

	dialog.ShowForm("Nouveau débiteur", "Créer", "Annuler", formItems, func(bool) {
		fmt.Println("Nom :", nameEntry.Text)
		fmt.Println("Balance :", balanceEntry.Text)
	}, appHelpers.GetMasterWindow())
}
