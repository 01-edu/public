package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func strRev(s string) string {
	runes := []rune(s)
	i := 0
	j := len(runes) - 1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello!",
		"Bonjour!",
		"Hola!",
	)
	for _, arg := range table {
		lib.Challenge("StrRev", student.StrRev, strRev, arg)
	}
}
