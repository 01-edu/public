package main

import "github.com/01-edu/z01"

func main() {
	table := append(z01.MultRandWords(),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghi jklmnop qrstuvwxyz ABCDEFGHI JKLMNOPQR STUVWXYZ ! ",
	)
	for _, s := range table {
		z01.ChallengeMain("switchcase", s)
	}
}
