package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"time"
)

var (
	progress    *widget.ProgressBar
	infProgress *widget.ProgressBarInfinite
	endProgress chan interface{}
)

func SplashScreenRun() {
	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		w := drv.CreateSplashWindow()
		//go func() {
		//	startProgress()
		//}()

		windowGroup := container.NewVBox(
			widget.NewLabel("Hello World"),
			widget.NewFileIcon(fine.)
			widget.NewProgressBar(),
			//makeProgress(w),
		)

		w.SetContent(windowGroup)
		w.Resize(fyne.NewSize(500, 300))
		w.CenterOnScreen()
		w.Show()

		//go func() {
		//	time.Sleep(time.Second * 5)
		//	stopProgress()
		//	w.Close()
		//}()
	} else {
		fmt.Println("Erro splash screen")
	}
}

func makeProgress(_ fyne.Window) fyne.CanvasObject {
	stopProgress()

	progress = widget.NewProgressBar()

	infProgress = widget.NewProgressBarInfinite()
	endProgress = make(chan interface{}, 1)
	startProgress()

	return container.NewVBox(
		widget.NewLabel("Percent"), progress,
	)
}

func startProgress() {
	progress.SetValue(0)
	select { // ignore stale end message
	case <-endProgress:
	default:
	}

	go func() {
		end := endProgress
		num := 0.0
		for num < 1.0 {
			time.Sleep(16 * time.Millisecond)
			select {
			case <-end:
				return
			default:
			}

			progress.SetValue(num)
			num += 0.002
		}

		progress.SetValue(1)

		// TODO make sure this resets when we hide etc...
		stopProgress()
	}()
	infProgress.Start()
}

func stopProgress() {
	if !infProgress.Running() {
		return
	}

	infProgress.Stop()
	endProgress <- struct{}{}
}
