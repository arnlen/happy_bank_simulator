package services

import "fyne.io/fyne/v2"

var appWindow fyne.Window

func GetAppWindow() fyne.Window {
	return appWindow
}

func SetAppWindow(window fyne.Window) {
	appWindow = window
}
