package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


func MainPage(app fyne.App) fyne.Window {
	window := app.NewWindow("MAIN")

	window.Resize(fyne.Size{Width: 600, Height: 600})

	welcomeText := canvas.NewText("Welcome to Invoice Creator", color.Black)
	welcomeText.TextStyle = fyne.TextStyle{Bold: true}
	createInvoiceBtn := widget.NewButton("Create invoice", func ()  {
		
	})
	addCompanyBtn := widget.NewButton("Add company", func ()  {
		
	})

	welcomeTextContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), welcomeText, layout.NewSpacer())

	buttonsContainer := container.New(layout.NewHBoxLayout(), createInvoiceBtn, addCompanyBtn)
	mainContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), welcomeTextContainer, buttonsContainer, layout.NewSpacer())
	window.SetContent(container.New(layout.NewCenterLayout(), mainContainer))

	return window
}