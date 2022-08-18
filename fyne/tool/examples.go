package tool

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

// TowWindows 实现两个窗口
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

//onChanged文本变化
func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello" + content + "!")
	}
	return out, in
}

func HelloWorld() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(makeUI()))
	w.ShowAndRun()
}

// Clock 一个时钟小部件
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

//画布和画布对象
func Canvas() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas() //画布

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect)
}
func setContentToText(c fyne.Canvas) {
	//green := color.NRGBA{R: 0xff, G: 0x33, A: 0xff}

}
