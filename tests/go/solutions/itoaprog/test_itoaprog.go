package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
)

func main() {
	for i := 0; i < 50; i++ {
		arg := z01.RandIntBetween(-2000000000, 2000000000)
		z01.Challenge("ItoaProg", Itoa, correct.Itoa, arg)
	}
}
