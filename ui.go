package main

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateUI() {
	fmt.Println("Opening UI")

	myApp := app.New()
	myWindow := myApp.NewWindow("Invoice generator")

	idInput := widget.NewEntry()
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
		saveToJson(Company{
			Id:              idInput.Text,
			OwnerName:       ownerNameInput.Text,
			CompanyName:     companyNameInput.Text,
			CompanyFullName: companyFullNameInput.Text,
			CompanyAddress:  companyAddressInput.Text,
			CompanyState:    companyStateInput.Text,
			CompanyEmail:    companyEmailInput.Text,
			Bank: BankAccount{
				SWIFT: bankSwiftNoInput.Text,
				IBAN:  bankIbanNoInput.Text,
			},
			PIB: companyPibInput.Text,
		})
		fmt.Println(ownerNameInput.Text, companyNameInput.Text)
	})

	content := container.NewVBox(
		widget.NewLabel("CREATE ID - UNIQUE NAME FOR COMPANY"),
		idInput,
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

func saveToJson(from Company) {
	jsonData, err := json.MarshalIndent(from, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filename := from.CompanyName + ".json"
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File saved in: " + filename)
}
