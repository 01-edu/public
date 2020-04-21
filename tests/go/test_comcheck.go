package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	table := append(z01.MultRandWords(),
		"01",
		"galaxy",
		"galaxy01",
		" 01",
		"as das d 01 asd",
		"as galaxy d 12",
		"as ds galaxy 01 asd")

	for _, s := range table {
		z01.ChallengeMain("comcheck", strings.Fields(s)...)
	}
}
