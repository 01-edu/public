package main

import "fmt"

func main() {

	for a := 9; a >= 2; {
		for b := a - 1; b >= 1; {
			for c := b - 1; c >= 0; {

				if a > b && b > c && (a+b+c) != 3 {
					fmt.Printf("%d%d%d, ", a, b, c)
				}
				if (a + b + c) == 3 {
					fmt.Printf("%d%d%d\n", a, b, c)
				}

				c--
			}
			b--
		}
		a--
	}

}
