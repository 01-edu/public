package main

func main(){

}

func FindPrevPrime(nbr int) int {
	if nbr < 2 {
		return 0
	}
	if IsPrime(nbr) {
		return nbr
	}
	return FindPrevPrime(nbr - 1)

}

func IsPrime(nb int) bool {

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
