package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateUI() {
	fmt.Println("Opening UI")

	myApp := app.New()
	myWindow := myApp.NewWindow("Invoice generator")

	ownerNameInput := widget.NewEntry()
	companyNameInput := widget.NewEntry()
	companyFullNameInput := widget.NewEntry()
	companyAddressInput := widget.NewEntry()
	companySityInput := widget.NewEntry()
	companyStateInput := widget.NewEntry()
	companyEmailInput := widget.NewEntry()
	companyPibInput := widget.NewEntry()

	bankSwiftNoInput := widget.NewEntry()
	bankIbanNoInput := widget.NewEntry()

	exportButton := widget.NewButton("Export", func() {
		// Here we should save this information for later usage
		fmt.Println(ownerNameInput.Text, companyNameInput.Text)
	})

	content := container.NewVBox(
		widget.NewLabel("OWNER NAME"),
		ownerNameInput,
		widget.NewLabel("COMPANY NAME"),
		companyNameInput,
		widget.NewLabel("COMPANY FULL NAME"),
		companyFullNameInput,
		widget.NewLabel("COMPANY ADDRESS"),
		companyAddressInput,
		widget.NewLabel("COMPANY SITY"),
		companySityInput,
		widget.NewLabel("COMPANY STATE"),
		companyStateInput,
		widget.NewLabel("COMPANY EMIL"),
		companyEmailInput,
		widget.NewLabel("COMPANY PIB/VAT NUMBER"),
		companyPibInput,
		widget.NewLabel("BANK SWIFT NUMBER"),
		bankSwiftNoInput,
		widget.NewLabel("BANK IBAN NUMBER"),
		bankIbanNoInput,

		layout.NewSpacer(),
		exportButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
