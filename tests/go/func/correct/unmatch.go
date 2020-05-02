package correct

// Returns the element of the slice that doesn't have a correspondant pair
func Unmatch(arr []int) int {
	var quant int
	for _, el := range arr {
		quant = 0
		for _, v := range arr {
			if v == el {
				quant++
			}
		}
		if quant%2 != 0 {
			return el
		}
	}
	return -1
}
