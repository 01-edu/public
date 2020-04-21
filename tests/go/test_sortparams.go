package main

import (
	"github.com/01-edu/z01"
)

func main() {
	args := []string{"1", "a", "2", "A", "3", "b", "4", "C"}
	z01.ChallengeMain("sortparams", args...)
}
