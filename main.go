package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
)

var App fyne.App
var Window fyne.Window

func main() {
	App = app.NewWithID("stacklist-balancing")

	Window = App.NewWindow("Stack Balancing")

	Window.Resize(fyne.NewSize(1600, 800))

	// binders =
	pathBind := binding.NewString()

	mainContent := container.New(
		layout.NewVBoxLayout(),
		HeaderWidget(pathBind),
		StackWidget(pathBind),
	)

	Window.SetContent(mainContent)

	Window.ShowAndRun()
}
