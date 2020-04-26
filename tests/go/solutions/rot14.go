package solutions

func Rot14(s string) (result string) {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			if r >= 'm' {
				r -= 12
			} else {
				r += 14
			}
		} else if r >= 'A' && r <= 'Z' {
			if r >= 'M' {
				r -= 12
			} else {
				r += 14
			}
		}
		result += string(r)
	}
	return result
}
