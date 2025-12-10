package ui

import (
	"masterc/logic"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMainTabs() *container.AppTabs {
	credTab := container.NewTabItem("Credentials", CredentialsForm(func(name, user, pass, plat string) error {
		return logic.HandleCrendetials(name, user, pass, plat)
	}))

	BotMenu := container.NewTabItem("Bot Menu", widget.NewLabel("Bot Menu"))
	listAccount := container.NewTabItem("List Account", NewAccountListTable())
	status := container.NewTabItem("Status", widget.NewLabel("Status"))

	tabs := container.NewAppTabs(credTab, BotMenu, listAccount, status)
	tabs.SetTabLocation(container.TabLocationLeading)
	return tabs
}
