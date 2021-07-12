package views

import (
	"fmt"
	"happy_bank_simulator/database"
	"happy_bank_simulator/initializers"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Overview(loansCounter int, borrowersCounter int, lendersCounter int, insurersCounter int) *fyne.Container {
	wipeDatabaseButton := widget.NewButton("Vider la base de données", func() {
		database.DropBD()
		fmt.Println("DP dropped")
		initializers.InitDB()
		fmt.Println("DP initialized")
	})

	populateDatabaseButton := widget.NewButton("Remplir la base", func() {
		fmt.Println("Populate database")
	})

	hbox := container.NewHBox(
		populateDatabaseButton,
		wipeDatabaseButton,
	)

	vbox := container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Nombre de crédits : %s", strconv.Itoa(loansCounter))),
		widget.NewLabel(fmt.Sprintf("Nombre de débiteurs : %s", strconv.Itoa(borrowersCounter))),
		widget.NewLabel(fmt.Sprintf("Nombre de créanciers : %s", strconv.Itoa(lendersCounter))),
		widget.NewLabel(fmt.Sprintf("Nombre d'assureurs : %s", strconv.Itoa(insurersCounter))),
	)

	return container.NewBorder(nil, hbox, nil, nil, vbox)
}
