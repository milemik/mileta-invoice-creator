package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-pdf/fpdf"
)

// Company Information
type From struct {
	OwnerName       string      `json:"ownerName"`
	CompanyName     string      `json:"companyName"`
	CompanyFullName string      `json:"companyFullName"`
	CompanyAddress  string      `json:"companyAddress"`
	CompanyCity     string      `json:"companyCity"`
	CompanyState    string      `json:"companyState"`
	CompanyEmail    string      `json:"companyEmail"`
	Bank            BankAccount `json:"bank"`
	PIB             string      `json:"pib"`
}

// Bank account
type BankAccount struct {
	SWIFT string `json:"swift"`
	IBAN  string `json:"iban"`
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide needed arguments\n\tcreate: create pdf\n\tsetup: configure program")
		os.Exit(1)
	}
	if args[1] == "setup" {
		CreateUI()
		// Add logic to create config.json file - or maybe write data to some DB
		fmt.Println("Starting setup")
		return
	}
	if args[1] != "create" {
		fmt.Println("Exiting")
		return
	}

	if len(args) < 3 {
		fmt.Println("Please provice invoice name!")
		return
	}

	fileName := args[2] + ".pdf"
	fmt.Println(args)

	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	headerSetup(pdf)
	drawLine(pdf, 10, 30, 200, 30)
	fromToInfo(pdf)
	drawLine(pdf, 10, 120, 200, 120)
	createTable(pdf)

	singDoc(pdf)
	footer(pdf)

	// Output to pdf
	err := pdf.OutputFileAndClose(fileName)
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

func headerSetup(pdf *fpdf.Fpdf) {
	// Header
	pdf.Cell(100, 10, "Invoice / Faktura: <BROJ FAKTURE>")
	pdf.Cell(-1, 10, "Dated / Datum fakture: <DATUM>")
	pdf.Cell(1, 20, "Value date / Datum prometa: <DATUM>")
	pdf.Cell(1, 30, "Trading place / Mesto prometa: Mountain View")
}

func fromToInfo(pdf *fpdf.Fpdf) {
	// FROM
	pdf.MoveTo(10, -10)
	pdf.Cell(100, 100, "From / Od:")
	pdf.Cell(50, 100, "Bill to / Komitent:")

	// Company name
	pdf.MoveTo(10, -10)
	pdf.Cell(100, 110, "From copany name")
	pdf.Cell(50, 110, "To company name")

	// Company Adress
	pdf.MoveTo(10, 0)
	pdf.Cell(100, 110, "Address / Adresa: Blabla Street from 55")
	pdf.Cell(50, 110, "Address / Adresa: Blabla Street 55")

	// Company Adress
	pdf.MoveTo(10, 10)
	pdf.Cell(100, 110, "City / Grad: <GradFrom>")
	pdf.Cell(50, 110, "City / Grad: <GradTo>")

	// Company Adress
	pdf.MoveTo(10, 20)
	pdf.Cell(100, 110, "Country / Država: <DrzavaFrom>")
	pdf.Cell(50, 110, "Country / Država: <DrzavaTo>")

	// VAT NO
	pdf.MoveTo(10, 30)
	pdf.Cell(100, 110, "Vat No/ PIB: 123456789")
	pdf.Cell(50, 110, "Vat No / Poreski broj: 12345678TT")

	// EMAIL
	pdf.MoveTo(10, 40)
	pdf.Cell(100, 110, "E-mail: <company123@company.com")
	pdf.Cell(100, 110, "E-mail: <company123To@company.com")

	// SWIFT
	pdf.MoveTo(10, 50)
	pdf.Cell(100, 110, "SWIFT: SWIFTHERE")

	// BANK NUMBER
	pdf.MoveTo(10, 60)
	pdf.Cell(100, 110, "IBAN:  BANKNUM123456789")
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
