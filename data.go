package main

// Company Information
type From struct {
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
}

// Bank account
type BankAccount struct {
	SWIFT string `json:"swift"`
	IBAN  string `json:"iban"`
}
