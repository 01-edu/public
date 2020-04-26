package correct

func UltimateDivMod(a, b *int) {
	temp := *a
	*a = *a / *b
	*b = temp % *b
}
