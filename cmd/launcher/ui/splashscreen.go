package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func SplashScreenRun() {
	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		w := drv.CreateSplashWindow()
		w.SetContent(widget.NewLabel("Hello World"))
		w.Resize(fyne.NewSize(500, 300))
		w.CenterOnScreen()
		w.Show()
	} else {
		fmt.Println("Erro splash screen")
	}
}
