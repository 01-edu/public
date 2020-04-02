package solutions

func SplitWhiteSpaces(str string) []string {

	answer := []string{}
	word := ""
	for _, r := range str {
		if isSeparator(r) {
			if word != "" {
				answer = append(answer, word)
				word = ""
			}
		} else {
			word = word + string(r)
		}
	}
	if word != "" {
		answer = append(answer, word)
	}
	return answer
}
