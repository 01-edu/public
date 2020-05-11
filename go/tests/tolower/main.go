package main

import (
	"strings"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func toLower(s string) string {
	return strings.ToLower(s)
}

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello! How are you?",
	)
	for _, arg := range table {
		lib.Challenge("ToLower", student.ToLower, toLower, arg)
	}
}
