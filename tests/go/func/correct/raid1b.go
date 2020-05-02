package correct

import "fmt"

func drawLineRaid1b(x int, s string) {
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

func Raid1b(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	strBeg := "/*\\"
	strMed := "* *"
	strEnd := "\\*/"

	if y >= 1 {
		drawLineRaid1b(x, strBeg)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLineRaid1b(x, strMed)
		}
	}
	if y > 1 {
		drawLineRaid1b(x, strEnd)
	}
}
