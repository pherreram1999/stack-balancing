package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

var App fyne.App
var Window fyne.Window

const maxLen = 100_000
const maxRand = 100

func main() {
	App = app.NewWithID("stacklist-balancing")

	Window = App.NewWindow("Stack Balancing")

	Window.Resize(fyne.NewSize(800, 350))

	// binders
	pathBind := binding.NewString()
	counterBind := binding.NewString()
	entryLenBind := binding.NewString()

	entryBind := binding.NewString()
	entryInput := widget.NewEntry()
	entryInput.Bind(entryBind)
	entryLen := 0
	entryInput.OnChanged = func(s string) {
		entryLen = len(s)
		_ = entryLenBind.Set(fmt.Sprintf("%d", entryLen))
	}
	entryLenLbl := widget.NewLabel("")
	entryLenLbl.Bind(entryLenBind)

	// semilla randmon
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	// methods  to entry
	makeRandomEntryBtn := widget.NewButton("Random Entry", func() {
		maxLength := r.Intn(maxRand/2) * 2
		middle := maxLength / 2
		randEntry := ""
		for i := 0; i < maxLength; i++ {
			if i < middle {
				randEntry += "0"
			} else {
				randEntry += "1"
			}
		}
		entryBind.Set(randEntry)
	})

	mainContent := container.New(
		layout.NewVBoxLayout(),
		entryInput,
		container.New(
			layout.NewHBoxLayout(),
			makeRandomEntryBtn,
			widget.NewLabel("Entry Len"),
			entryLenLbl,
		),
		HeaderWidget(pathBind, counterBind),
		StackWidget(pathBind, counterBind, entryBind, &entryLen),
	)

	Window.SetContent(mainContent)
	Window.ShowAndRun()
}
