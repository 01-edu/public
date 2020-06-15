// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// level 3: doopprog

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 4 {
		return
	}

	if args[2] != "+" && args[2] != "-" && args[2] != "/" && args[2] != "*" && args[2] != "%" {
		fmt.Println(0)
		return
	}

	// arg1 := getArg(args[1])

	arg1, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(0)
		return
	}
	//fmt.Println("Arg1: ", arg1)

	arg2, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println(0)
		return
	}
	//fmt.Println("Arg2: ", arg2)

	if args[2] == "+" {
		Add(int64(arg1), int64(arg2))
	} else if args[2] == "-" {
		Subs(int64(arg1), int64(arg2))
	} else if args[2] == "/" {
		Div(int64(arg1), int64(arg2))
	} else if args[2] == "*" {
		Multip(int64(arg1), int64(arg2))
	} else if args[2] == "%" {
		Modulo(int64(arg1), int64(arg2))
	}

}

var max int64 = 9223372036854775807
var min int64 = -9223372036854775808

func Add(x, y int64) {
	if x >= 0 && y < 0 {
		Subs(x, -y)
		return
	} else if x < 0 && y >= 0 {
		Subs(y, -x)
		return
	} else if x < 0 && y < 0 {
		if x >= min-y {
			fmt.Println(x + y)
			return
		} else {
			fmt.Println(0)
			return
		}
	}

	if x <= max-y {
		fmt.Println(x + y)
	} else {
		fmt.Println(0)
		return
	}
}

func Subs(x, y int64) {
	if x >= 0 && y < 0 {
		Add(x, -y)
		return
	} else if x < 0 && y >= 0 {
		Add(-y, x)
		return
	} else if x < 0 && y < 0 {
		Subs(-y, -x)
		return
	}

	if x >= min+y {
		fmt.Println(x - y)
		return
	} else {
		fmt.Println(0)
		return
	}
}

func Multip(x, y int64) {
	if x >= 0 && y >= 0 {
		if x <= max/y {
			fmt.Println(x * y)
			return
		} else {
			fmt.Println(0)
			return
		}
	} else if x < 0 && y < 0 {
		if x > max/y {
			fmt.Println(x * y)
			return
		} else {
			fmt.Println(0)
			return
		}
	} else {
		if x >= min/y {
			fmt.Println(x * y)
			return
		} else {
			fmt.Println(0)
			return
		}
	}
}

func Div(x, y int64) {
	if y == 0 {
		fmt.Println("No division by 0")
		return
	}
	if x == min && y == -1 {
		fmt.Println(0)
		return
	}
	fmt.Println(x / y)
	return
}

func Modulo(x, y int64) {
	if y == 0 {
		fmt.Println("No modulo by 0")
		return
	}
	if y == min && x == -1 {
		fmt.Println(0)
		return
	}
	fmt.Println(x % y)
	return
}

//9223372036854775807
// func main() {
// 	if len(os.Args) == 4 {
// 		var result int
// 		firstArg, err := strconv.Atoi(os.Args[1])

// 		if err != nil {
// 			fmt.Println(0)
// 			return
// 		}

// 		operator := os.Args[2]
// 		secondArg, err1 := strconv.Atoi(os.Args[3])

// 		if err1 != nil {
// 			fmt.Println(0)
// 			return
// 		}

// 		if secondArg == 0 && operator == "/" {
// 			fmt.Println("No division by 0")
// 			return
// 		} else if secondArg == 0 && operator == "%" {
// 			fmt.Println("No modulo by 0")
// 			return
// 		} else if operator == "+" {
// 			result = firstArg + secondArg
// 			if !((result > firstArg) == (secondArg > 0)) {
// 				fmt.Println(0)
// 				return
// 			}
// 		} else if operator == "-" {
// 			result = firstArg - secondArg
// 			if !((result < firstArg) == (secondArg > 0)) {
// 				fmt.Println(0)
// 				return
// 			}
// 		} else if operator == "/" {
// 			result = firstArg / secondArg
// 		} else if operator == "*" {
// 			result = firstArg * secondArg
// 			if firstArg != 0 && (result/firstArg != secondArg) {
// 				fmt.Println(0)
// 				return
// 			}
// 		} else if operator == "%" {
// 			result = firstArg % secondArg
// 		}
// 		fmt.Println(result)
// 	}
// }
