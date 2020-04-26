package main

func ForEach(f func(int), a []int) {
	for _, el := range a {
		f(el)
	}
}

func main() {
}
