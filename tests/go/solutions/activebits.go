package solutions

//Function that return the number of active bits in the number passed as the argument
func ActiveBits(n int) uint {
	total := 0
	for ; n > 1; n = n / 2 {
		total += n % 2
	}
	total += n

	return uint(total)
}
