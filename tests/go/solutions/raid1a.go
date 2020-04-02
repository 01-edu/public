package solutions

import "github.com/01-edu/z01"

func drawLine(x int, str string) {
	strConverted := []rune(str)
	beg := strConverted[0]
	med := strConverted[1]
	end := strConverted[2]
	if x >= 1 {
		z01.PrintRune(beg)
	}
	if x > 2 {
		for i := 0; i < (x - 2); i++ {
			z01.PrintRune(med)
		}
	}
	if x > 1 {
		z01.PrintRune(end)
	}
	z01.PrintRune('\n')
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
