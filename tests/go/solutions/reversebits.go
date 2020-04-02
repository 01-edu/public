package solutions

import (
	"math/bits"
)

func ReverseBits(by byte) byte {
	return bits.Reverse8(uint8(by))
}
