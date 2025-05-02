package ui

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milemik/pdf-vezba/internal/database"
	"github.com/milemik/pdf-vezba/internal/helpers/db"
	"github.com/milemik/pdf-vezba/internal/helpers/utils"
	"github.com/milemik/pdf-vezba/internal/pdf_creator"
)

func CreateInvoice(app fyne.App) fyne.Window {
	window := app.NewWindow("Create Invoice")
	userHomeDir := db.GetOutputDir()
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
	dateDayText := canvas.NewText("Day", color.Black)
	dateMonthText := canvas.NewText("Month", color.Black)
	dateYearText := canvas.NewText("Year", color.Black)
	dayInMonthInp := widget.NewEntry()
	monthInYearInp := widget.NewEntry()
	yearInp := widget.NewEntry()

	dateContainer := container.New(layout.NewHBoxLayout(), dateDayText, dayInMonthInp, dateMonthText, monthInYearInp, dateYearText, yearInp)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Invoice ID", Widget: invoiceIdInput},
			{Text: "Price per day", Widget: pricePerHourInput},
			{Text: "Days worked", Widget: workedHoursInput},
			{Text: "Select base company", Widget: baseCompSelect},
			{Text: "Select target company", Widget: targetCompSelect},
			{Text: "Select date", Widget: dateContainer},
		},
		OnSubmit: func() {
			log.Println("Form submited")
			log.Println(baseCompSelect.Selected, targetCompSelect.Selected)


			strDate := utils.ValidateDateInput(dayInMonthInp.Text, monthInYearInp.Text, yearInp.Text)

			log.Println("DATE:"+strDate)

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
			pdf_creator.CreatePDF(fileName, baseComp, targetComp, pricePerHourInput.Text, workedHoursInput.Text, userHomeDir, strDate)
			ShowPopUp(app, "Success", "Invoice created successfully")
			window.Close()
		},
	}

	window.SetContent(form)

	return window
}
