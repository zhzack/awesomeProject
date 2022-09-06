package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func mainss() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetFullScreen(true)

	hello := widget.NewLabel("Hello Fyne!")

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			w.SetFullScreen(false)
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
