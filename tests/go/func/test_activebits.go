package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	args := []int{lib.RandIntBetween(2, 20)}
	args = append(args, lib.MultRandIntBetween(2, 20)...)
	args = append(args, lib.MultRandIntBetween(2, 20)...)

	for _, v := range args {
		lib.Challenge("ActiveBits", student.ActiveBits, correct.ActiveBits, v)
	}
}
