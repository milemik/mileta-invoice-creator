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

func TestBaseIdsListInit(t *testing.T) {
	comp := AllCompanies{}

	result := comp.BaseIdsList()
	if len(result) != 0 {
		t.Errorf("Expect that we have 0 base companies and we got %d", len(result))
	}
}

func TestBaseIdsList(t *testing.T) {
	comp := AllCompanies{Base: []Company{
		{Id: "1"},
		{Id: "2"},
		{Id: "3"},
	}}

	result := comp.BaseIdsList()
	if len(result) != 3 {
		t.Errorf("Expect that we have 3 base companies and we got %d", len(result))
	}

	if result[0] != "1" {
		t.Errorf("Expect first id in List to be 1 and we got %s", result[0])
	}
	if result[1] != "2" {
		t.Errorf("Expect first id in List to be 1 and we got %s", result[0])
	}
	if result[2] != "3" {
		t.Errorf("Expect first id in List to be 1 and we got %s", result[0])
	}
}

func TestTargetIdsListInit(t *testing.T) {
	comp := AllCompanies{}

	result := comp.TargetIdsList()
	if len(result) != 0 {
		t.Errorf("Expect that we have 0 base companies and we got %d", len(result))
	}
}

func TestTargetIdsList(t *testing.T) {
	comp := AllCompanies{All: []Company{
		{Id: "1"},
		{Id: "2"},
		{Id: "3"},
	}}

	result := comp.TargetIdsList()
	if len(result) != 3 {
		t.Errorf("Expect that we have 3 base companies and we got %d", len(result))
	}

	if result[0] != "1" {
		t.Errorf("Expect first id in List to be 1 and we got %s", result[0])
	}
	if result[1] != "2" {
		t.Errorf("Expect first id in List to be 1 and we got %s", result[0])
	}
	if result[2] != "3" {
		t.Errorf("Expect first id in List to be 1 and we got %s", result[0])
	}
}
