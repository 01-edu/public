package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestIsUpper(t *testing.T) {
	// 15 unvalid strings in the table
	table := z01.MultRandASCII()

	// 15 valid lowercase strings of random size between 1 and 20 letters in the table
	for i := 0; i < 15; i++ {
		size := z01.RandIntBetween(1, 20)
		randLow := z01.RandLower()
		if len(randLow) <= size {
			table = append(table, randLow)
		} else {
			table = append(table, randLow[:size])
		}
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
		"HELLO",
		"HELLO!",
	)
	for _, arg := range table {
		z01.Challenge(t, student.IsUpper, solutions.IsUpper, arg)
	}
}
