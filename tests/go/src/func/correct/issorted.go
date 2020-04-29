package correct

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
