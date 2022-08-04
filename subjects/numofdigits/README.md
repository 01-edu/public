## number of digits

### Instructions

Write a function that returns the number of digits in a positive integer `n`.
- if the number is negative returns `0`.

### Expected function

```go
func Numofdigits(num int) int {

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
	fmt.Println(Numofdigits(3))
	fmt.Println(Numofdigits(245))
	fmt.Println(Numofdigits(-1))
	fmt.Println(Numofdigits(885))
	fmt.Println(Numofdigits(8574))
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
