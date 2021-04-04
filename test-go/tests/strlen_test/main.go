package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func strLen(s string) int {
	return len([]rune(s))
}

func main() {
	randomStringCharset := "a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ"

	table := []string{}
	for i := 0; i < 10; i++ {
		randomLenghtOfWord := lib.RandIntBetween(1, 20)
		randomStrRandomLenght := lib.RandStr(randomLenghtOfWord, randomStringCharset)
		table = append(table, randomStrRandomLenght)
	}
	table = append(table, "Héllo!")
	table = append(table, randomStringCharset)

	for _, s := range table {
		lib.Challenge("StrLen", student.StrLen, strLen, s)
	}
}

// TODO: refactor, simplify, no need for specific charset : check lib
