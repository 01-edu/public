package solutions

func FindPrevPrime(nbr int) int {
	if nbr < 2 {
		return 0
	}
	if IsPrime(nbr) {
		return nbr
	}
	return FindPrevPrime(nbr - 1)

}
