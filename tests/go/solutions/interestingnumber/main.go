package main

func isInteresting(n int) bool {
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	if s%7 == 0 {
		return true
	}
	return false
}

func InterestingNumber(n int) int {
	for n > 0 {
		if isInteresting(n) == true {
			return n
		}
		n++
	}
	return -1
}

func main() {

}
