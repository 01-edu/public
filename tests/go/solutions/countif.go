package solutions

func CountIf(f func(string) bool, arr []string) int {

	counter := 0
	for _, el := range arr {
		if f(el) {
			counter++
		}
	}

	return counter

}
