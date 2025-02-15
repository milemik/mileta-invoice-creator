package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milemik/pdf-vezba/internal/database"
)

func CreateUI(myApp fyne.App, baseDir string) {
	fmt.Println("Opening UI")

	myWindow := myApp.NewWindow("Invoice generator")
	myWindow.Resize(fyne.NewSize(700, 700))

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

	exportButton := widget.NewButton("Export", func() {
		// Here we should save this information for later usage
		database.SaveToDB(database.Company{
			Id:              idInput.Text,
			OwnerName:       ownerNameInput.Text,
			CompanyName:     companyNameInput.Text,
			CompanyFullName: companyFullNameInput.Text,
			CompanyAddress:  companyAddressInput.Text,
			CompanyState:    companyStateInput.Text,
			CompanyEmail:    companyEmailInput.Text,
			Bank: database.BankAccount{
				SWIFT: bankSwiftNoInput.Text,
				IBAN:  bankIbanNoInput.Text,
			},
			PIB:           companyPibInput.Text,
			IsBaseCompany: isBaseCompanyInput.Checked,
		}, baseDir)
		log.Println("Added: ", idInput.Text)
		showPopUp(myApp, "ADDED: "+companyFullNameInput.Text)
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
		isBaseCompanyInput,

		layout.NewSpacer(),
		exportButton,
	)

	scrollableContent := container.NewScroll(content)

	myWindow.SetContent(scrollableContent)
	myWindow.Show()

}

func showPopUp(app fyne.App, content string) {
	// Very simple PopUp shown when we successfully add new company!
	addPopUpWindow := app.NewWindow("Company Added")
	addPopUpWindow.Resize(fyne.NewSize(100, 100))
	addPopUpWindow.SetContent(container.NewVBox(widget.NewLabel(content)))
	addPopUpWindow.Show()
}
