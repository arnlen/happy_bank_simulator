package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func InitApp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Happy Bank Simulator")
	myWindow.Resize(fyne.NewSize(1024, 768))

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	nameEntry := widget.NewEntry()
	amountEntry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Nom", Widget: nameEntry},
			{Text: "Montant", Widget: amountEntry}},
		OnSubmit: func() { // optional, handle form submission
			fmt.Println("Form submitted:", nameEntry.Text, amountEntry.Text)
			val := fmt.Sprintf("%d - %s %s", data.Length()+1, nameEntry.Text, amountEntry.Text)
			data.Append(val)
		},
	}

	borrowersTabContent := container.NewBorder(form, nil, nil, nil, list)

	tabs := container.NewAppTabs(
		container.NewTabItem("Débiteurs", borrowersTabContent),
		container.NewTabItem("Créanciers", widget.NewLabel("Tableau des créanciers")),
		container.NewTabItem("Assureurs", widget.NewLabel("Tableau des assureurs")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
