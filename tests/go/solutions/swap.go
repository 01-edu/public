package solutions

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
