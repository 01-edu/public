package solutions

//Join the elements of the slice using the sep as glue between each element
func Join(arstr []string, sep string) string {
	res := ""
	n := len(arstr)
	for i, a := range arstr {
		res += a
		if i < n-1 {
			res += sep
		}
	}
	return res
}
