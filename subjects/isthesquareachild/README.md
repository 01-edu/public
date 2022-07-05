## is-square

### Instructions

Write a function `isSquare` that takes two numbers and returns `true` if the second number is the square of the first and `false` otherwise.
- The function should return `false` if any of the parameters are negative
- Check only if the second parameter is the square of the first one.

### Expected Function
```go
func isSquare(number int, square int) bool {
    // your code here
}
```

### Usage

Here is a possible program to test your function:

``` go
package main

import "fmt"

func main(){
    fmt.Println(isSquare(4, 16))
    fmt.Println(isSquare(5, 23))
    fmt.Println(isSquare(5, 25))
    fmt.Println(isSquare(2, 27))
    fmt.Println(isSquare(6, 36))
    fmt.Println(isSquare(-10, 100))
    fmt.Println(isSquare(100,10))
    fmt.Println(isSquare(8, -64))
}
```

and its output:

```console
$ go run . 
true
false
true
false
true
false
false
false
```

