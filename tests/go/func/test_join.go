package main

import (
	"../lib"
	"./correct"
	"./student"
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
		posA := lib.RandIntBetween(0, len(args)-1)
		// random position for the slice of separators
		posS := lib.RandIntBetween(0, len(seps)-1)

		lib.Challenge("Join", student.Join, correct.Join, args[posA], seps[posS])
	}
}
