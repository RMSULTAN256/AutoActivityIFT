package ui

import (
	"encoding/json"
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

type UserShow struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Platform string `json:"platform"`
}

func NewAccountListTable() fyne.CanvasObject {
	var users []UserShow

	table := widget.NewTable(
		func() (int, int) {
			return len(users) + 1, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)

			if id.Row == 0 {
				switch id.Col {
				case 0:
					label.SetText("Nama")
				case 1:
					label.SetText("Username")
				case 2:
					label.SetText("Platform")
				}
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Alignment = fyne.TextAlignCenter
				return
			}

			u := users[id.Row-1]
			label.TextStyle = fyne.TextStyle{}

			switch id.Col {
			case 0:
				label.SetText(u.Name)

			case 1:
				label.SetText(u.Username)

			case 2:
				label.SetText(u.Platform)
			}
		},
	)

	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 300)
	table.SetColumnWidth(2, 150)

	scroll := container.NewScroll(table)

	go func() {
		u := url.URL{
			Scheme: "ws",
			Host:   "127.0.0.1:5544",
			Path:   "/api/v1/ws",
		}
		log.Println("connecting to", u.String())

		conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Println("ws dial error:", err)
			return
		}
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("ws read error:", err)
				return
			}

			var list []UserShow
			if err := json.Unmarshal(msg, &list); err != nil {
				log.Println("Bad Payload:", err, "raw:", string(msg))
				continue
			}

			fyne.Do(func() {
				users = list
				table.Refresh()
			})
		}
	}()

	return scroll
}
