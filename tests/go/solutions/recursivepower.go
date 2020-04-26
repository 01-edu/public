package solutions

import "math"

func RecursivePower(nb, power int) int {
	if power < 0 {
		return 0
	}
	return int(math.Pow(float64(nb), float64(power)))
}
