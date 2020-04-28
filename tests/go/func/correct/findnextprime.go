package correct

func FindNextPrime(nb int) int {
	if isPrime(nb) {
		return nb
	}
	return FindNextPrime(nb + 1)
}
