package solutions

import (
	"fmt"
)

func ForEach(f func(int), arr []int) {

	for _, el := range arr {
		f(el)
	}

}

func Add0(nbr int) {

	fmt.Println(nbr)
}

func Add1(nbr int) {

	fmt.Println(nbr + 1)
}

func Add2(nbr int) {

	fmt.Println(nbr + 2)
}

func Add3(nbr int) {

	fmt.Println(nbr + 3)
}
