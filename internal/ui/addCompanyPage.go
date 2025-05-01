package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/milemik/pdf-vezba/internal/database"
	"github.com/milemik/pdf-vezba/internal/helpers/db"
)

func AddCompany(app fyne.App) fyne.Window {
	window := app.NewWindow("ADD COMPANY")
	userHomeDir := db.GetOutputDir()

	idInput := widget.NewEntry()
	ownerNameInput := widget.NewEntry()
	companyNameInput := widget.NewEntry()
	companyFullNameInput := widget.NewEntry()
	companyAddressInput := widget.NewEntry()
	companySityInput := widget.NewEntry()
	companyStateInput := widget.NewEntry()
	companyEmailInput := widget.NewEntry()
	companyPibInput := widget.NewEntry()
	isBaseCompanyInput := widget.NewCheck("IS BASE COMPANY", func(b bool) {
		log.Println("isBaseSetTo: ", b)
	})

	bankSwiftNoInput := widget.NewEntry()
	bankIbanNoInput := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "ID for company", Widget: idInput},
			{Text: "Owner Name", Widget: ownerNameInput},
			{Text: "Company Name", Widget: companyNameInput},
			{Text: "Company Full Name", Widget: companyFullNameInput},
			{Text: "Address", Widget: companyAddressInput},
			{Text: "Company city", Widget: companySityInput},
			{Text: "Company state", Widget: companyStateInput},
			{Text: "Company email", Widget: companyEmailInput},
			{Text: "PIB", Widget: companyPibInput},
			{Text: "IS BASE COMPANY", Widget: isBaseCompanyInput},
			{Text: "SWIFT", Widget: bankSwiftNoInput},
			{Text: "IBAN", Widget: bankIbanNoInput},
		},
		OnSubmit: func() {
			// Here we should save this information for later usage
			database.SaveToDB(database.Company{
				Id:              idInput.Text,
				OwnerName:       ownerNameInput.Text,
				CompanyName:     companyNameInput.Text,
				CompanyFullName: companyFullNameInput.Text,
				CompanyAddress:  companyAddressInput.Text,
				CompanyCity:     companySityInput.Text,
				CompanyState:    companyStateInput.Text,
				CompanyEmail:    companyEmailInput.Text,
				Bank: database.BankAccount{
					SWIFT: bankSwiftNoInput.Text,
					IBAN:  bankIbanNoInput.Text,
				},
				PIB:           companyPibInput.Text,
				IsBaseCompany: isBaseCompanyInput.Checked,
			}, userHomeDir)
			log.Println("Added: ", idInput.Text)
			ShowPopUp(app, "Copmany added", "Added: "+companyFullNameInput.Text)
			window.Close()
		},
	}

	window.SetContent(form)
	return window
}
