package correct

func PriorPrime(x int) int {
	ans := 0
	ok := 0
	for i := 2; i < x; i++ {
		ok = 1
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				ok = 0
			}
		}
		if ok == 1 {
			ans += i
		}
	}
	return ans
}
