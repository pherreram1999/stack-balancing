package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
	"stackbalancing/balancingSymbols"
	"stackbalancing/stacklist"
	"time"
)

func NewBoxSymbol(symbol rune) *fyne.Container {
	SymbolBox := canvas.NewRectangle(color.RGBA{R: 241, G: 245, B: 249, A: 255})
	SymbolBoxLbl := widget.NewLabel(string(symbol))
	SymbolBoxLbl.TextStyle.Bold = true
	SymbolBox.Resize(fyne.NewSize(35, 35))
	return container.NewWithoutLayout(SymbolBox, SymbolBoxLbl)
}

func StackWidget(pathBind binding.String) *fyne.Container {

	getFileContent := func() (string, error) {
		path, err := pathBind.Get()

		if err != nil {
			return "", err
		}

		res, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}

		return string(res), nil
	}

	symbolBind := binding.NewString()
	SymbolBox := canvas.NewRectangle(color.RGBA{R: 241, G: 245, B: 249, A: 255})
	SymbolBoxLbl := widget.NewLabel("")
	SymbolBoxLbl.TextStyle.Bold = true
	SymbolBoxLbl.Bind(symbolBind)
	SymbolBox.Resize(fyne.NewSize(35, 35))

	SymbolBoxContainer := container.NewWithoutLayout(SymbolBox, SymbolBoxLbl)
	inputSymbolLbl := widget.NewLabel("Input")
	InputSymbolContainer := container.New(layout.NewHBoxLayout(), inputSymbolLbl, SymbolBoxContainer)

	// stack Container
	stackContainer := container.NewWithoutLayout()

	// symbols

	pushSymbols := balancingSymbols.GetPushSymbols()
	popSymbols := balancingSymbols.GetPopSymbols()

	stack := &stacklist.StackList[*StackItem]{}

	stackActionBtn := widget.NewButton("Balance Stack", func() {
		text, err := getFileContent()
		if err != nil {
			ShowError(err)
			return
		}

		for _, char := range text {
			symbolBind.Set(string(char))
			if pushSymbols.Is(char) {

				symbolBox := NewBoxSymbol(char)

				stackContainer.Add(symbolBox)

				stacklist.Push(
					&stack,
					&StackItem{
						Symbol: char,
						Widget: symbolBox,
					},
				)
			} else if popSymbols.Is(char) {
				stacklist.Pop(&stack)
			}
			time.Sleep(time.Second)
		}
	})

	return container.New(
		layout.NewVBoxLayout(),
		container.New(
			layout.NewHBoxLayout(),
			stackActionBtn,
			InputSymbolContainer,
		),
		stackContainer,
	)
}
