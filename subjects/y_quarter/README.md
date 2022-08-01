## Y_quarter

### Instructions

Write a function called `Y_quarter()` that takes  month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.

- if the number is not between 1 and 12 return `-1`.


### Expected function

```go
func Y_quarter(month int) int {
}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(Y_quarter(2))
	fmt.Println(Y_quarter(5))
    fmt.Println(Y_quarter(8))
    fmt.Println(Y_quarter(11))
    fmt.Println(Y_quarter(13))
    fmt.Println(Y_quarter(-5))
}
```
And its output :

```go
$ go run . 
1
2
3
4
-1
-1
```