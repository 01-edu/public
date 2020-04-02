package main

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Sweetproblem(a, b, c int) int {
	if a > b {
		f := a
		a = b
		b = f
	}
	if a > c {
		f := a
		a = c
		c = f
	}
	if b > c {
		f := b
		b = c
		c = f
	}
	ans := a
	if c-b >= a {
		c -= a
	} else {
		a -= c - b
		half := a / 2
		c -= half
		b -= a - half
	}
	ans += min2(b, c)
	return ans
}

func main() {

}
