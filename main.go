package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	a := app.New()

	w := a.NewWindow("Launcher App")
	w.Resize(fyne.NewSize(400, 200))

	str := binding.NewString()
	err := str.Set("Initial value")
	if err != nil {
		return
	}

	text := widget.NewLabelWithData(str)
	w.SetContent(text)

	go func() {
		time.Sleep(time.Second * 2)
		err := str.Set("A new string")
		if err != nil {
			return
		}
	}()

	w.ShowAndRun() // Finally Running our App
}
