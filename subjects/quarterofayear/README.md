## quarterofayear

### Instructions

Write a function `QuarterOfAYear()` that takes an `int`, from 1 to 12, as an argument and returns an `int` with the respective quarter of the year to which it belongs.

- If the number is not between 1 and 12 return `-1`.

### Expected function

```go
func QuarterOfAYear(month int) int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(QuarterOfAYear(2))
	fmt.Println(QuarterOfAYear(5))
	fmt.Println(QuarterOfAYear(9))
	fmt.Println(QuarterOfAYear(11))
	fmt.Println(QuarterOfAYear(13))
	fmt.Println(QuarterOfAYear(-5))
}

```

And its output:

```console
$ go run .
1
2
3
4
-1
-1
```
