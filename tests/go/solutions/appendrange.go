package solutions

func AppendRange(min, max int) []int {
	size := max - min
	answer := []int{}

	if size <= 0 {
		return nil
	}

	for i := min; i < max; i++ {
		answer = append(answer, i)
	}

	return answer
}
