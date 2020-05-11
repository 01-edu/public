package main

import (
	"strconv"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func itoa(i int) string {
	return strconv.Itoa(i)
}

func main() {
	for i := 0; i < 50; i++ {
		arg := lib.RandIntBetween(-2000000000, 2000000000)
		lib.Challenge("Itoa", student.Itoa, itoa, arg)
	}
}
