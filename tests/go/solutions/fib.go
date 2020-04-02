package solutions

func rec(a, b, cnt int) int {
	if a > b {
		return -1
	}
	if a == b {
		// fmt.Printf("%d\n", cnt)
		return cnt
	}
	if rec(a*2, b, cnt+1) != -1 {
		return rec(a*2, b, cnt+1)
	}
	if rec(a*3, b, cnt+1) != -1 {
		return rec(a*3, b, cnt+1)
	}
	return -1
}

func Fib(n int) int {
	if n <= 0 {
		return 0
	}
	t1 := 0
	t2 := 1
	for i := 2; i <= n; i++ {
		t1 = t1 + t2
		tmp := t1
		t1 = t2
		t2 = tmp
	}
	return t2
}
