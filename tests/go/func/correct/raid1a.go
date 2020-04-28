package correct

import "fmt"

func drawLine(x int, s string) {
	beg := s[0]
	med := s[1]
	end := s[2]
	if x >= 1 {
		fmt.Printf("%c", beg)
	}
	if x > 2 {
		for i := 0; i < (x - 2); i++ {
			fmt.Printf("%c", med)
		}
	}
	if x > 1 {
		fmt.Printf("%c", end)
	}
	fmt.Println()
}

func printTheLines(x, y int, strBeg, strMed, strEnd string) {
	if y >= 1 {
		drawLine(x, strBeg)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLine(x, strMed)
		}
	}
	if y > 1 {
		drawLine(x, strEnd)
	}
}

func Raid1a(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	printTheLines(x, y, "o-o", "| |", "o-o")
}
