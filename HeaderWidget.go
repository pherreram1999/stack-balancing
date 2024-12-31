package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HeaderWidget(pathBind, counterBind binding.String) *fyne.Container {
	pathLbl := widget.NewLabel("")
	pathLbl.Bind(pathBind)
	counterTitleLbl := widget.NewLabel("Stack Count")
	counterTitleLbl.TextStyle.Bold = true
	counterLbl := widget.NewLabel("")
	counterLbl.Bind(counterBind)

	return container.New(
		layout.NewHBoxLayout(),
		counterTitleLbl,
		counterLbl,
	)
}
