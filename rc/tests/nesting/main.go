package main

import "fmt"

func myFunction(s string) bool {
	return s > "m"
}

func main() {
	// myFunction := func(s string) bool {
	// 	return s < "m"
	// }

	fmt.Printf("Does %s comes before \"m\" %v\n", "name", myFunction("name"))
	fmt.Printf("Does %s comes before \"m\" %v\n", "change", myFunction("change"))
}
