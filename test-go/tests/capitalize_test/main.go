package main

import (
	"strings"
	"unicode"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func capitalize(s string) string {
	r := []rune(strings.ToLower(s))

	if unicode.Is(unicode.Latin, r[0]) {
		r[0] = unicode.ToUpper(r[0])
	}

	for i := 1; i < len(r); i++ {
		if !unicode.IsDigit(r[i-1]) && !unicode.Is(unicode.Latin, r[i-1]) {
			r[i] = unicode.ToUpper(r[i])
		}
	}
	return string(r)
}

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello! How are you? How+are+things+4you?",
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"9a",
		"9a LALALA!",
	)
	for _, arg := range table {
		lib.Challenge("Capitalize", student.Capitalize, capitalize, arg)
	}
}
