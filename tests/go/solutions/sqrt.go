package solutions

import "math"

func Sqrt(value int) int {
	sr := math.Sqrt(float64(value))
	if math.Mod(sr, 1) == 0 {
		return int(sr)
	}
	return 0
}
