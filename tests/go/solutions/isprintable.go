package solutions

func IsPrintable(s string) bool {
	for _, r := range s {
		if !(r >= 32 && r <= 127) {
			return false
		}
	}
	return true
}
