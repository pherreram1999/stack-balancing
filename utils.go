package main

import "fyne.io/fyne/v2/dialog"

func ShowError(err error) {
	dialog.ShowError(err, Window)
}

func ShowInfo(title, msg string) {
	dialog.ShowInformation(title, msg, Window)
}
