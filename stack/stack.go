package stack

type StackList struct {
	Symbol rune
	Next   *StackList
}

type StackEachFunc func(symbol rune)

func Push(stack **StackList, symbol rune) {
	nuevoNodo := &StackList{
		Symbol: symbol,
	} // nuevo nodo insertar
	nuevoNodo.Next = *stack // indicamos que el nuevo nodo si siguente apunta a la pila (HEAD)
	*stack = nuevoNodo      // ahora la pila (HEAD) apunta al nuevo nodo
}

func Pop(stack **StackList) rune {
	if *stack == nil {
		return 0 // sin nada en la pila
	}
	popNode := *stack     // el nodo que se saca es el primero
	*stack = popNode.Next // el nodo que sacamos, su siguiente nodo ahora es la pila (HEAD)
	popNode.Next = nil    // del nodo de sacamos quitamos referencias
	return popNode.Symbol // regresamos el symbolo asociado
}

func (s *StackList) ForEach(callback StackEachFunc) {
	nav := s
	for nav != nil {
		callback(nav.Symbol)
		nav = nav.Next
	}
}
