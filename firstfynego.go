package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hi world")

	w.SetContent(widget.NewLabel("Hi world again"))
	w.ShowAndRun()
}
