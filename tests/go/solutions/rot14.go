package solutions

func Rot14(str string) string {
	arrayRune := []rune(str)
	var result string

	for i := 0; i < len(arrayRune); i++ {
		if arrayRune[i] >= 'a' && arrayRune[i] <= 'z' {
			if arrayRune[i] >= 'm' {
				arrayRune[i] = arrayRune[i] - 12
			} else {
				arrayRune[i] = arrayRune[i] + 14
			}
		} else if arrayRune[i] >= 'A' && arrayRune[i] <= 'Z' {
			if arrayRune[i] >= 'M' {
				arrayRune[i] = arrayRune[i] - 12
			} else {
				arrayRune[i] = arrayRune[i] + 14
			}
		}
		result += string(arrayRune[i])
	}
	return result
}
