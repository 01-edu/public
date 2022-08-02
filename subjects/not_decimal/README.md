## Not_decimal

### Instructions

Write a function called `Not_decimal()` that takes a string in forme of a float number with the decimal point and returns an integer without the decimal point (you muliply by 10^n to remove `.`).

- If the number is doesn'the have a decimal point or  there is only zero after `.` return the number followed by a newline (`'\n'`).
- If the argument is empty return newline (`'\n'`).
- If The argument is not a number return it followed by newline (`'\n'`).


### Expected function

```go
func Not_decimal(dec string) int {
}
```
### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Print(Not_decimal("0.1"))
	fmt.Print(Not_decimal("174.2"))
	fmt.Print(Not_decimal("0.1255"))
	fmt.Print(Not_decimal("1.20525856"))
	fmt.Print(Not_decimal("-0.0f00d00"))
	fmt.Print(Not_decimal(""))
	fmt.Print(Not_decimal("-19.525856"))
	fmt.Print(Not_decimal("1952"))
}
```
And its output :

```go
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