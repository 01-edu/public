## descendappendrange

### Instructions

- Write a function that takes an `int` max and an `int` min as parameters. The function should return a slice of `int`s with all the values between the max and min.

- The `max` must be included, and `min` must be excluded.

- If `max` is inferior than or equal to `min`, return an empty slice.

- `make()` is not allowed for this exercise.

### Expected function

```go
func DescendAppendRange(max, min int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.DescendAppendRange(10, 5))
	fmt.Println(piscine.DescendAppendRange(5, 10))
}
```

And its output:

```console
$ go run . | cat -e
[10 9 8 7 6]
[]
$
```
