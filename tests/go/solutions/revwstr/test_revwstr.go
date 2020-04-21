package main

import (
	"github.com/01-edu/z01"
)

func main() {
	table := []string{}

	table = append(table,
		"",
		"abcdefghijklm",
		"the time of contempt precedes that of indifference",
		"he stared at the mountain",
		"qw qw e qwsa d")

	//3 valid random sentences with no spaces at the beginning nor the end and only one space for separator.
	for i := 0; i < 3; i++ {
		numberOfWords := z01.RandIntBetween(1, 6)
		sentence := z01.RandAlnum()
		for j := 0; j < numberOfWords; j++ {
			sentence = sentence + " " + z01.RandAlnum()
		}
		sentence = sentence + z01.RandAlnum()
		table = append(table, sentence)
	}

	for _, s := range table {
		z01.ChallengeMain("revwstr", s)
	}

	z01.ChallengeMain("revwstr")
	z01.ChallengeMain("revwstr", "1param", "2param", "3param", "4param")
}
