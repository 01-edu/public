package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghi jklmnop qrstuvwxyz ABCDEFGHI JKLMNOPQR STUVWXYZ ! ",
	)
	for _, s := range table {
		lib.ChallengeMain("switchcase", s)
	}
}
