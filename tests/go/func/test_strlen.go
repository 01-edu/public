package main

import (
	"../lib"
	"./correct"
	"./student"
)

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
		lib.Challenge("StrLen", correct.StrLen, student.StrLen, s)
	}
}
