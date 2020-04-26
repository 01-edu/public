package correct

func DivMod(a, b int, div, mod *int) {
	*div = a / b
	*mod = a % b
}
