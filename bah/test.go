package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tester Menu")

	Menu1 := "Credentials Menu"
	Menu2 := "Bot Menu"
	Menu3 := "List Account"
	Menu4 := "Status"

	size := fyne.NewSize(300, 40)
	btnsize := fyne.NewSize(100, 40)

	input := widget.NewEntry()
	input.PlaceHolder = "Input Username.."
	input2 := widget.NewEntry()
	input2.PlaceHolder = "Input Password.."
	label1 := widget.NewLabel("")
	btn := widget.NewButton("Submit", func() { label1.SetText("Username: " + input.Text + " Password: " + input2.Text + "") })

	inputsize := container.NewGridWrap(size, input)
	input2size := container.NewGridWrap(size, input2)
	btnsizer := container.NewGridWrap(btnsize, btn)

	btnRow := container.NewHBox(
		layout.NewSpacer(),
		btnsizer,
		layout.NewSpacer(),
	)

	content := container.NewVBox(
		inputsize,
		input2size,
		btnRow,
		label1,
	)

	contentInBorder := container.NewBorder(
		nil, nil,
		layout.NewSpacer(),
		layout.NewSpacer(),
		container.NewCenter(content),
	)

	innerTab := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Tab 1 Content")),
		container.NewTabItem("Tab 2", widget.NewLabel("Tab 2 Content")),
		container.NewTabItem("Tab 3", widget.NewLabel("Tab 3 Content")),
	)

	Tab := container.NewAppTabs(
		container.NewTabItem(Menu1, contentInBorder),
		container.NewTabItem(Menu2, widget.NewLabel("Bot Menu")),
		container.NewTabItem(Menu3, widget.NewLabel("List Account")),
		container.NewTabItem(Menu4, widget.NewLabel("Status")),
		container.NewTabItem("List all", innerTab))

	Tab.SetTabLocation(container.TabLocationLeading)
	w.SetContent(Tab)
	w.Resize(fyne.NewSize(700, 500))
	w.Show()
	a.Run()
}
