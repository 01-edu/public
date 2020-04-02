package solutions

func IsSortedBy1(a, b int) int {
	if a-b < 0 {
		return -1
	} else if a-b > 0 {
		return 1
	} else {
		return 0
	}
}

func IsSortedBy10(a, b int) int {
	if a-b < 0 {
		return -10
	} else if a-b > 0 {
		return 10
	} else {
		return 0
	}
}

func IsSortedByDiff(a, b int) int {
	return a - b
}

func IsSorted(f func(int, int) int, arr []int) bool {

	ascendingOrdered := true
	descendingOrdered := true

	for i := 1; i < len(arr); i++ {
		if !(f(arr[i-1], arr[i]) >= 0) {
			ascendingOrdered = false
		}
	}

	for i := 1; i < len(arr); i++ {
		if !(f(arr[i-1], arr[i]) <= 0) {
			descendingOrdered = false
		}
	}

	return ascendingOrdered || descendingOrdered

}
