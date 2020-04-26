package solutions

func isInteresting(n int) bool {
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s%7 == 0
}

func InterestingNumber(n int) int {
	for {
		if isInteresting(n) {
			return n
		}
		n++
	}
}
