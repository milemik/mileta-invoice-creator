package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-pdf/fpdf"

	"github.com/milemik/pdf-vezba/internal/database"
	"github.com/milemik/pdf-vezba/internal/helpers/db"
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
		createPDF(fileName, baseComp, targetComp, pricePerHourInput.Text, workedHoursInput.Text, userHomeDir)
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

func createPDF(filename string, baseComp, toComp database.Company, pricePerHour, hoursWorked, outputDir string) {
	// PDF CREATE
	pdf := fpdf.New("P", "mm", "A4", "")

	pricePerHourInt, err := strconv.ParseFloat(pricePerHour, 64)
	if err != nil {
		log.Println(err)
		return
	}

	hoursWorkedInt, err := strconv.ParseFloat(hoursWorked, 64)
	if err != nil {
		log.Println(err)
		return
	}

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	// facNum := "1/2024" // This should be from user input
	headerSetup(pdf, filename)
	drawLine(pdf, 10, 30, 200, 30)
	// Should be read from DB/JSON?

	fromToInfo(pdf, baseComp, toComp)
	drawLine(pdf, 10, 120, 200, 120)

	createTable(pdf, pricePerHourInt, hoursWorkedInt)
	drawLine(pdf, 10, 165, 200, 165)

	singDoc(pdf)
	footer(pdf)

	pdfPath := filepath.Join(outputDir, filename+".pdf")
	log.Println("PDF PATH: " + pdfPath)
	// Output to pdf
	err = pdf.OutputFileAndClose(pdfPath)
	if err != nil {
		fmt.Println(err)
	}
}

func drawLine(pdf *fpdf.Fpdf, x1, y1, x2, y2 float64) {
	// Settings for line
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(0.5)
	pdf.SetLineCapStyle("round")
	pdf.Line(x1, y1, x2, y2)
}

func headerSetup(pdf *fpdf.Fpdf, facNum string) {
	// Header
	now := time.Now().Format("01/02/2006")
	pdf.Cell(100, 10, "Invoice / Faktura: "+facNum)
	pdf.Cell(-1, 10, "Dated / Datum fakture: "+now)
	pdf.Cell(1, 20, "Value date / Datum prometa: "+now)
	pdf.Cell(1, 30, "Trading place / Mesto prometa: Mountain View")
}

func fromToInfo(pdf *fpdf.Fpdf, baseComp, to database.Company) {
	// FROM
	pdf.MoveTo(10, -10)
	pdf.Cell(100, 100, "From / Od:")
	pdf.Cell(50, 100, "Bill to / Komitent:")

	// Company name
	pdf.MoveTo(10, -10)
	pdf.Cell(100, 110, baseComp.CompanyFullName)
	pdf.Cell(50, 110, to.CompanyFullName)

	// Company Adress
	pdf.MoveTo(10, 0)
	pdf.Cell(100, 110, "Address / Adresa: "+baseComp.CompanyAddress)
	pdf.Cell(50, 110, "Address / Adresa: "+to.CompanyAddress)

	// Company Adress
	pdf.MoveTo(10, 10)
	pdf.Cell(100, 110, "City / Grad: "+baseComp.CompanyCity)
	pdf.Cell(50, 110, "City / Grad: "+to.CompanyCity)

	// Company Adress
	pdf.MoveTo(10, 20)
	pdf.Cell(100, 110, "Country / Država: "+baseComp.CompanyState)
	pdf.Cell(50, 110, "Country / Država: "+to.CompanyState)

	// VAT NO
	pdf.MoveTo(10, 30)
	pdf.Cell(100, 110, "Vat No/ PIB: "+baseComp.PIB)
	pdf.Cell(50, 110, "Vat No / Poreski broj: "+to.PIB)

	// EMAIL
	pdf.MoveTo(10, 40)
	pdf.Cell(100, 110, "E-mail: "+baseComp.CompanyEmail)
	pdf.Cell(100, 110, "E-mail: "+to.CompanyEmail)

	// SWIFT
	pdf.MoveTo(10, 50)
	pdf.Cell(100, 110, "SWIFT: "+baseComp.Bank.SWIFT)

	// BANK NUMBER
	pdf.MoveTo(10, 60)
	pdf.Cell(100, 110, "IBAN: "+baseComp.Bank.IBAN)
}

