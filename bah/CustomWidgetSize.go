package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CustomSize() {
	a := app.New()
	w := a.NewWindow("Custom Size")
	w.Resize(fyne.NewSize(700, 500))
	size := fyne.NewSize(200, 100)

	input := widget.NewEntry()
	input.PlaceHolder = "Input Nama..."

	fixedInput := container.NewGridWrap(size, input)

	content := container.NewVBox(fixedInput)
	Tab := container.NewAppTabs(
		container.NewTabItem("Credentials store", content),
		container.NewTabItem("Commands", widget.NewLabel("Bot Menu")),
		container.NewTabItem("Status", widget.NewLabel("List Account")),
		container.NewTabItem("Logs", widget.NewLabel("Status")))

	Tab.SetTabLocation(container.TabLocationLeading)
	w.SetContent(Tab)
	w.Show()
	a.Run()
}
