package main

import "testing"

func TestAddCompany(t *testing.T) {
	comp := AllCompanies{}

	compToAdd := Company{Id: "TestId"}
	comp.AddCompany(compToAdd)

	if len(comp.All) != 1 {
		t.Errorf("Expected to have 1 company and got %d", len(comp.All))
	}
	// Try to add same company again
	comp.AddCompany(compToAdd)

	if len(comp.All) != 1 {
		t.Errorf("Expected to have 1 company and got %d", len(comp.All))
	}

	if len(comp.Base) != 0 {
		t.Errorf("Number of base companies should be 0 and we got %d", len(comp.Base))
	}
}

func TestAddBaseCompany(t *testing.T) {
	comp := AllCompanies{}

	compToAdd := Company{Id: "TestTest", IsBaseCompany: true}

	comp.AddCompany(compToAdd)

	if len(comp.All) != 0 {
		t.Errorf("Expect that number of target companies is 0 and we got %d", len(comp.All))
	}

	if len(comp.Base) != 1 {
		t.Errorf("Expect that number of Base companies is 1 and we got %d", len(comp.Base))
	}

}
