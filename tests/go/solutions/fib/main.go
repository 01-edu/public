package main

func Fib(n int) int {
	if n <= 0 {
		return 0
	}
	t1 := 0
	t2 := 1
	for i := 2; i <= n; i++ {
		t1 += t2
		tmp := t1
		t1 = t2
		t2 = tmp
	}
	return t2
}

func main() {
}
