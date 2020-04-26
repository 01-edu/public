package solutions

import "math"

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	result := math.Pow(float64(nb), float64(power))
	return int(result)
}
