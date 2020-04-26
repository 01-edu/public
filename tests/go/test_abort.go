package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	arg := z01.MultRandInt()
	arg = append(arg, z01.RandInt())
	for i := 0; i < 15; i++ {
		z01.Challenge("Abort", student.Abort, correct.Abort, arg[0], arg[1], arg[2], arg[3], arg[4])
		arg = z01.MultRandInt()
		arg = append(arg, z01.RandInt())
	}
}
