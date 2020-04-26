package correct

func BasicJoin(a []string) (b string) {
	for _, s := range a {
		b += s
	}
	return b
}
