package app

import "fyne.io/fyne/v2"

var (
	app          fyne.App
	masterWindow fyne.Window
)

func SetApp(newApp *fyne.App) {
	app = *newApp
}

func GetApp() fyne.App {
	return app
}

func SetMasterWindow(newWindow *fyne.Window) {
	masterWindow = *newWindow
}

func GetMasterWindow() fyne.Window {
	return masterWindow
}
