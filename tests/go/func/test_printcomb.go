package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	z01.Challenge("PrintComb", student.PrintComb, correct.PrintComb)
}
