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

	// window := app.NewWindow("Hello")

	// hello := widget.NewLabel("Hello Fyne!")
	// window.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// ))

	// ---

	// myWindow := myApp.NewWindow("Canvas")
	// myCanvas := myWindow.Canvas()

	// green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	// text := canvas.NewText("Text", green)
	// text.TextStyle.Bold = true
	// myCanvas.SetContent(text)
	// go changeContent(myCanvas)

	// myWindow.Resize(fyne.NewSize(100, 100))

	// ---

	// myWindow := myApp.NewWindow("Box Layout")

	// text1 := canvas.NewText("Hello", color.White)
	// text2 := canvas.NewText("There", color.White)
	// text3 := canvas.NewText("(right)", color.White)
	// content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	// text4 := canvas.NewText("centered", color.White)
	// centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	// myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	// myWindow.ShowAndRun()

	// myApp := app.New()
	// myWindow := myApp.NewWindow("Grid Layout")

	// text1 := canvas.NewText("1", color.White)
	// text2 := canvas.NewText("2", color.White)
	// text3 := canvas.NewText("3", color.White)
	// grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	// myWindow.SetContent(grid)

	// ---

	// myWindow := myApp.NewWindow("Grid Wrap Layout")

	// text1 := canvas.NewText("1", color.White)
	// text2 := canvas.NewText("2", color.White)
	// text3 := canvas.NewText("3", color.White)
	// grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
	// 	text1, text2, text3)
	// myWindow.SetContent(grid)

	// myWindow.Resize(fyne.NewSize(180, 75))

	// ---

	// myWindow := myApp.NewWindow("Border Layout")

	// top := canvas.NewText("top bar", color.White)
	// left := canvas.NewText("left", color.White)
	// middle := canvas.NewText("content", color.White)
	// content := container.New(layout.NewBorderLayout(top, nil, left, nil),
	// 	top, left, middle)
	// myWindow.SetContent(content)

	// ---

	// myWindow := myApp.NewWindow("Form Layout")

	// label1 := canvas.NewText("Label 1", color.Black)
	// value1 := canvas.NewText("Value", color.White)
	// label2 := canvas.NewText("Label 2", color.Black)
	// value2 := canvas.NewText("Something", color.White)
	// grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2)
	// myWindow.SetContent(grid)

	// ---

	// myWindow := myApp.NewWindow("TabContainer Widget")

	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
	// 	container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	// 	container.NewTabItem("Pouet", widget.NewLabel("This is fuuuun")),
	// )

	// tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	// // tabs.SetTabLocation(container.TabLocationLeading)

	// myWindow.SetContent(tabs)

	// ---

	// myWindow := myApp.NewWindow("Entry Widget")

	// content := container.NewVBox(
	// 	widget.NewLabel("The top row of the VBox"),
	// 	container.NewHBox(
	// 		widget.NewLabel("Label 1"),
	// 		widget.NewLabel("Label 2"),
	// 	),
	// )

	// content.Add(widget.NewButton("Add more items", func() {
	// 	content.Add(widget.NewLabel("Added"))
	// }))

	// myWindow.SetContent(content)

	// ---

	// myWindow := myApp.NewWindow("Button Widget")

	// content := widget.NewButton("click me", func() {
	// 	log.Println("tapped")
	// })

	// //content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	// //	log.Println("tapped home")
	// //})

	// myWindow.SetContent(content)

	// ---

	// myWindow := myApp.NewWindow("Entry Widget")

	// input := widget.NewEntry()
	// input.SetPlaceHolder("Enter text...")
	// input2 := widget.NewEntry()
	// input2.SetPlaceHolder("Email")

	// content := container.NewVBox(input, input2, widget.NewButton("Save", func() {
	// 	log.Println("Content was:", input.Text)
	// }))

	// myWindow.SetContent(content)

	// ---

	// myWindow := myApp.NewWindow("ProgressBar Widget")

	// progress := widget.NewProgressBar()
	// infinite := widget.NewProgressBarInfinite()

	// go func() {
	// 	for i := 0.0; i <= 1.0; i += 0.01 {
	// 		time.Sleep(time.Millisecond * 250)
	// 		progress.SetValue(i)
	// 	}
	// }()

	// myWindow.SetContent(container.NewVBox(progress, infinite))

	// ---

	// myWindow := myApp.NewWindow("Toolbar Widget")

	// toolbar := widget.NewToolbar(
	// 	widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
	// 		log.Println("New document")
	// 	}),
	// 	widget.NewToolbarSeparator(),
	// 	widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	// 	widget.NewToolbarSpacer(),
	// 	widget.NewToolbarAction(theme.HelpIcon(), func() {
	// 		log.Println("Display help")
	// 	}),
	// )

	// content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	// myWindow.SetContent(content)

	// ---

	// var data = []string{"a", "string", "list"}

	// myWindow := myApp.NewWindow("List Widget")

	// list := widget.NewList(
	// 	func() int {
	// 		return len(data)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("template")
	// 	},
	// 	func(i widget.ListItemID, o fyne.CanvasObject) {
	// 		o.(*widget.Label).SetText(data[i])
	// 	})

	// myWindow.SetContent(list)

	// ---

	// myWindow := myApp.NewWindow("Simple")

	// str := binding.NewString()
	// str.Set("Initial value")

	// text := widget.NewLabelWithData(str)
	// myWindow.SetContent(text)

	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	str.Set("A new string")
	// }()

	// ---

	// myWindow := myApp.NewWindow("Two Way")

	// str := binding.NewString()
	// str.Set("Hi!")

	// myWindow.SetContent(container.NewVBox(
	// 	widget.NewLabelWithData(str),
	// 	widget.NewEntryWithData(str),
	// ))

	// ---

	myWindow := myApp.NewWindow("List Data")

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

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	myWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))

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
