package main

import (
	"strconv"

	student "student"

	"lib"
)

func main() {
	for i := 0; i < 50; i++ {
		lib.Challenge("Itoa", student.Itoa, strconv.Itoa, lib.RandInt())
	}
}
