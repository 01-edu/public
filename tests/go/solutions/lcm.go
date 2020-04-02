package solutions

//This solution is the comparing file of the staff
// Because the solution is a function,
//
//1) here the package is solutions
//2) it does not need an empty func main(){}
//3) its location is 1 level below the folder of the nauuo_test.go file

func gcd(first, second int) int {
	if (second == 0) {
		return (first)
	}
	return (gcd(second, first % second))
}

func Lcm(first, second int) int {
	return (first / gcd(second, first % second) * second)
}