package correct

func SwapBits(n byte) byte {
	return (n >> 4) | (n << 4)
}
