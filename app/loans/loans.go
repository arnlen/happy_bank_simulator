package loans

import (
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/app/helpers"
	"happy_bank_simulator/models"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func RenderIndex() *fyne.Container {
	loansTableData := getLoanTableData()

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

	newButtonString := strings.Title(configs.Loan.String)
	newButton := widget.NewButtonWithIcon(newButtonString, theme.ContentAddIcon(), func() {
		fmt.Println("newButton", newButtonString)
		RenderNew()
	})

	return container.NewBorder(nil, newButton, nil, nil, table)
}

func RenderNew() {
	loanList := GetLoanStringList()

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

	dialog.ShowForm(strings.Title(configs.Loan.String), "Créer", "Annuler", formItems, func(bool) {
		fmt.Println("Nom :", nameEntry.Text)
		fmt.Println("Balance :", balanceEntry.Text)
	}, helpers.GetMasterWindow())
}

func getLoanTableData() [][]string {
	loans := models.ListLoans()

	loanTableData := [][]string{
		{"ID", "Débiteur", "Créancier", "Assureur", "Montant", "Durée"}}

	for _, loan := range loans {
		lenders := loan.Lenders
		lendersString := fmt.Sprintf("%s lenders", strconv.Itoa(len(lenders)))
		insurers := loan.Insurers
		insurersString := fmt.Sprintf("%s insurers", strconv.Itoa(len(insurers)))

		loanRow := []string{
			strconv.Itoa(int(loan.ID)),
			loan.Borrower.Name,
			lendersString,
			insurersString,
			fmt.Sprintf("%1.2f €", loan.Amount),
			fmt.Sprintf("%s mois", strconv.Itoa(int(loan.Duration))),
		}

		loanTableData = append(loanTableData, loanRow)
	}

	return loanTableData
}

func GetLoanStringList() []string {
	loans := models.ListLoans()
	var loanStringList []string

	for _, loan := range loans {
		string := fmt.Sprintf("%s - %1.2f € ", strconv.Itoa(int(loan.ID)), loan.Amount)
		loanStringList = append(loanStringList, string)
	}

	return loanStringList
}
