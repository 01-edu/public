package main

import "fmt"

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

const Aircraft1 = 1

func main() {
	donnie := Pilot{
		Name:     "Donnie",
		Life:     100.0,
		Age:      24,
		Aircraft: Aircraft1,
	}
	fmt.Println(donnie)
}
