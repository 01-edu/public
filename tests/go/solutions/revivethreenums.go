package solutions

func more(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func max(a, b, c, d int) int {
	if a >= b && a >= c && a >= d {
		return a
	}
	if b >= a && b >= c && b >= d {
		return b
	}
	if c >= a && c >= b && c >= d {
		return c
	}
	if d >= a && d >= b && d >= c {
		return d
	}
	return -1
}

func Revive_three_nums(a, b, c, d int) int {
	maxi := -111
	if a != max(a, b, c, d) {
		maxi = more(maxi, max(a, b, c, d)-a)
	}
	if b != max(a, b, c, d) {
		maxi = more(maxi, max(a, b, c, d)-b)
	}
	if c != max(a, b, c, d) {
		maxi = more(maxi, max(a, b, c, d)-c)
	}
	if d != max(a, b, c, d) {
		maxi = more(maxi, max(a, b, c, d)-d)
	}
	return maxi
}
