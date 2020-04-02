package solutions

func IsPositive(value int) bool {
	return value > 0
}

func IsNegative0(value int) bool {
	return value < 0
}

func Map(f func(int) bool, arr []int) []bool {

	arrBool := make([]bool, len(arr))

	for i, el := range arr {
		arrBool[i] = f(el)
	}

	return arrBool

}
