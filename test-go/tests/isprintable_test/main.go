package main

import (
	"math/rand"
	"unicode"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func main() {
	// 15 unvalid strings in the table
	table := lib.MultRandASCII()

	// 15 valid lowercase strings of random size between 1 and 20 letters in the table
	for i := 0; i < 15; i++ {
		letters := []rune("\a\b\f\r\n\v\t")
		size := lib.RandIntBetween(1, 20)
		r := make([]rune, size)
		for i := range r {
			r[i] = letters[rand.Intn(len(letters))]
		}
		table = append(table, string(r))
	}

	// Special cases added to table
	table = append(table,
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"HelloHowareyou",
		"What's this 4?",
		"Whatsthis4",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!",
		"0123456789",
		"01,02,03",
		"abcdefghijklmnopqrstuvwxyz",
		"hello",
		"hello!",
		"hello\n",
		"\n",
	)
	for _, arg := range table {
		lib.Challenge("IsPrintable", student.IsPrintable, isPrintable, arg)
	}
}
