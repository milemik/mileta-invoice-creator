package db

import (
	"log"
	"os"
	"path/filepath"

	"github.com/milemik/pdf-vezba/internal/database"
)

func GetBaseCompanies() []string {
	userHomeDir := GetOutputDir()
	companies, _, err := database.GetDataFromDB(userHomeDir)
	if err != nil {
		log.Println(err)
		// Maybe some popup saying error reading data!?
	}
	return companies.BaseIdsList()
}

func GetOutputDir() string {

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("ERROR: ", err)
	}
	outputPath := filepath.Join(userHomeDir, "MInvoceCreator")
	err = os.Mkdir(outputPath, 0755)
	if err != nil {
		log.Println("ERROR: ", err)
	}
	return outputPath
}

func GetTargetCompanies() []string {
	userHomeDir := GetOutputDir()
	companies, _, err := database.GetDataFromDB(userHomeDir)
	if err != nil {
		log.Println(err)
		// Maybe some popup saying error reading data!?
	}
	return companies.TargetIdsList()
}

func DataGetAllIds() []string {
	userHomeDir := GetOutputDir()
	companies, _, err := database.GetDataFromDB(userHomeDir)
	if err != nil {
		log.Println(err)
		// Maybe some popup saying error reading data!?
	}
	return companies.GetAllIds()
}
