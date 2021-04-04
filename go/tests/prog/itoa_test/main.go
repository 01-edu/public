package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	for i := 0; i < 50; i++ {
		lib.Challenge("Itoa", student.Itoa, strconv.Itoa, lib.RandInt())
	}
}
