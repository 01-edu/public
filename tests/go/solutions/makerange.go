package solutions

func MakeRange(min, max int) []int {
	size := max - min

	if size <= 0 {
		return nil
	}
	answer := make([]int, size)
	for i := range answer {
		answer[i] = min
		min++
	}
	return answer
}
