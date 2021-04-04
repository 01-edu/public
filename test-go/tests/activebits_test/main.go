package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

// Function that return the number of active bits in the number passed as the argument
func activeBits(n int) (total int) {
	for ; n > 1; n /= 2 {
		total += n % 2
	}
	total += n
	return
}

func main() {
	args := []int{lib.RandIntBetween(2, 20)}
	args = append(args, lib.MultRandIntBetween(2, 20)...)
	args = append(args, lib.MultRandIntBetween(2, 20)...)

	for _, v := range args {
		lib.Challenge("ActiveBits", student.ActiveBits, activeBits, v)
	}
}
