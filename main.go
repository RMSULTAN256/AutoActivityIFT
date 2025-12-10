package main

import (
	"masterc/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Bot Menu")
	w.Resize(fyne.NewSize(1000, 600))
	w.SetContent(ui.NewMainTabs())
	w.ShowAndRun()
}
