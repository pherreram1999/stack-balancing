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

type StackSymbolWidget struct {
	Symbol string
	Wg     fyne.Widget
}

var windowParent fyne.Window

var codeTextBindig binding.String
var appFyne fyne.App

func main() {
	appFyne = app.NewWithID("stack-balancing")
	windowParent = appFyne.NewWindow("Stack Balancing")

	codeTextBindig = binding.NewString()

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

	balancingBtn := widget.NewButton("Balancing", BalancingWindow)

	buttonContainer := container.New(layout.NewHBoxLayout(), btnFile, balancingBtn)

	textarea := widget.NewMultiLineEntry()
	textarea.Bind(codeTextBindig)
	textarea.SetMinRowsVisible(35)

	cont := container.NewVBox(buttonContainer, textarea)

	windowParent.SetContent(cont)

	windowParent.Resize(fyne.NewSize(1200, 720))

	windowParent.ShowAndRun()
}

func BalancingWindow() {
	/*codeText, err := codeTextBindig.Get()
	if err != nil {
		dialog.ShowError(err, windowParent)
	}*/

	win := appFyne.NewWindow("Stack animation")

	rowContainer := container.New(layout.NewVBoxLayout())

	win.SetContent(rowContainer)

	win.Show()
}