func createTable(pdf *fpdf.Fpdf, pricePerHour, hoursWorked float64) {
	// Invoice Info
	pdf.MoveTo(10, 70)
	// Table Header
	cellWidth := float64(40)
	pdf.Cell(cellWidth, 110, "ITEM")
	pdf.Cell(cellWidth, 110, "UNIT")
	pdf.Cell(cellWidth, 110, "QUATITY")
	pdf.Cell(cellWidth, 110, "PRICE")
	pdf.Cell(cellWidth, 110, "TOTAL")
	pdf.MoveTo(10, 76)
	pdf.Cell(cellWidth, 110, "(Vrsta usluge)")
	pdf.Cell(cellWidth, 110, "(Jedinica mere)")
	pdf.Cell(cellWidth, 110, "(KOLICINA)")
	pdf.Cell(cellWidth, 110, "(USD) (CENA)")
	pdf.Cell(cellWidth, 110, "(USD) (TOTAL)")
	// Table content
	pdf.MoveTo(10, 80)
	// Callculate amount
	summary := pricePerHour * hoursWorked

	pdf.Cell(cellWidth, 120, "Programerske")
	pdf.Cell(cellWidth, 125, "Hours/Sat")
	pdf.Cell(cellWidth, 125, fmt.Sprintf("%.0f", hoursWorked))
	pdf.Cell(cellWidth, 125, fmt.Sprintf("%.3f", pricePerHour))
	pdf.Cell(cellWidth, 125, fmt.Sprintf("%.2f", summary))

	pdf.MoveTo(10, 86)
	pdf.Cell(cellWidth, 120, "usluge (Programing)")

	pdf.MoveTo(10, 90)
	pdf.Cell(cellWidth*4, 130, "UKUPNO/SUM")
	pdf.Cell(cellWidth, 130, fmt.Sprintf("%.2f", summary))
}

func singDoc(pdf *fpdf.Fpdf) {

	pdf.MoveTo(20, 130)
	pdf.Cell(100, 100, "Usluge izvršio / services performed")
	pdf.Cell(100, 100, "Usluge primio / services received")

	drawLine(pdf, 25, 195, 85, 195)
	drawLine(pdf, 121, 195, 181, 195)

	pdf.MoveTo(30, 150)
	pdf.Cell(100, 100, "Ime i prezime / full name")
	pdf.Cell(100, 100, "Ime i prezime / full name")

	drawLine(pdf, 25, 215, 85, 215)
	drawLine(pdf, 121, 215, 181, 215)

	pdf.MoveTo(35, 170)
	pdf.Cell(100, 100, "Potpis/ signature")
	pdf.Cell(100, 100, "Potpis/ signature")

}

func footer(pdf *fpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(10, 240)
	pdf.Cell(0, 0, "NOTE / KOMENTAR:")

	pdf.SetFont("Arial", "I", 12)
	pdf.MoveTo(10, 245)
	pdf.Cell(0, 0, "Not in the VAT system / Poreski obaveznik nije u sistemu PDV-a.")

	pdf.MoveTo(10, 250)
	pdf.Cell(0, 0, "Complaints are received 10 days from receipt of invoice / Žalbe se primaju do 10 dana")

	pdf.MoveTo(10, 255)
	pdf.Cell(0, 0, "od prijema facture")

	pdf.MoveTo(10, 260)
	pdf.Cell(0, 0, "Payment within 30 days of receipt of the invoice / Valuta plaćanja je 30 dana od dana")

	pdf.MoveTo(10, 265)
	pdf.Cell(0, 0, "prijema facture.")
}
