package solutions

func TrimAtoi(str string) int {
	chars := []rune(str)
	for i := range str {
		chars[i] = rune(str[i])
	}
	var numbers []rune
	for i := range chars {
		if (chars[i] >= '0' && chars[i] <= '9') || (len(numbers) == 0 && (chars[i]) == '-' || chars[i] == '+') {
			numbers = append(numbers, chars[i])
		}
	}
	if len(numbers) == 0 || (len(numbers) == 1 && (numbers[0] == '-' || numbers[0] == '+')) {
		return 0
	}

	res, i, sign := 0, 0, 1

	if numbers[0] == '-' {
		sign = -1
		i++
	} else if numbers[0] == '+' {
		i++
	}
	for ; i < len(numbers); i++ {
		res = res*10 + int(numbers[i]) - '0'
	}

	return sign * res
}
