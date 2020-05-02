package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	type nodeTest struct {
		data []string
	}

	table := []nodeTest{}
	for i := 0; i < 5; i++ {
		val := nodeTest{
			data: lib.MultRandWords(),
		}
		table = append(table, val)
	}

	for _, arg := range table {
		for _, s := range arg.data {
			lib.Challenge("Rot14", correct.Rot14, student.Rot14, s)
		}
	}
}
