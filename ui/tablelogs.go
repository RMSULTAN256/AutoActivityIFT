package ui

import (
	"encoding/json"
	"log"
	"masterc/logic"
	"masterc/models"
	"net/url"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
)

func Logging() fyne.CanvasObject {
	var TablesData []models.LogsTable
	var dataMutex sync.RWMutex

	table := widget.NewTable(
		func() (int, int) {
			dataMutex.RLock()
			defer dataMutex.RUnlock()
			return len(TablesData) + 1, 7
		},
		func() fyne.CanvasObject {

			label := widget.NewLabel("No Data")
			label.Truncation = fyne.TextTruncateEllipsis
			label.Alignment = fyne.TextAlignCenter
			return label
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)

			label.Alignment = fyne.TextAlignCenter
			label.TextStyle = fyne.TextStyle{}
			label.Importance = widget.MediumImportance

			if id.Row == 0 {
				label.TextStyle = fyne.TextStyle{Bold: true}
				switch id.Col {
				case 0:
					label.SetText("Nama")
				case 1:
					label.SetText("Sebelum")
				case 2:
					label.SetText("Mulai")
				case 3:
					label.SetText("Berikutnya")
				case 4:
					label.SetText("Aksi")
				case 5:
					label.SetText("Status")
				case 6:
					label.SetText("Message")
				}
				return
			}

			// --- DATA ---
			dataMutex.RLock()

			if id.Row-1 >= len(TablesData) {
				dataMutex.RUnlock()
				label.SetText("")
				return
			}

			u := TablesData[id.Row-1]
			dataMutex.RUnlock()

			// Isi Data
			label.Alignment = fyne.TextAlignCenter
			switch id.Col {
			case 0:
				label.SetText(u.Name)
			case 1:
				label.SetText(u.PrevTime)
			case 2:
				label.SetText(u.StartTime)
			case 3:
				label.SetText(u.NextSchedule)
			case 4:
				label.SetText(u.Action)
			case 5:
				label.SetText(u.Status)

				if u.Status == "ERROR" {
					label.Importance = widget.DangerImportance
				} else {
					label.Importance = widget.MediumImportance
				}

			case 6:
				label.SetText(u.ErrorMessage)
				label.Alignment = fyne.TextAlignLeading
				label.Wrapping = fyne.TextWrapWord
			}
		},
	)

	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 200)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 200)
	table.SetColumnWidth(4, 100)
	table.SetColumnWidth(5, 100)
	table.SetColumnWidth(6, 1200)

	labelwarn := binding.NewString()
	labelwarn.Set("Klik tombol clear untuk menghapus semua data")

	btnclear := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {

		err := logic.ResetLogsData()
		if err != nil {
			labelwarn.Set("Gagal menghapus gambar" + err.Error())
		}
		dataMutex.Lock()
		TablesData = []models.LogsTable{}
		dataMutex.Unlock()

		table.Refresh()

		labelwarn.Set("Data berhasil dihapus")
		go func() {
			time.Sleep(2 * time.Second)

			labelwarn.Set("Klik tombol clear untuk menghapus semua data")

		}()
	})
	labelWithla := widget.NewLabelWithData(labelwarn)
	BtnOnly := container.NewHBox(btnclear, labelWithla)

	content := container.NewBorder(
		BtnOnly,
		nil, nil, nil,
		table,
	)

	scroll := container.NewScroll(content)
	go func() {
		u := url.URL{
			Scheme: "ws",
			Host:   "127.0.0.1:5544",
			Path:   "/api/v1/ws/logs",
		}
		log.Println("connecting to WS:", u.String())

		for {
			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Println("ws dial error (retrying in 2s):", err)
				time.Sleep(2 * time.Second)
				continue
			}

			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					log.Println("ws read error:", err)
					break
				}

				func() {
					dataMutex.Lock()
					defer dataMutex.Unlock()

					var incomingList []models.LogsTable

					if err := json.Unmarshal(msg, &incomingList); err == nil {
						TablesData = incomingList
					} else {
						var singleLog models.LogsTable
						if err2 := json.Unmarshal(msg, &singleLog); err2 == nil {
							found := false

							for i, row := range TablesData {
								if row.Name == singleLog.Name && row.StartTime == singleLog.StartTime {
									TablesData[i] = singleLog
									found = true
									break
								}
							}
							if !found {
								TablesData = append(TablesData, singleLog)
							}
						}
					}
				}()
				fyne.Do(func() {
					table.Refresh()
					table.ScrollToBottom()
				})
				time.Sleep(100 * time.Millisecond)
			}
			conn.Close()
		}
	}()
	return scroll
}

//func LogsConfig() fyne.CanvasObject {
//
//}
