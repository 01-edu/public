package student_test

import (
	"math/rand"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestIsPrintable(t *testing.T) {
	// 15 unvalid strings in the table
	table := z01.MultRandASCII()

	// 15 valid lowercase strings of random size between 1 and 20 letters in the table
	for i := 0; i < 15; i++ {
		letters := []rune("\a\b\f\r\n\v\t")
		size := z01.RandIntBetween(1, 20)
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
		z01.Challenge(t, student.IsPrintable, solutions.IsPrintable, arg)
	}
}
