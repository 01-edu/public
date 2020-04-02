package main

func Reachablenumber(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		if n < 10 {
			cnt += 8
			break
		} else {
			n++
		}
		for n%10 == 0 {
			n /= 10
		}
	}
	return cnt
}

func main() {

}
