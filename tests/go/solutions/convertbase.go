package solutions

func ConvertNbrBase(n int, base string) string {
	var result string
	length := len(base)

	for n >= length {
		result = string(base[(n%length)]) + result
		n = n / length
	}
	result = string(base[n]) + result

	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	resultIntermediary := AtoiBase(nbr, baseFrom)

	resultFinal := ConvertNbrBase(resultIntermediary, baseTo)

	return resultFinal
}
