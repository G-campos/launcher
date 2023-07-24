package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func LauncherRun(w fyne.Window) {
	w.SetContent(widget.NewLabel("Hello World!"))
	w.Resize(fyne.NewSize(300, 500))

	w.Show()
}
