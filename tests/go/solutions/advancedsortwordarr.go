package solutions

func CompArray(a, b string) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	} else {
		return 1
	}
}

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 1; i < len(array); i++ {
		if f(array[i], array[i-1]) < 0 {
			array[i], array[i-1] = array[i-1], array[i]
			i = 0
		}

	}

}
