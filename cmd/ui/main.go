package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	hello := widget.NewLabel("")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Pal :)")
		}),
	))

	w.Show()

	go func() {
		for range time.Tick(time.Second) {
			updateTime(hello)
		}
	}()

	// second window
	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	a.Run()

	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
