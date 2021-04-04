package main

import (
	"strconv"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	rand := []string{
		"0",
		"4000",
		"5000",
		"12433",
		"hello",
		"good luck",
	}
	for i := 0; i < 7; i++ {
		rand = append(rand, strconv.Itoa(lib.RandIntBetween(0, 4000)))
	}
	for _, v := range rand {
		lib.ChallengeMain("romannumbers", v)
	}
}
