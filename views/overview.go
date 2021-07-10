package views

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Overview(loansCounter int, borrowersCounter int, lendersCounter int, insurersCounter int) *fyne.Container {
	wipeDatabaseButton := widget.NewButton("Vider la base de données", func() {
		fmt.Println("Drop database")
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
