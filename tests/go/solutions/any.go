package solutions

func Any(f func(string) bool, arr []string) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}

	return false
}
