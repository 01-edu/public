package main

import "fmt"

//This solution is the placeholder of the student solution
// for an exercise in the exam asking for a Function
//Remember the disclaimer!!!!
//1) here the package is main
//2) It does need an empty func main(){}

func gcd(first, second int) int {
	if second == 0 {
		return (first)
	}
	return gcd(second, first%second)
}

func Lcm(first, second int) int {
	return (first / gcd(first, second) * second)
}
func main() {
	fmt.Println(Lcm(3, 0))
	fmt.Println(Lcm(4, 0))
}
