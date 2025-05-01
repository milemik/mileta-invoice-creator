package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milemik/pdf-vezba/internal/helpers/db"
	"github.com/milemik/pdf-vezba/internal/helpers/utils"
)

func MainPage(app fyne.App) fyne.Window {
	window := app.NewWindow("MAIN")

	window.Resize(fyne.Size{Width: 600, Height: 600})

	welcomeText := canvas.NewText("Welcome to Invoice Creator", color.Black)
	welcomeText.TextStyle = fyne.TextStyle{Bold: true}
	createInvoiceBtn := widget.NewButton("Create invoice", func() {
		invoiceWindow := CreateInvoice(app)
		invoiceWindow.Show()
	})
	addCompanyBtn := widget.NewButton("Add company", func() {
		addCompanyWindow := AddCompany(app)
		addCompanyWindow.Show()
	})
	deleteCompanyBtn := widget.NewButton("Delete company", func() {
		deleteWindow := DeleteCompany(app)
		deleteWindow.Show()
	})

	openFileLocationBtn := widget.NewButton("Open invoces location", func() {
		err := utils.OpenFileLocation(db.GetOutputDir())
		if err != nil {
			ShowPopUp(app, "Error", "Failed to open file location: "+err.Error())
		}
	})

	welcomeTextContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), welcomeText, layout.NewSpacer())

	buttonsContainer := container.New(layout.NewHBoxLayout(), createInvoiceBtn, addCompanyBtn, deleteCompanyBtn)
	fileLocationContainer := container.New(layout.NewCenterLayout(), openFileLocationBtn)
	mainContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), welcomeTextContainer, buttonsContainer, fileLocationContainer, layout.NewSpacer())
	window.SetContent(container.New(layout.NewCenterLayout(), mainContainer))

	return window
}
