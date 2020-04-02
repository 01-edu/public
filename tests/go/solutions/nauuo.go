package solutions

//This solution is the comparing file of the staff
// Because the solution is a function,
//
//1) here the package is solutions
//2) it does not need an empty func main(){}
//3) its location is 1 level below the folder of the nauuo_test.go file

func Nauuo(plus, minus, rand int) string {
	if rand == 0 {
		if plus > minus {
			return ("+")
		}
		if plus < minus {
			return ("-")
		}
		if plus == minus {
			return ("0")
		}
	}
	if plus > minus+rand {
		return ("+")
	}
	if plus+rand < minus {
		return ("-")
	}
	if (plus+rand >= minus) && (plus-rand <= minus) {
		return ("?")
	}
	if (minus+rand >= plus) && (minus-rand <= plus) {
		return ("?")
	}
	return ("?")
}
