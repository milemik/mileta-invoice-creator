package utils

import (
	"log"
	"strconv"
	"time"
)


func ValidateDateInput(day, month, year string) string {
	dayInt, err := strconv.Atoi(day)
	if err != nil {
		log.Fatalf("Could not convert %q to int", year)
	}
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		log.Fatalf("Could not convert %q to int", month)
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		log.Fatalf("Could not convert %q to int", day)
	}
	return time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC).Format("02/01/2006")
}