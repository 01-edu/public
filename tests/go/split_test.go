package student_test

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestSplit(t *testing.T) {

	separators := []string{"!=HA=!",
		"!==!",
		"    ",
		"|=choumi=|",
		"|<=>|",
		z01.RandStr(3, z01.RuneRange('A', 'Z')),
		"<<==123==>>",
		"[<>abc<>]"}

	type node struct {
		str string
		sep string
	}
	table := []node{}
	//15 random slice of strings

	for i := 0; i < 15; i++ {
		separator := separators[rand.Intn(len(separators))]
		val := node{
			str: strings.Join(z01.MultRandAlnum(), separator),
			sep: separator,
		}
		table = append(table, val)
	}

	table = append(table,
		node{str: "HelloHAhowHAareHAyou?", sep: "HA"})

	for _, arg := range table {
		z01.Challenge(t, student.Split, solutions.Split, arg.str, arg.sep)
	}
}
