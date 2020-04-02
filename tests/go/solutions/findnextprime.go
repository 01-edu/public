package solutions

import (
	"math"
)

func FindNextPrime(value int) int {
	isPrime := func(value int) bool {
		if value < 2 {
			return false
		}
		limit := int(math.Floor(math.Sqrt(float64(value))))
		i := 2
		for i <= limit {
			if value%i == 0 {
				return false
			}
			i++
		}
		return true
	}

	if isPrime(value) {
		return value
	}
	return FindNextPrime(value + 1)
}
