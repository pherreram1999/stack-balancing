package main

import (
	"fmt"
	"stackbalancing/stack"
)

func main() {
	word := "Hello"

	var pila *stack.StackList[rune]

	fmt.Println("----- Pushed -----")
	for _, runa := range word {
		fmt.Println("pushed: ", string(runa))
		stack.Push(&pila, runa)
	}
	fmt.Println("----- Pop -----")
	for i := 0; i < len(word); i++ {
		s := stack.Pop(&pila)
		if s == 0 {
			fmt.Println("Sin nada en la pila")
		}
		fmt.Println("pop: ", string(s))
	}
}
