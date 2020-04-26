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

func FindPrevPrime(nbr int) int {
	if nbr < 2 {
		return 0
	}
	if isPrime(nbr) {
		return nbr
	}
	return FindPrevPrime(nbr - 1)
}
