package ui

import (
	"masterc/ui/strict"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ActivityBot(
	OnStart func(name string, schedule, action, remainTime string) error,
	OnStop func(name string) error) fyne.CanvasObject {

	inputName := widget.NewEntry()
	inputName.SetPlaceHolder("Input Nama...")
	inputTime := widget.NewSelect([]string{"5", "10", "20", "30", "Random"}, nil)
	inputTime.PlaceHolder = "Input Waktu... (Minutes)"
	inputAction := widget.NewSelect([]string{"idle", "story", "scroll"}, nil)
	inputAction.PlaceHolder = "Input Action..."
	inputRemain := strict.NewNumericalEntry()
	inputRemain.SetPlaceHolder("Input Waktu hidup bot..")

	label := widget.NewLabel("")
	label.Alignment = fyne.TextAlignCenter
	label.Wrapping = fyne.TextWrapWord

	btnOn := widget.NewButton("Submit", func() {
		if inputName.Text == "" {
			label.SetText("Nama tidak boleh kosong...")
			return
		}
		if inputAction.Selected != "idle" && inputRemain.Text == "" {
			label.SetText("Waktu hidup tidak boleh kosong...")
			return
		}

		name := inputName.Text
		Action := inputAction.Selected
		Schedule := inputTime.Selected
		RemainTime := inputRemain.Text

		label.SetText("Menjalankan " + Action + " untuk " + inputName.Text + "Setiap" + Schedule + " Menit" + "dengan nyawa bot:" + RemainTime + " Jam")

		err := OnStart(name, Schedule, Action, RemainTime)
		if err != nil {
			label.SetText("Gagal menjalankan bot: " + err.Error())
		} else {
			label.SetText("Berhasil menjalankan bot")
		}
		label.Refresh()
	})
	btnoff := widget.NewButton("Stop Bot", func() {
		err := OnStop(inputName.Text)
		if err != nil {
			label.SetText("Gagal menonaktifkan bot: " + err.Error())
		} else {
			label.SetText("Berhasil menonaktifkan bot")
		}
	})
	//btnRestart := widget.NewButton("Restart", func() {})

	fieldSize := fyne.NewSize(300, 40)
	nameWrap := container.NewGridWrap(fieldSize, inputName)
	scheduleWrap := container.NewGridWrap(fieldSize, inputTime)
	actionWrap := container.NewGridWrap(fieldSize, inputAction)
	RemainWrap := container.NewGridWrap(fieldSize, inputRemain)

	inputAction.OnChanged = func(s string) {
		if s == "idle" {
			scheduleWrap.Hide()
		} else {
			scheduleWrap.Show()
		}
	}

	form := container.NewVBox(
		nameWrap,
		actionWrap,
		scheduleWrap,
		RemainWrap,
		container.NewHBox(layout.NewSpacer(), btnOn, btnoff, layout.NewSpacer()),
		label)
	return container.NewCenter(form)
}
