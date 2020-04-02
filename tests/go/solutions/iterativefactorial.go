package solutions

import "math/bits"

func IterativeFactorial(nb int) int {
	limit := 12
	if bits.UintSize == 64 {
		limit = 20
	}
	if nb < 0 || nb > limit {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return nb * IterativeFactorial(nb-1)
}
