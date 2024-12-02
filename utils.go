package main

import "fyne.io/fyne/v2/dialog"

func ShowError(err error) {
	dialog.ShowError(err, Window)
}
