package utils

import "testing"


func TestDateContainerInit(t *testing.T) {
	expected := "01/02/0003"
	res := ValidateDateInput("1","2","3")

	if res != expected {
		t.Fatalf("Expected %q and got %q", expected, res)
	}
}


func TestDateContainerBad(t *testing.T) {
	expected := "01/02/0003"
	res := ValidateDateInput("100","2","3")

	if res != expected {
		t.Fatalf("Expected %q and got %q", expected, res)
	}
}