## pilot

### Instructions

Write a `go` file so that the following program compiles.

### Usage

```go
package main

import (
	"fmt"
	student ".."
)

func main() {
	var donnie student.Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = student.AIRCRAFT1

	fmt.Println(donnie)
}
```
