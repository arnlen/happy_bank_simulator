package helpers

import "fyne.io/fyne/v2"

var masterWindow fyne.Window

func GetMasterWindow() fyne.Window {
	return masterWindow
}

func SetMasterWindow(newWindow *fyne.Window) {
	masterWindow = *newWindow
}
