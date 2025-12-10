package ui

import (
	"masterc/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CredentialsForm(onSubmit func(name, user, pass, plat string) error) fyne.CanvasObject {

	inputName := widget.NewEntry()
	inputName.SetPlaceHolder("Input Nama (tanpa spasi)")
	inputUser := widget.NewEntry()
	inputUser.SetPlaceHolder("Input Username...")
	inputPass := widget.NewPasswordEntry()
	inputPass.SetPlaceHolder("Input Password...")
	inputPlatform := widget.NewSelect([]string{"Instagram", "Facebook", "Tiktok"}, nil)

	label := widget.NewLabel("")
	label.Alignment = fyne.TextAlignCenter
	label.Wrapping = fyne.TextWrapWord

	btn := widget.NewButton("Submit", func() {
		label.SetText("Menambahkan akun...")
		label.Refresh()
		name := inputName.Text
		user := inputUser.Text
		pass := inputPass.Text
		platform := inputPlatform.Selected

		if platform == "" {
			label.SetText("Pilih platform dulu....")
			return
		}

		if name == "" || user == "" || pass == "" {
			label.SetText("Jangan dibiarkan kosong fieldnya")
			return
		}
		err := logic.HandleCrendetials(name, user, pass, platform)
		if err != nil {
			label.SetText("Gagal menambahkan akun: " + err.Error())
		} else {
			label.SetText("Berhasil disimpan")
		}
		label.Refresh()
	})

	fieldSize := fyne.NewSize(300, 40)
	nameWrap := container.NewGridWrap(fieldSize, inputName)
	userWrap := container.NewGridWrap(fieldSize, inputUser)
	passWrap := container.NewGridWrap(fieldSize, inputPass)
	inputPlat := container.NewGridWrap(fieldSize, inputPlatform)

	form := container.NewVBox(
		nameWrap,
		userWrap,
		passWrap,
		inputPlat,
		container.NewHBox(layout.NewSpacer(), btn, layout.NewSpacer()),
		label,
	)

	return container.NewCenter(form)
}
