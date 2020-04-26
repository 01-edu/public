package solutions

func gcd(first, second int) int {
	if second == 0 {
		return first
	}
	return gcd(second, first%second)
}

func Lcm(first, second int) int {
	return first / gcd(second, first%second) * second
}
