package tool

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func HelloWorld() {
	myApp := app.New()
	myWindow := myApp.NewWindow("hello")
	myWindow.SetContent(widget.NewLabel("hello"))

	clock := widget.NewLabel("")
	myWindow.SetContent(clock)
	clock.SetText(time.Now().Format("Time: 03:04:05"))

	myWindow.Show()
	myApp.Run()
	fmt.Println("exit")
}
