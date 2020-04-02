package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestIndex(t *testing.T) {
	type node struct {
		s      string
		toFind string
	}

	table := []node{}

	// the first 15 values are valid for this test
	for i := 0; i < 15; i++ {
		wordToTest := z01.RandASCII()
		firstLetterIndex := z01.RandIntBetween(0, (len(wordToTest)-1)/2)
		lastLetterIndex := z01.RandIntBetween(firstLetterIndex, len(wordToTest)-1)

		val := node{
			s:      wordToTest,
			toFind: wordToTest[firstLetterIndex:lastLetterIndex],
		}
		table = append(table, val)
	}

	// the next 15 values are supposed to be invalid for this test
	for i := 0; i < 15; i++ {
		wordToTest := z01.RandASCII()
		wrongMatch := z01.RandASCII()

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
		z01.Challenge(t, student.Index, solutions.Index, arg.s, arg.toFind)
	}
}
