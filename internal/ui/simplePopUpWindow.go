package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ShowPopUp(app fyne.App, title string, message string) {
	window := app.NewWindow(title)
	window.Resize(fyne.NewSize(300, 100))
	window.SetContent(widget.NewLabel(message))
	window.Show()
}
