package main

import (
	"math/bits"
)

func ReverseBits(by byte) byte {
	return bits.Reverse8(uint8(by))
}

func main() {
	// this main is for testing purposes only, uncomment if needed
	// a := byte(0x26)
	// fmt.Printf("%08b\n", a)
	// fmt.Printf("%08b\n", ReverseBits(a))
}
