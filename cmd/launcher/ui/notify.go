package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Notify(title, content, label string) {
	widget.NewButton(label, func() {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   title,
			Content: content,
		})
	})
}
