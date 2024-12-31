package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
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

func StackWidget(pathBind, counterBind, entryBind binding.String) *fyne.Container {

	timerBind := binding.NewFloat()
	timerBind.Set(minTimer)

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
	pushLbl := canvas.NewText("PUSH:", color.Black)
	popLbl := canvas.NewText("POP:", color.Black)
	popLbl.Move(fyne.NewPos(100, 0))
	stackContainer := container.NewWithoutLayout(pushLbl, popLbl)

	// symbols

	var stack *stacklist.StackList[*StackItem]

	// moveStack mueva la pila entera,y los desplaza en x "axis" veces
	moveStack := func(axis float32) {
		nav := stack
		t, _ := timerBind.Get()
		duration := time.Millisecond * time.Duration(t)
		for nav != nil {
			item := nav.Item
			// lo desplazamos hacia la derecha "axis" veces para cada elmento
			newAxis := item.Xaxis + axis

			move := canvas.NewPositionAnimation(
				fyne.NewPos(item.Xaxis, topHeigth),
				fyne.NewPos(newAxis, topHeigth),
				duration,
				item.Widget.Move,
			)
			item.Xaxis = newAxis
			move.Start() // anima la transicion
			nav = nav.Next
		}
		time.Sleep(duration)

	}

	clearStack := func() {
		for stack != nil {
			stack.Item.Widget.Hide()
			stack = stack.Next
		}
	}

	stackActionBtn := widget.NewButton("Balance Stack", func() {
		clearStack()
		stack = nil // limpiamos la pila
		t, _ := timerBind.Get()
		duration := time.Millisecond * time.Duration(t)
		text, err := entryBind.Get()
		if err != nil {
			ShowError(err)
			return
		}

		if text[0] != '0' {
			ShowError(errors.New("La cadena no inicia con 0"))
			return
		}

		counter := 0 // lleva el conteo de la pila

		for _, char := range text {
			_ = symbolBind.Set(string(char))
			if char == '0' { // push
				counter++
				symbolBox := NewBoxSymbol(char)
				stackContainer.Add(symbolBox)
				symbolBox.Move(fyne.NewPos(50, 0))
				moveStack(boxSize)
				time.Sleep(duration) // esperar y desplazamos
				// movemos nuestra nuestro simbolo de entrada a la posicion del stack
				moveSymbolBox := canvas.NewPositionAnimation(
					fyne.NewPos(50, 0),
					fyne.NewPos(0, topHeigth),
					duration,
					symbolBox.Move,
				)
				moveSymbolBox.Start()
				stacklist.Push(
					&stack,
					&StackItem{
						Symbol: char,
						Widget: symbolBox,
					},
				)
				time.Sleep(duration) // esperamos que se termine de animar
			} else {
				counter--
				popItem := stacklist.Pop(&stack)
				if popItem == nil {
					ShowInfo("Pila no balanceada", "Se trata de sacar mas elemento de los existentes")
					clearStack()
					return
				}
				popMove := canvas.NewPositionAnimation(
					fyne.NewPos(popItem.Xaxis, topHeigth),
					fyne.NewPos(150, 0),
					duration,
					popItem.Widget.Move,
				)
				popMove.Start()
				time.Sleep(duration * 2) // esperamos la animacion
				// una vez que termine lo quitamos
				popItem.Widget.Hide()
				popItem = nil // para que lo recoja el colector
				moveStack(-boxSize)

			}
			_ = counterBind.Set(fmt.Sprintf("%d", counter))
			time.Sleep(duration)

		}

		if counter > 0 {
			ShowInfo("NO balanceado :(", "La pila aun presenta elementos sin más simbolos que procesar")
		} else {
			ShowInfo("Pila banceada", "Sin elementos en la pila y sin mas simbolos que procesar")
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
