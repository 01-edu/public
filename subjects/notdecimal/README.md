## notdecimal

### Instructions

Write a function called `NotDecimal()` that takes as an argument a `string` in form of a float number with the decimal point and returns a string converted to `int` without the decimal point (you will have to multiply it by 10^n to remove the `.`).

- If the number doesn't have a decimal point or there is only a zero after the `.` return the number followed by a newline `\n`.
- If the argument is empty return a newline `\n`.
- If the argument is not a number return it followed by a newline `\n`.

### Expected function

```go
func NotDecimal(dec string) string {

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
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
```

And its output:

```console
$ go run .  | cat -e
1$
1742$
1255$
120525856$
-0.0f00d00$
$
-19525856$
1952$
```
