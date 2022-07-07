## Square-root

### Instructions
Write a function that takes a number and returns the square root of that number.
- The square root of a number is the number divided by two until the number is less than or equal to one.
- If the number is less than zero return `-1`.

### Expected function

```go
func SquareRoot(number int) int {
    // Your code here
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
    fmt.Println(SquareRoot(9))
    fmt.Println(SquareRoot(16))
    fmt.Println(SquareRoot(25))
    fmt.Println(SquareRoot(26))
    fmt.Println(SquareRoot(0))
    fmt.Println(SquareRoot(-1))
    fmt.Println(SquareRoot(1))
}
```

and the output should be:

```console
$ go run . 
3
4
5
5
0
-1
1
```
