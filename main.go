package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// Gorm init
	db, err := gorm.Open(sqlite.Open("happy_dev.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&Loan{},
		&Borrower{},
	)

	myApp := app.New()
	myWindow := myApp.NewWindow("Happy Bank Simulator")
	myWindow.Resize(fyne.NewSize(1024, 768))

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	// template := container.NewHBox(
	// 	widget.NewLabel("template"),
	// )Ÿ

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

	// Create test loan
	testLoan := Loan{
		Model:            gorm.Model{},
		StartDate:        "27/06/2021",
		EndDate:          "27/06/2022",
		Duration:         12,
		Amount:           1000,
		InitialDeposit:   100,
		MonthlyCredit:    1.2,
		MonthlyInsurance: 12.4,
	}

	fmt.Printf("%+v\n", testLoan)

	dateString := "01/02/2021"
	date, _ := time.Parse("02/01/2006", dateString)
	fmt.Println(date.Format("2 Jan. 2006"))
}

type Borrower struct {
	gorm.Model
	Loans []Loan
}

type Loan struct {
	gorm.Model
	BorrowerID       uint
	StartDate        string
	EndDate          string
	Duration         int32
	Amount           float32
	InitialDeposit   float32
	MonthlyCredit    float32
	MonthlyInsurance float32
}

type MonthlyPayment struct {
	gorm.Model
	Loan     Loan
	Borrower Borrower
	Amount   float32
}

func showAnother(a fyne.App) {
	time.Sleep(time.Second * 5)

	win := a.NewWindow("Shown later")
	win.SetContent(widget.NewLabel("5 seconds later"))
	win.Resize(fyne.NewSize(200, 200))
	win.Show()

	time.Sleep(time.Second * 2)
	win.Close()
}

func changeContent(c fyne.Canvas) {
	time.Sleep(time.Second * 2)

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	c.SetContent(canvas.NewRectangle(blue))

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewLine(color.Gray{Y: 180}))

	time.Sleep(time.Second * 2)
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewImageFromResource(theme.FyneLogo()))
}
