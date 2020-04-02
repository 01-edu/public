package solutions

func NRune(s string, n int) rune {
	if n > len(s) || n < 1 {
		return 0
	}
	runes := []rune(s)
	return runes[n-1]
}
