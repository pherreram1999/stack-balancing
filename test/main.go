package main

import (
	"fmt"
	"stackbalancing/stacklist"
)

func main() {
	word := "Hello"

	var pila *stacklist.StackList[rune]

	fmt.Println("----- Pushed -----")
	for _, runa := range word {
		fmt.Println("pushed: ", string(runa))
		stacklist.Push(&pila, runa)
	}
	fmt.Println("----- Pop -----")
	for i := 0; i < len(word); i++ {
		s := stacklist.Pop(&pila)
		if s == 0 {
			fmt.Println("Sin nada en la pila")
		}
		fmt.Println("pop: ", string(s))
	}
}
