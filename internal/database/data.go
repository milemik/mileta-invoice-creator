package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type AllCompanies struct {
	Base []Company `json:"base"`
	All  []Company `json:"all"`
}

// Company Information
type Company struct {
	Id              string      `json:"id"`
	OwnerName       string      `json:"ownerName"`
	CompanyName     string      `json:"companyName"`
	CompanyFullName string      `json:"companyFullName"`
	CompanyAddress  string      `json:"companyAddress"`
	CompanyCity     string      `json:"companyCity"`
	CompanyState    string      `json:"companyState"`
	CompanyEmail    string      `json:"companyEmail"`
	Bank            BankAccount `json:"bank"`
	PIB             string      `json:"pib"`
	IsBaseCompany   bool        `json:"isBaseCompany"`
}

// Bank account
type BankAccount struct {
	SWIFT string `json:"swift"`
	IBAN  string `json:"iban"`
}

func (all *AllCompanies) AddCompany(new Company) []Company {
	comp, _ := all.GetBaseCompById(new.Id)

	comp, _ = all.GetTargetCompById(new.Id)

	if len(comp.Id) > 0 {
		log.Println("Company already exists in DB")
		return all.All
	}

	if new.IsBaseCompany {
		all.Base = append(all.Base, new)
		return all.Base
	} else {
		all.All = append(all.All, new)
		return all.All
	}
}

func (all *AllCompanies) BaseIdsList() []string {
	var ids []string
	for _, company := range all.Base {
		ids = append(ids, company.Id)
	}
	return ids
}

func (all *AllCompanies) TargetIdsList() []string {
	var ids []string
	for _, company := range all.All {
		ids = append(ids, company.Id)
	}
	return ids
}

func (all *AllCompanies) DeleteFromList(company Company) error {
	if company.IsBaseCompany {
		companyIndex, err := getCompanyIndex(company, all.Base)
		if err != nil {
			return err
		}
		all.Base = append(all.Base[:companyIndex], all.Base[companyIndex+1:]...)
	} else {
		companyIndex, err := getCompanyIndex(company, all.All)
		if err != nil {
			return err
		}
		all.All = append(all.All[:companyIndex], all.All[companyIndex+1:]...)
	}
	return nil
}

func getCompanyIndex(company Company, compList []Company) (int, error) {
	companyIndex := -1
	for index, comp := range compList {
		if comp == company {
			companyIndex = index
			break
		}
	}
	if companyIndex < 0 {
		return companyIndex, errors.New("Company doesn't exist")
	}
	return companyIndex, nil
}

func GetDBLocation(baseDir string) string {
	return filepath.Join(baseDir, "all.json")
}

// GetDataFromDB get data from database if file don't exist it will create new one
func GetDataFromDB(baseDir string) (AllCompanies, string, error) {
	dbFile := GetDBLocation(baseDir)
	var allCopmanies AllCompanies
	data, err := os.ReadFile(dbFile)
	if err != nil {
		log.Println("ERROR: ", err)
		errMsg := err.Error()
		if strings.Contains(errMsg, "all.json:") {
			initalDBSetup(dbFile, allCopmanies)
		}
		return allCopmanies, dbFile, err
	}

	err = json.Unmarshal(data, &allCopmanies)
	if err != nil {
		fmt.Println("Could not unmarshal data", err)
		return allCopmanies, dbFile, err
	}
	return allCopmanies, dbFile, nil
}

func initalDBSetup(dbFile string, allCopmanies AllCompanies) {
	_, err := os.Create(dbFile)
	if err != nil {
		log.Println(err)
		return
	}
	WriteToDb(allCopmanies, dbFile)
}

func WriteToDb(allCopmanies AllCompanies, dbFile string) {
	jsonData, err := json.MarshalIndent(allCopmanies, "", " ")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(dbFile, jsonData, 0644)
	if err != nil {
		log.Println(err)
	}
}

func SaveToDB(from Company, baseDir string) {
	allCopmanies, mainFile, err := GetDataFromDB(baseDir)
	if err != nil {
		log.Println(err)
		return
	}

	_ = allCopmanies.AddCompany(from)

	WriteToDb(allCopmanies, mainFile)
	log.Println("File saved in: " + mainFile)
}

func (all *AllCompanies) GetTargetCompById(id string) (Company, error) {
	for _, company := range all.All {
		if company.Id == id {
			return company, nil
		}
	}
	return Company{}, errors.New("TARGET: Could not found company with ID: " + id)
}

func (all *AllCompanies) GetBaseCompById(id string) (Company, error) {
	for _, company := range all.Base {
		if company.Id == id {
			return company, nil
		}
	}
	return Company{}, errors.New("BASE: Could not found company with ID: " + id)
}

func (all *AllCompanies) GetCompById(id string) (Company, error) {
	comp, err := all.GetBaseCompById(id)
	if err != nil {
		comp, err = all.GetTargetCompById(id)
		if err != nil {
			return Company{}, errors.New("Could not find company with ID: " + id)
		}
	}
	return comp, nil
}

func (all *AllCompanies) GetAllIds() []string {
	return append(all.BaseIdsList(), all.TargetIdsList()...)
}
