package main

func main() {
}

func Slice(arr []string, nbr ...int) []string {

	if len(nbr) == 0 {
		return arr
	}

	first := nbr[0]
	if len(nbr) == 1 {
		if first < 0 {
			first = len(arr) + first
			if first < 0 {
				return arr
			}
		}
		return arr[first:]
	} else {
		second := nbr[1]

		first = ifNegative(arr, first)
		second = ifNegative(arr, second)

		if first > second {
			return []string{}
		}

		return arr[first:second]
	}
}

func ifNegative(arr []string, n int) int {
	if n < 0 {
		n = len(arr) + n
	}

	if n < 0 {
		n = 0
	} else if n > len(arr) {
		n = len(arr)
	}

	return n
}
