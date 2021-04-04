package main

import (
	"strings"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	letters := lib.Lower + lib.Upper + " "
	a := []string{""}
	for i := 0; i < 10; i++ {
		a = append(a, lib.RandStr(lib.RandIntBetween(2, 20), letters))
	}

	for _, v := range a {
		lib.ChallengeMain("rotatevowels", strings.Fields(v)...)
	}
	lib.ChallengeMain("rotatevowels", "Hello World")
	lib.ChallengeMain("rotatevowels", "HEllO World", "problem solved")
	lib.ChallengeMain("rotatevowels", "str", "shh", "psst")
	lib.ChallengeMain("rotatevowels", "happy thoughts", "good luck")
	lib.ChallengeMain("rotatevowels", "al's elEphAnt is overly underweight!")
	lib.ChallengeMain("rotatevowels", "aEi", "Ou")
}
