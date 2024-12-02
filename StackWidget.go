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

const boxSize = 40

func NewBoxSymbol(symbol rune) *fyne.Container {
	SymbolBox := canvas.NewRectangle(color.RGBA{R: 241, G: 245, B: 249, A: 255})
	SymbolBoxLbl := widget.NewLabel(string(symbol))
	SymbolBoxLbl.TextStyle.Bold = true
	SymbolBox.Resize(fyne.NewSize(boxSize, boxSize))
	return container.NewWithoutLayout(SymbolBox, SymbolBoxLbl)
}

const minTimer float64 = 400
const topHeigth float32 = 50

func StackWidget(pathBind binding.String) *fyne.Container {

	timerBind := binding.NewFloat()
	timerBind.Set(minTimer)

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

	// slide timer
	slideTimer := widget.NewSlider(minTimer, 1000)
	slideTimer.Bind(timerBind)
	slideContainer := container.NewWithoutLayout(slideTimer)
	slideTimer.Resize(fyne.NewSize(250, boxSize))
	slideTimeLbl := widget.NewLabel("")
	slideInputContent := container.New(
		layout.NewHBoxLayout(),
		canvas.NewText("Time", color.Black),
		slideTimeLbl,
		slideContainer,
	)
	// input symbol
	symbolBind := binding.NewString()
	SymbolBox := canvas.NewRectangle(color.RGBA{R: 241, G: 245, B: 249, A: 255})
	SymbolBoxLbl := widget.NewLabel("")
	SymbolBoxLbl.TextStyle.Bold = true
	SymbolBoxLbl.Bind(symbolBind)
	SymbolBox.Resize(fyne.NewSize(boxSize, boxSize))

	SymbolBoxContainer := container.NewWithoutLayout(SymbolBox, SymbolBoxLbl)
	SymbolBoxContainer.Resize(fyne.NewSize(500, boxSize))
	inputSymbolLbl := widget.NewLabel("Input")
	InputSymbolContainer := container.New(layout.NewHBoxLayout(), inputSymbolLbl, SymbolBoxContainer)

	// stack Container
	stackContainer := container.NewWithoutLayout()

	// symbols

	pushSymbols := balancingSymbols.GetPushSymbols()
	popSymbols := balancingSymbols.GetPopSymbols()

	var stack *stacklist.StackList[*StackItem]

	moveStack := func(axis float32) {
		nav := stack
		t, _ := timerBind.Get()
		duration := time.Millisecond * time.Duration(t)
		for nav != nil {
			item := nav.Item
			newAxis := item.Xaxis + axis

			move := canvas.NewPositionAnimation(
				fyne.NewPos(item.Xaxis, topHeigth),
				fyne.NewPos(newAxis, topHeigth),
				duration,
				item.Widget.Move,
			)
			item.Xaxis = newAxis
			move.Start()
			nav = nav.Next
		}
		time.Sleep(duration)

	}

	stackActionBtn := widget.NewButton("Balance Stack", func() {
		t, _ := timerBind.Get()
		duration := time.Millisecond * time.Duration(t)
		text, err := getFileContent()
		if err != nil {
			ShowError(err)
			return
		}

		for _, char := range text {
			symbolBind.Set(string(char))
			if pushSymbols.Is(char) {
				moveStack(boxSize)
				symbolBox := NewBoxSymbol(char)
				stackContainer.Add(symbolBox)
				move := canvas.NewPositionAnimation(
					fyne.NewPos(0, 0),
					fyne.NewPos(0, topHeigth),
					duration, symbolBox.Move,
				)
				move.Start()
				stacklist.Push(
					&stack,
					&StackItem{
						Symbol: char,
						Widget: symbolBox,
					},
				)
				time.Sleep(duration)
			} else if popSymbols.Is(char) {
				moveStack(-boxSize)
				stacklist.Pop(&stack)
			}
			time.Sleep(duration)
		}

	})

	headerContainer := container.New(
		layout.NewHBoxLayout(),
		stackActionBtn,
		InputSymbolContainer,
	)

	return container.New(
		layout.NewVBoxLayout(),
		headerContainer,
		slideInputContent,
		stackContainer,
	)
}
