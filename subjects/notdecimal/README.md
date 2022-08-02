## Not_decimal

### Instructions

Write a function called `Not_decimal()` that takes a float number with the decimal point and return an integer with the decimal point (you muliply by 10^n to remove `.`).

- if the number is doesn'the have a decimal point or  there is only zero after `.` return the number.


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
	fmt.Println(Not_decimal(0.1))
	fmt.Println(Not_decimal(1.2))
	fmt.Println(Not_decimal(0.15))
	fmt.Println(Not_decimal(1.256))
	fmt.Println(Not_decimal(0.00009))
}
```
And its output :

```go
$ go run . 
1
12
15
1256
9
```