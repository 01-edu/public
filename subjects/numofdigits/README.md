## number of digits

### Instructions

Write a function that returns the number of digits in a nonnegative integer n.
- if the number is negative returns 0.

### Expected function

```go
func Numofdigits(num int) int {

}
```

### Usage

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Numofdigits(3))
	fmt.Println(piscine.Numofdigits(245))
	fmt.Println(piscine.Numofdigits(-1))
	fmt.Println(piscine.Numofdigits(885))
	fmt.Println(piscine.Numofdigits(8574))
}
```

And its output :

```console
$ go run .
1
3
0
3
4
