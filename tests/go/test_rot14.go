package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	type nodeTest struct {
		data []string
	}

	table := []nodeTest{}
	for i := 0; i < 5; i++ {
		val := nodeTest{
			data: z01.MultRandWords(),
		}
		table = append(table, val)
	}

	for _, arg := range table {
		for _, s := range arg.data {
			z01.Challenge("Rot14", correct.Rot14, student.Rot14, s)
		}
	}
}
