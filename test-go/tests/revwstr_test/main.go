package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func main() {
	table := []string{
		"",
		"abcdefghijklm",
		"the time of contempt precedes that of indifference",
		"he stared at the mountain",
		"qw qw e qwsa d",
	}

	// 3 valid random sentences with no spaces at the beginning nor the end and only one space for separator.
	for i := 0; i < 3; i++ {
		numberOfWords := lib.RandIntBetween(1, 6)
		sentence := lib.RandAlnum()
		for j := 0; j < numberOfWords; j++ {
			sentence += " " + lib.RandAlnum()
		}
		sentence += lib.RandAlnum()
		table = append(table, sentence)
	}

	for _, s := range table {
		lib.ChallengeMain("revwstr", s)
	}

	lib.ChallengeMain("revwstr")
	lib.ChallengeMain("revwstr", "1param", "2param", "3param", "4param")
}
