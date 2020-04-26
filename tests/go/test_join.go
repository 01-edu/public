package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	seps := []string{" ", "-", " ,", "_", "SPC", " . "}

	args := [][]string{
		{"hello", "how", "are", "you", "doing"},
		{"fine", "and", "you"},
		{"I'm", "O.K."},
	}

	for i := 0; i < 5; i++ {
		// random position for the slice of arguments
		posA := z01.RandIntBetween(0, len(args)-1)
		// random position for the slice of separators
		posS := z01.RandIntBetween(0, len(seps)-1)

		z01.Challenge("Join", student.Join, correct.Join, args[posA], seps[posS])
	}
}
