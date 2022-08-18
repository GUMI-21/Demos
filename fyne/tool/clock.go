package tool

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

// Clock var wg sync.WaitGroup
//时钟
func Clock() {

	a := app.New()
	w := a.NewWindow("clock")

	clock := widget.NewLabel("")
	UpdateTime(clock)
	w.SetContent(clock)
	w.Resize(fyne.NewSize(300, 300))

	go func() {
		for range time.Tick(time.Second) {
			UpdateTime(clock)
		}
	}()
	w.ShowAndRun()
}
