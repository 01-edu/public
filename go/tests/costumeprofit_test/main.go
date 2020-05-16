package main

import (
	"strconv"

	"lib"
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
