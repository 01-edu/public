package correct

import "fmt"

func drawLineRaid1e(x int, s string) {
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

func Raid1e(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	strBeg := "ABC"
	strMed := "B B"
	strEnd := "CBA"

	if y >= 1 {
		drawLineRaid1e(x, strBeg)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLineRaid1e(x, strMed)
		}
	}
	if y > 1 {
		drawLineRaid1e(x, strEnd)
	}
}
