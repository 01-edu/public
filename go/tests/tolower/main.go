package main

import (
	"strings"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func main() {
	table := append(lib.MultRandASCII(), "Hello! How are you?")
	for _, arg := range table {
		lib.Challenge("ToLower", student.ToLower, strings.ToLower, arg)
	}
}
