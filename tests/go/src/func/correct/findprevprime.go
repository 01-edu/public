package correct

func isPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	}

	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if isPrime(nb) {
		return nb
	}
	return FindPrevPrime(nb - 1)
}
