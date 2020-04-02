package solutions

func StrLen(str string) int {
	len := 0

	strConverted := []rune(str)
	for i, _ := range strConverted {
		len = i + 1
	}
	return len
}
