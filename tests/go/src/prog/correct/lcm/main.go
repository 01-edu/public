package main

import "fmt"

func gcd(first, second int) int {
	if second == 0 {
		return first
	}
	return gcd(second, first%second)
}

func Lcm(first, second int) int {
	return first / gcd(first, second) * second
}

func main() {
	fmt.Println(Lcm(3, 0))
	fmt.Println(Lcm(4, 0))
}
