package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-pdf/fpdf"
)

func main() {
	myApp := app.New()
	baseWindow := myApp.NewWindow("Invoice Creator")
	baseWindow.Resize(fyne.NewSize(700, 1000))

	setupBtn := widget.NewButton("Setup", func() {
		CreateUI(myApp)
	})

	exportBtn := widget.NewButton("Create", func() {

		fileName := "example.pdf"
		baseComp := Company{
			Id:              "BASE UI Test",
			OwnerName:       "FirstName LastName",
			CompanyName:     "Base Company",
			CompanyFullName: "UI Base Company LTD.",
			CompanyAddress:  "Samle Address 123",
			CompanyCity:     "City",
			CompanyState:    "State",
			CompanyEmail:    "email@base.com",
			Bank: BankAccount{
				SWIFT: "DS122434345435345345",
				IBAN:  "DWADWADADEDFE",
			},
			PIB: "12324234234",
		}
		// Should be read from DB/JSON?
		toComp := Company{
			Id:              "ToComp",
			OwnerName:       "FirstNameTo LastNameTo",
			CompanyName:     "To Company",
			CompanyFullName: "To Company LTD.",
			CompanyAddress:  "To Address 123",
			CompanyCity:     "ToCity",
			CompanyState:    "ToState",
			CompanyEmail:    "email@to.com",
			PIB:             "45455556565",
		}
		createPDF(fileName, baseComp, toComp)
	})

	content := container.NewVBox(
		setupBtn,
		exportBtn,
	)

	baseWindow.SetContent(content)
	baseWindow.ShowAndRun()
	os.Exit(1)
}

func createPDF(filename string, baseComp, toComp Company) {
	// PDF CREATE
	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	facNum := "1/2024" // This should be from user input
	headerSetup(pdf, facNum)
	drawLine(pdf, 10, 30, 200, 30)
	// Should be read from DB/JSON?

	fromToInfo(pdf, baseComp, toComp)
	drawLine(pdf, 10, 120, 200, 120)
	createTable(pdf)

	singDoc(pdf)
	footer(pdf)

	// Output to pdf
	err := pdf.OutputFileAndClose(filename)
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

func fromToInfo(pdf *fpdf.Fpdf, baseComp, to Company) {
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

func createTable(pdf *fpdf.Fpdf) {
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
	// Calculate amount
	pricePerHour := 10
	hoursWorked := 50
	summary := pricePerHour * hoursWorked

	pdf.Cell(cellWidth, 120, "Programerske")
	pdf.Cell(cellWidth, 125, "Hours/Sat")
	pdf.Cell(cellWidth, 125, strconv.Itoa(pricePerHour))
	pdf.Cell(cellWidth, 125, strconv.Itoa(hoursWorked))
	pdf.Cell(cellWidth, 125, strconv.Itoa(summary))

	pdf.MoveTo(10, 86)
	pdf.Cell(cellWidth, 120, "usluge (Programing)")

	pdf.MoveTo(10, 90)
	pdf.Cell(cellWidth*4, 130, "UKUPNO/SUM")
	pdf.Cell(cellWidth, 130, strconv.Itoa(summary))

}

func singDoc(pdf *fpdf.Fpdf) {

	pdf.MoveTo(20, 120)
	pdf.Cell(100, 100, "Usluge izvršio / services performed")
	pdf.Cell(100, 100, "Usluge primio / services received")

	drawLine(pdf, 25, 185, 85, 185)
	drawLine(pdf, 121, 185, 181, 185)

	pdf.MoveTo(30, 140)
	pdf.Cell(100, 100, "Ime i prezime / full name")
	pdf.Cell(100, 100, "Ime i prezime / full name")

	drawLine(pdf, 25, 205, 85, 205)
	drawLine(pdf, 121, 205, 181, 205)

	pdf.MoveTo(35, 160)
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
