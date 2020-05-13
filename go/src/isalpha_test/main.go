package main

import (
	"unicode"

	student "student"

	"lib"
)

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Hex_Digit, r) {
			return false
		}
	}
	return true
}

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello! €How are you?",
		"a",
		"z",
		"!",
		"HelloHowareyou",
		"What's this 4?",
		"Whatsthis4",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!",
	)
	for _, arg := range table {
		lib.Challenge("IsAlpha", student.IsAlpha, isAlpha, arg)
	}
}
