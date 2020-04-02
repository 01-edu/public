package main

import (
	student ".."
	"fmt"
)

func main() {
	var donnie student.Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = student.AIRCRAFT1

	fmt.Println(donnie)
}
