package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	"github.com/leofigy/valk/server"
)

func main() {
	a := app.New()
	w := a.NewWindow("Valk")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Valk",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	//addLabel := widget.NewLabel("Listen address")
	addValue := widget.NewEntry()
	addValue.SetPlaceHolder("enter a listener address in format 0.0.0.0:5000")
	addValue.Resize(fyne.NewSize(100, 100))
	//label2 := widget.NewLabel("Server Name")
	serverName := widget.NewEntry()
	serverName.SetPlaceHolder("Enter a friendly name for the server")

	// wires definition
	state := make(chan server.ServerConfig)

	go server.InitBackendListener(state)

	stopBotton := widget.NewButton("stop", func() {
		state <- server.ServerConfig{
			Current: server.Stop,
			Address: "",
		}
	})

	stopBotton.Importance = widget.DangerImportance

	// container
	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "address value", Widget: addValue},
			{Text: "server friendly name", Widget: serverName},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", addValue.Text)
			state <- server.ServerConfig{
				Current: server.Start,
				Address: addValue.Text,
			}
		},
		SubmitText: "Start",
	}

	controls := container.NewHSplit(form, stopBotton)
	controls.SetOffset(1.0)
	w.SetContent(controls)
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	//w.Resize(fyne.NewSize(350, 120))
	w.Resize(fyne.NewSize(700, 100))
	w.ShowAndRun()
}
