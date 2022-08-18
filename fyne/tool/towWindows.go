package tool

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func TowWindows() {
	a := app.New()
	w := a.NewWindow("Hello")

	w.SetContent(widget.NewLabel("Hello"))
	w.Show()

	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More Content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	a.Run()
}
