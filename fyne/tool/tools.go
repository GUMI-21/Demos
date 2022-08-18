package tool

import (
	"fyne.io/fyne/v2/widget"
	"time"
)

func UpdateTime(clock *widget.Label) {
	clock.SetText(time.Now().Format("Time: 03:04:05"))
}
