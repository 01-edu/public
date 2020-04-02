package main

func LastRune(s string) rune {
	runes := []rune(s)
	index := len(runes) - 1
	return runes[index]
}

func main() {
}
