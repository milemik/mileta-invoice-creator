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

func TestGetTargetCompByIdInit(t *testing.T) {
	comp := AllCompanies{}

	_, err := comp.GetTargetCompById("1")

	if err == nil {
		t.Error("Expect to see error")
	}
}

func TestGetTargetCompById(t *testing.T) {
	comp := AllCompanies{All: []Company{
		{Id: "1"},
		{Id: "11"},
		{Id: "2"},
	}}

	res, err := comp.GetTargetCompById("1")

	if err != nil {
		t.Error("Error raised and we expect not to see error")
	}

	if res.Id != "1" {
		t.Errorf("We expect that company in result has id of 1 and we got %s", res.Id)
	}
}

func TestGetBaseCompByIdInit(t *testing.T) {
	comp := AllCompanies{}

	_, err := comp.GetBaseCompById("1")

	if err == nil {
		t.Error("Expect to see error")
	}
}

func TestGetBaseCompById(t *testing.T) {
	comp := AllCompanies{Base: []Company{
		{Id: "1"},
		{Id: "11"},
		{Id: "2"},
	}}

	res, err := comp.GetBaseCompById("1")

	if err != nil {
		t.Error("Error raised and we expect not to see error")
	}

	if res.Id != "1" {
		t.Errorf("We expect that company in result has id of 1 and we got %s", res.Id)
	}
}

func TestDeleteCompanyFromList(t *testing.T) {
	comp := AllCompanies{Base: []Company{
		{Id: "11", IsBaseCompany: true},
		{Id: "12", IsBaseCompany: true},
	}, All: []Company{
		{Id: "1", IsBaseCompany: false},
		{Id: "2", IsBaseCompany: false},
	}}

	compToDelete := comp.All[0]
	err := comp.DeleteFromList(compToDelete)
	if err != nil {
		t.Errorf("Got error %s", err)
	}

	if len(comp.All) != 1 {
		t.Errorf("Expected to have one company in All and got %d", len(comp.All))
	}

	for _, c := range comp.All {
		if c.Id == compToDelete.Id {
			t.Errorf("Expect that company with ID %s is deleted form list", compToDelete.Id)
		}
	}
	compToDelete = comp.Base[0]
	err = comp.DeleteFromList(compToDelete)
	if err != nil {
		t.Errorf("Got error that we didn't expect: %s", err)
	}
	for _, c := range comp.Base {
		if c.Id == compToDelete.Id {
			t.Errorf("Expected that company with ID %s is removed", compToDelete.Id)
		}
	}
}

func TestGetAllCompanies(t *testing.T) {
	all := AllCompanies{
		Base: []Company{
			{Id: "1"},
			{Id: "2"},
		},
		All: []Company{
			{Id: "11"},
			{Id: "12"},
		},
	}

	allComps := all.GetAllIds()
	if len(allComps) != 4 {
		t.Errorf("Expect to get 4 companies and got %d", len(allComps))
	}
}

func TestGetCompanyByIdFoo(t *testing.T) {
	all := AllCompanies{
		Base: []Company{
			{Id: "1"},
			{Id: "2"},
		},
		All: []Company{
			{Id: "11"},
			{Id: "12"},
		},
	}

	compToGet := all.Base[1]
	c, err := all.GetCompById(compToGet.Id)
	if err != nil {
		t.Error(err)
	}
	if c.Id != compToGet.Id {
		t.Errorf("Got company with ID %s and expected to get company with ID %s", c.Id, compToGet.Id)
	}

	compToGet = all.All[0]
	c, err = all.GetCompById(compToGet.Id)
	if err != nil {
		t.Error(err)
	}
	if c.Id != compToGet.Id {
		t.Errorf("Got company with ID %s and expected to get company with ID %s", c.Id, compToGet.Id)
	}
}
