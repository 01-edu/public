package main

import (
	"strconv"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	for i := 0; i < 25; i++ {
		lib.ChallengeMain("costumeprofit",
			strconv.Itoa(lib.RandIntBetween(0, 1000)),
			strconv.Itoa(lib.RandIntBetween(0, 1000)),
			strconv.Itoa(lib.RandIntBetween(0, 1000)),
			strconv.Itoa(lib.RandIntBetween(0, 1000)),
			strconv.Itoa(lib.RandIntBetween(0, 1000)),
			strconv.Itoa(lib.RandIntBetween(0, 1000)),
		)
	}
}
