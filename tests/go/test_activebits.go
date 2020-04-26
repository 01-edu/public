package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	args := []int{z01.RandIntBetween(2, 20)}
	args = append(args, z01.MultRandIntBetween(2, 20)...)
	args = append(args, z01.MultRandIntBetween(2, 20)...)

	for _, v := range args {
		z01.Challenge("ActiveBits", student.ActiveBits, correct.ActiveBits, v)
	}
}
