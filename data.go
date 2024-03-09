package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
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
	comp, err := all.GetBaseCompById(new.Id)
	if err != nil {
		log.Println(err)
		return all.Base
	}

	comp, err = all.GetTargetCompById(new.Id)
	if err != nil {
		log.Println(err)
		return all.All
	}

	if len(comp.Id) < 1 {
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

func GetDataFromDB() (AllCompanies, string, error) {
	dbFile := "all.json"
	var allCopmanies AllCompanies
	data, err := os.ReadFile(dbFile)

	err = json.Unmarshal(data, &allCopmanies)
	if err != nil {
		fmt.Println("Could not unmarshal data", err)
		return allCopmanies, dbFile, err
	}
	return allCopmanies, dbFile, nil
}

func SaveToDB(from Company) {
	allCopmanies, mainFile, err := GetDataFromDB()
	if err != nil {
		log.Println(err)
		return
	}

	_ = allCopmanies.AddCompany(from)

	jsonData, err := json.MarshalIndent(allCopmanies, "", " ")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(mainFile, jsonData, 0644)
	if err != nil {
		log.Println(err)
	}
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
