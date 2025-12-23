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

	BotMenu := container.NewTabItem("Bot Menu", ActivityBot(
		func(name, schedule, action, remaintime string) error {
			return logic.BrowserIdle(name, schedule, action, remaintime)
		},
		func(name string) error {
			return logic.BrowserClose(name)
		}))

	//innerBotMenu := container.NewAppTabs(
	//	BotMenu,
	//	container.NewTabItem("Configuration", widget.NewLabel("Configuration")),
	//)

	innerTabsAccount := container.NewAppTabs(
		container.NewTabItem("List Account", NewAccountListTable()),
		container.NewTabItem("Configuration", widget.NewLabel("Configuration")),
	)
	listAccount := container.NewTabItem("List Account", innerTabsAccount)

	status := container.NewTabItem("Status", widget.NewLabel("Status"))

	innerLogTabs := container.NewAppTabs(
		container.NewTabItem("Logs", Logging()),
		container.NewTabItem("Configuration", widget.NewLabel("Configuration")),
	)
	Logs := container.NewTabItem("Logs", innerLogTabs)

	tabs := container.NewAppTabs(credTab, BotMenu, listAccount, status, Logs)
	tabs.SetTabLocation(container.TabLocationTop)
	return tabs
}
