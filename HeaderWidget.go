package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
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

	readFileBtn := widget.NewButton("Select File", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				ShowError(err)
				return
			}
			if reader == nil {
				return
			}
			pathBind.Set(reader.URI().Path())

		}, Window)
	})

	return container.New(
		layout.NewHBoxLayout(),
		readFileBtn,
		pathLbl,
		counterTitleLbl,
		counterLbl,
	)
}
