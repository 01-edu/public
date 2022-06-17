## Square-root

### Instructions
Write a function that takes a number and returns the square root of that number.
- If the number is less than one or the number does not have an integer square root, return `-1`.
- The square root of a number is the number divided by two until the number is less than or equal to one.

### Expected function
```go
func SquareRoot(number int) int {
    // Your code here
}
```

### Usage

```go
package main

import "fmt"

func main() {
    fmt.Println(SquareRoot(9))
    fmt.Println(SquareRoot(16))
    fmt.Println(SquareRoot(25))
    fmt.Println(SquareRoot(26))
    fmt.Println(SquareRoot(0))
}
```

and the output should be:

```console
$ go run . 
3
4
5
-1
-1
```