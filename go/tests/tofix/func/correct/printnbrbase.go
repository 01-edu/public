package correct

import "fmt"

func PrintNbrBase(n int, base string) {
	if validBase(base) {
		length := len(base)
		sign := 1
		rbase := []rune(base)
		if n < 0 {
			fmt.Print("-")
			sign = -1
		}
		if n < length && n >= 0 {
			fmt.Printf("%c", rbase[n])
		} else {
			PrintNbrBase(sign*(n/length), base)
			fmt.Printf("%c", rbase[sign*(n%length)])
		}
	} else {
		fmt.Print("NV")
	}
}
