package main

import "fyne.io/fyne/v2"

type StackItem struct {
	Symbol rune
	Widget *fyne.Container
	Xaxis  float32
}
