package main

import (
	"fmt"
	"image/color"
	"time"

	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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
	w2.SetContent(widget.NewButton("generate color", func() {
		w3 := a.NewWindow("random color")
		w3.Resize(fyne.NewSize(100, 100))
		myCanvas := w3.Canvas()
		blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
		rect := canvas.NewRectangle(blue)
		myCanvas.SetContent(rect)
		w3.Show()
		go func() {
			time.Sleep(time.Second * 5)
			rect.FillColor = randomColor()
			rect.Refresh()
		}()

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

func randomColor() color.NRGBA {
	rand.Seed(time.Now().UnixNano())
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	return color.NRGBA{R: r, G: g, B: b, A: 255}
}
