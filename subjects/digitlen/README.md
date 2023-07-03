## digitlen

### Instructions

Write a function `DigitLen()` that takes two integers as arguments and returns the times the first `int` can be divided by the second until it reaches zero.

- The second `int` must be between **_2_** and **_36_**. If not, the function returns `-1`.
- If the first `int` is negative, reverse the sign and count the digits.

### Expected function

```go
func DigitLen(n, base int) int {

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
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
```

And its output:

```console
$ go run . | cat -e
3$
7$
2$
-1$
```
