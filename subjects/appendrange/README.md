## appendrange

### Instructions

Write a function that takes an `int` min and an `int` max as parameters. The function should return a slice of `int`s with all the values between the min and max.

Min is included, and max is excluded.

If min is greater than or equal to max, a `nil` slice is returned.

`make` is not allowed for this exercise.

### Expected function

```go
func AppendRange(min, max int) []int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(5, 10))
	fmt.Println(piscine.AppendRange(10, 5))
}
```

And its output :

```console
$ go run .
[5 6 7 8 9]
[]
$
```
