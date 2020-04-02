package solutions

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Volumechanger(a, b int) int {
	return abs(a-b)/5 + abs(a-b)%5/2 + abs(a-b)%5%2
}
