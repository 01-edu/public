package main

//This solution is the placeholder of the student solution
// for an exercise in the exam asking for a Function
//Remember the disclaimer!!!!
//1) here the package is main
//2) It does need an empty func main(){}

func Priorprime(x int) int {
	ans := 0
	ok := 0
	for i := 2; i < x; i++ {
		ok = 1
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				ok = 0
			}
		}
		if ok == 1 {
			ans += i
		}
	}
	return ans
}

func main() {

}
