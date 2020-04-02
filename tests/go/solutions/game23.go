package solutions

func Game23(a, b int) int {
	if a > b {
		return -1
	}
	if a == b {
		// fmt.Printf("%d\n", cnt)
		return 0
	}
	if Game23(a*2, b) != -1 {
		return 1 + Game23(a*2, b)
	}
	if Game23(a*3, b) != -1 {
		return 1 + Game23(a*3, b)
	}
	return -1
}
