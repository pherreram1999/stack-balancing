package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
)

var windowParent fyne.Window

func main() {
	a := app.New()
	windowParent = a.NewWindow("Stack Balancing")

	codeTextBindig := binding.NewString()

	// button OpenFile
	btnFile := widget.NewButton("Open File", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, windowParent)
			}
			path := reader.URI().Path()
			data, err := os.ReadFile(path)
			if err != nil {
				dialog.ShowError(err, windowParent)
			}
			codeTextBindig.Set(string(data))
		}, windowParent)
	})

	buttonContainer := container.New(layout.NewHBoxLayout(), btnFile)

	textarea := widget.NewMultiLineEntry()
	textarea.Bind(codeTextBindig)
	textarea.SetMinRowsVisible(35)

	cont := container.NewVBox(buttonContainer, textarea)

	windowParent.SetContent(cont)

	windowParent.Resize(fyne.NewSize(1200, 720))

	windowParent.ShowAndRun()
}
