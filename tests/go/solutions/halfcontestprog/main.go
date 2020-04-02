package main

func Halfcontest(h1, m1, h2, m2 int) int {
	t1 := h1*60 + m1
	t2 := h2*60 + m2
	t2 = (t2 + t1) / 2
	h2 = t2 / 60
	m2 = t2 % 60
	return h2*100 + m2
}

func main() {

}
