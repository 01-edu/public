package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := []string{
		" ",
		"Hello 78 World!    4455 /",
	}
	for l := 0; l < 7; l++ {
		a := z01.RandIntBetween(5, 20)
		b := z01.RandASCII()
		table = append(table, z01.RandStr(a, b))
	}

	for _, arg := range table {
		z01.Challenge("AlphaCount", student.AlphaCount, correct.AlphaCount, arg)
	}
}
