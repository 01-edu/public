package student_test

import (
	"math/rand"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestNRune(t *testing.T) {
	type node struct {
		word string
		n    int
	}

	table := []node{}

	for i := 0; i < 30; i++ {
		wordToInput := z01.RandASCII()
		val := node{
			word: wordToInput,
			n:    rand.Intn(len(wordToInput)) + 1,
		}
		table = append(table, val)
	}

	table = append(table,
		node{word: "Hello!", n: 3},
		node{word: "Salut!", n: 2},
		node{word: "Ola!", n: 4},
		node{word: "♥01!", n: 1},
		node{word: "Not", n: 6},
		node{word: "Also not", n: 9},
		node{word: "Tree house", n: 5},
		node{word: "Whatever", n: 0},
		node{word: "Whatshisname", n: -2},
	)

	for _, arg := range table {
		z01.Challenge(t, student.NRune, solutions.NRune, arg.word, arg.n)
	}
}
