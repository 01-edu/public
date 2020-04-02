package solutions

func Compact(slice *[]string) int {
	count := 0
	var compacted []string
	for _, v := range *slice {
		if v != "" {
			count++
			compacted = append(compacted, v)
		}
	}
	*slice = compacted
	return count
}
