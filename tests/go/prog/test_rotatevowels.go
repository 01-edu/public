package main

import (
	"strings"

	"../../lib"
)

func main() {
	Lower := lib.RuneRange('a', 'z')
	Upper := lib.RuneRange('A', 'Z')
	letters := Lower + Upper + " "
	var arr []string
	for i := 0; i < 10; i++ {
		str := lib.RandStr(lib.RandIntBetween(2, 20), letters)
		arr = append(arr, str)
	}
	arr = append(arr, "")

	for _, v := range arr {
		lib.ChallengeMain("rotatevowels", strings.Fields(v)...)
	}
	lib.ChallengeMain("rotatevowels", "Hello World")
	lib.ChallengeMain("rotatevowels", "HEllO World", "problem solved")
	lib.ChallengeMain("rotatevowels", "str", "shh", "psst")
	lib.ChallengeMain("rotatevowels", "happy thoughts", "good luck")
	lib.ChallengeMain("rotatevowels", "al's elEphAnt is overly underweight!")
	lib.ChallengeMain("rotatevowels", "aEi", "Ou")
}
