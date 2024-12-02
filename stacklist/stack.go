package stacklist

type StackList[T any] struct {
	Item T
	Next *StackList[T]
}

func ForEach[T any](stack **StackList[T], callback func(item T)) {
	nav := *stack
	for nav != nil {
		callback(nav.Item)
		nav = nav.Next
	}
}

func Push[T any](stack **StackList[T], item T) {
	nuevoNodo := &StackList[T]{
		Item: item,
	} // nuevo nodo insertar
	nuevoNodo.Next = *stack // indicamos que el nuevo nodo si siguente apunta a la pila (HEAD)
	*stack = nuevoNodo      // ahora la pila (HEAD) apunta al nuevo nodo
}

func Pop[T any](stack **StackList[T]) T {
	if *stack == nil {
		var zero T
		return zero
	}
	popNode := *stack     // el nodo que se saca es el primero
	*stack = popNode.Next // el nodo que sacamos, su siguiente nodo ahora es la pila (HEAD)
	popNode.Next = nil    // del nodo de sacamos quitamos referencias
	return popNode.Item   // regresamos el symbolo asociado
}
