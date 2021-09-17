## pilot

### Instructions

- Create a directory called `pilot`.
- Inside the directory `pilot` create a file `main.go`.
- Copy the code below to `main.go` and add the code needed so that the program compiles.

> Note: You can only add code, not delete.

### Usage

```go
package main

import "fmt"

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

const AIRCRAFT1 = 1
```
