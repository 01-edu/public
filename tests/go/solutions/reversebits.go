package solutions

import "math/bits"

func ReverseBits(octet byte) byte {
	return bits.Reverse8(uint8(octet))
}
