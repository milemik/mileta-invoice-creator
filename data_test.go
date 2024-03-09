package main

import "testing"

func TestAddCompany(t *testing.T) {
	comp := AllCompanies{}

	compToAdd := Company{Id: "TestId"}

	comp.AddCompany(compToAdd)

	if len(comp.All) < 1 {
		t.Errorf("Expected to have 1 company and got %d", len(comp.All))
	}
}
