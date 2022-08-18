package tool

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

//var wg sync.WaitGroup

func Clock() {

	a := app.New()
	w := a.NewWindow("clock")

	clock := widget.NewLabel("")
	UpdateTime(clock)
	w.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			UpdateTime(clock)
		}
	}()
	w.ShowAndRun()
}
