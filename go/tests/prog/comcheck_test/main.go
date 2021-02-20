package main

import (
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		"01",
		"galaxy",
		"galaxy01",
		" 01",
		"as das d 01 asd",
		"as galaxy d 12",
		"as ds galaxy 01 asd")

	for _, s := range table {
		lib.ChallengeMain("comcheck", strings.Fields(s)...)
	}
}
