package main

//This solution is the placeholder of the student solution
// for an exercise in the exam asking for a Function
//Remember the disclaimer!!!!
//1) here the package is main
//2) It does need an empty func main(){}

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

func main() {

}
