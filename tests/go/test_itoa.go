package main

import (
	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	for i := 0; i < 50; i++ {
		arg := z01.RandIntBetween(-2000000000, 2000000000)
		z01.Challenge("Itoa", student.Itoa, solutions.Itoa, arg)
	}
}
