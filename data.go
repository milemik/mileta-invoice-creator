package main

import (
	"encoding/json"
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
	if new.IsBaseCompany {
		all.Base = append(all.Base, new)
		return all.Base
	} else {
		all.All = append(all.All, new)
		return all.All
	}
}

func GetDataFromDB() (AllCompanies, string) {
	dbFile := "all.json"
	var allCopmanies AllCompanies
	data, err := os.ReadFile(dbFile)

	err = json.Unmarshal(data, &allCopmanies)
	if err != nil {
		fmt.Println("Could not unmarshal data", err)
		os.Exit(1)
	}
	return allCopmanies, dbFile
}

func SaveToDB(from Company) {
	allCopmanies, mainFile := GetDataFromDB()

	_ = allCopmanies.AddCompany(from)

	jsonData, err := json.MarshalIndent(allCopmanies, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = os.WriteFile(mainFile, jsonData, 0644)
	if err != nil {
		log.Println(err)
	}
	log.Println("File saved in: " + mainFile)
}
