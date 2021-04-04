package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func rot14(s string) (result string) {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			if r >= 'm' {
				r -= 12
			} else {
				r += 14
			}
		} else if r >= 'A' && r <= 'Z' {
			if r >= 'M' {
				r -= 12
			} else {
				r += 14
			}
		}
		result += string(r)
	}
	return result
}

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
			lib.Challenge("Rot14", rot14, student.Rot14, s)
		}
	}
}
