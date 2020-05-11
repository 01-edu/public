package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
)

func strLen(str string) int {
	len := 0

	strConverted := []rune(str)
	for i, _ := range strConverted {
		len = i + 1
	}
	return len
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
		lib.Challenge("StrLen", strLen, student.StrLen, s)
	}
}
