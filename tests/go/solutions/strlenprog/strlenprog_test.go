package main

import (
	"testing"

	solutions "../../solutions"
	"github.com/01-edu/z01"
)

func TestStrLenProg(t *testing.T) {

	randomStringCharset := "a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ"

	table := []string{}
	for i := 0; i < 10; i++ {
		randomLenghtOfWord := z01.RandIntBetween(1, 20)
		randomStrRandomLenght := z01.RandStr(randomLenghtOfWord, randomStringCharset)
		table = append(table, randomStrRandomLenght)
	}
	table = append(table, "Héllo!")
	table = append(table, randomStringCharset)

	for _, s := range table {
		z01.Challenge(t, StrLen, solutions.StrLen, s)
	}
}
