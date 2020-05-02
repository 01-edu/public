package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	for i := 0; i < 50; i++ {
		arg := lib.RandIntBetween(-2000000000, 2000000000)
		lib.Challenge("Itoa", student.Itoa, correct.Itoa, arg)
	}
}
