package main

import (
	"os"

	"fyne.io/fyne/v2/app"

	"github.com/milemik/pdf-vezba/internal/ui"
)

func main() {
	myApp := app.New()
	window := ui.MainPage(myApp)
	window.ShowAndRun()
	os.Exit(1)
}
