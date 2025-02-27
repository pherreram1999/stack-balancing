package balancingSymbols

type BalancingSymbol map[rune]bool

func GetPushSymbols() BalancingSymbol {
	return map[rune]bool{
		'{': true,
		'[': true,
		'(': true,
	}
}

func GetPopSymbols() BalancingSymbol {
	return map[rune]bool{
		'}': true,
		']': true,
		')': true,
	}
}

// Is indica si el simbolo dado se debe o no agregar a la cola
func (symbols *BalancingSymbol) Is(symbol rune) bool {
	if _, ok := (*symbols)[symbol]; ok {
		return true
	}
	return false
}
