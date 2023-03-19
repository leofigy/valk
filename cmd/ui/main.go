package main

import (
	"crypto/tls"
	"log"
	"os"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	"github.com/leofigy/valk/server"
)

const (
	VALK_CERTS = "VALK_CERTS"
)

var (
	certsFolder = "./certs"
	serverCert  = "server.crt"
	serverKey   = "server.key"
)

func init() {
	if os.Getenv(VALK_CERTS) != "" {
		certsFolder = os.Getenv(VALK_CERTS)
	}

	serverCert = path.Join(certsFolder, serverCert)
	serverKey = path.Join(certsFolder, serverKey)
}

func main() {
	var security *tls.Config

	if cer, err := tls.LoadX509KeyPair(serverCert, serverKey); err == nil {
		security = &tls.Config{Certificates: []tls.Certificate{cer}}
	} else {
		log.Println("WARNING: Loading ", err)
	}

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
	addValue.Resize(fyne.NewSize(100, 100))
	addValue.SetText("0.0.0.0:4001")
	//label2 := widget.NewLabel("Server Name")
	serverName := widget.NewEntry()
	serverName.SetText("ValkServer")

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
				Current:  server.Start,
				Address:  addValue.Text,
				Security: security,
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
