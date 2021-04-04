package main

import (
	"strings"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func index(s string, substr string) int {
	return strings.Index(s, substr)
}

func main() {
	type node struct {
		s      string
		toFind string
	}

	table := []node{}

	// the first 15 values are valid for this test
	for i := 0; i < 15; i++ {
		wordToTest := lib.RandASCII()
		firstLetterIndex := lib.RandIntBetween(0, (len(wordToTest)-1)/2)
		lastLetterIndex := lib.RandIntBetween(firstLetterIndex, len(wordToTest)-1)

		val := node{
			s:      wordToTest,
			toFind: wordToTest[firstLetterIndex:lastLetterIndex],
		}
		table = append(table, val)
	}

	// the next 15 values are supposed to be invalid for this test
	for i := 0; i < 15; i++ {
		wordToTest := lib.RandASCII()
		wrongMatch := lib.RandASCII()

		val := node{
			s:      wordToTest,
			toFind: wrongMatch,
		}
		table = append(table, val)
	}
	// those are the test values from the README examples
	table = append(table,
		node{s: "Hello!", toFind: "l"},
		node{s: "Salut!", toFind: "alu"},
		node{s: "Ola!", toFind: "hOl"},
	)

	for _, arg := range table {
		lib.Challenge("Index", student.Index, index, arg.s, arg.toFind)
	}
}
