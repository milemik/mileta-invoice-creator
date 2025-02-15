package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/milemik/pdf-vezba/internal/database"
	"github.com/milemik/pdf-vezba/internal/helpers/db"
	"github.com/milemik/pdf-vezba/internal/pdf_creator"
	"github.com/milemik/pdf-vezba/internal/ui"
)

func main() {
	myApp := app.New()
	baseWindow := myApp.NewWindow("Invoice Creator")
	baseWindow.Resize(fyne.NewSize(700, 1000))
	userHomeDir := db.GetOutputDir()

	// TODO: Use binding to refresh list after adding new company!?
	companies, _, err := database.GetDataFromDB(userHomeDir)
	if err != nil {
		log.Println(err)
		// Maybe some popup saying error reading data!?
	}
	invoiceIdInput := widget.NewEntry()
	pricePerHourInput := widget.NewEntry()
	workedHoursInput := widget.NewEntry()

	baseCompSelect := widget.NewSelect(db.GetBaseCompanies(), func(s string) {
		log.Println("BASE SELECTED: " + s)
	})

	targetCompSelect := widget.NewSelect(db.GetTargetCompanies(), func(s string) {
		log.Println("TARGET SELECTED: " + s)
	})

	setupBtn := widget.NewButton("Setup", func() {
		ui.CreateUI(myApp, userHomeDir)
	})

	exportBtn := widget.NewButton("Create", func() {
		log.Println(baseCompSelect.Selected, targetCompSelect.Selected)

		baseComp, err := companies.GetBaseCompById(baseCompSelect.Selected)
		if err != nil {
			// We can show popup for error
			log.Println(err)
			return
		}
		targetComp, err := companies.GetTargetCompById(targetCompSelect.Selected)
		if err != nil {
			// We can show popup for error
			log.Println(err)
			return
		}

		fileName := invoiceIdInput.Text
		if len(fileName) < 1 {
			fileName = "test.pdf"
		}
		pdf_creator.CreatePDF(fileName, baseComp, targetComp, pricePerHourInput.Text, workedHoursInput.Text, userHomeDir)
	})

	// Select for delete
	selectForDelete := widget.NewSelect(db.DataGetAllIds(), func(s string) {
		log.Println("Selected for delete :", s)
	})

	delCompanyBtn := widget.NewButton("Delete company", func() {
		idToDelete := selectForDelete.Selected
		c, err := companies.GetCompById(idToDelete)
		if err != nil {
			log.Printf("ERROR : %s", err)
		}
		log.Printf("Deleting %s", c.Id)
		err = companies.DeleteFromList(c)
		if err != nil {
			log.Printf("ERROR: %s", err)
		}
		// Write updated data to DB
		database.WriteToDb(companies, database.GetDBLocation(userHomeDir))
	})

	refreshButton := widget.NewButton("Refresh", func() {
		log.Println("Refreshing")
		baseCompSelect.Options = db.GetBaseCompanies()
		targetCompSelect.Options = db.GetTargetCompanies()
		selectForDelete.Options = db.DataGetAllIds()
		baseCompSelect.Refresh()
		targetCompSelect.Refresh()
		selectForDelete.Refresh()
	})

	// Location info
	locInfo := widget.NewLabel("Invoices will be created in: " + userHomeDir)

	content := container.NewVBox(
		setupBtn,
		baseCompSelect,
		targetCompSelect,
		widget.NewLabel("INVOICE ID"),
		invoiceIdInput,
		widget.NewLabel("PRICE PER HOUR"),
		pricePerHourInput,
		widget.NewLabel("HOURS WORKED"),
		workedHoursInput,
		exportBtn,
		locInfo,
		selectForDelete,
		delCompanyBtn,
		refreshButton,
	)

	baseWindow.SetContent(content)
	baseWindow.ShowAndRun()
	os.Exit(1)
}
