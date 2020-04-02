package solutions

//This solution is the comparing file of the staff
// Because the solution is a function,
//
//1) here the package is solutions
//2) it does not need an empty func main(){}
//3) its location is 1 level below the folder of the nauuo_test.go file

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