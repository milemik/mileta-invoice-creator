package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/milemik/pdf-vezba/internal/database"
	"github.com/milemik/pdf-vezba/internal/helpers/db"
)

func DeleteCompany (app fyne.App) fyne.Window {
	window := app.NewWindow("DELETE COMPANY")
	userHomeDir := db.GetOutputDir()
	companies, _, err := database.GetDataFromDB(userHomeDir)
	if err != nil {
		log.Println(err)
		// Maybe some popup saying error reading data!?
	}

	// Select for delete
	selectForDelete := widget.NewSelect(db.DataGetAllIds(), func(s string) {
		log.Println("Selected for delete :", s)
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Select company to delete", Widget: selectForDelete},
		},
		OnSubmit: func() {
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
			window.Close()
		},
	}

	window.SetContent(form)
	return window
}