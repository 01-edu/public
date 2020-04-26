package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	arg := lib.MultRandInt()
	arg = append(arg, lib.RandInt())
	for i := 0; i < 15; i++ {
		lib.Challenge("Abort", student.Abort, correct.Abort, arg[0], arg[1], arg[2], arg[3], arg[4])
		arg = lib.MultRandInt()
		arg = append(arg, lib.RandInt())
	}
}
