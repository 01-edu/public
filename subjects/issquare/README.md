## is-square

### Instructions

Write a function `IsSquare` that takes two numbers and returns `true` if the second number is the square of the first and `false` otherwise.
- The function should return `false` if any of the parameters are negative
- Check only if the second parameter is the square of the first one.

### Expected Function
```go
func IsSquare(number int, square int) bool {
    // your code here
}
```

### Usage

Here is a possible program to test your function:

``` go
package main

import "fmt"

func main(){
    fmt.Println(IsSquare(4, 16))
    fmt.Println(IsSquare(5, 23))
    fmt.Println(IsSquare(5, 25))
    fmt.Println(IsSquare(2, 27))
    fmt.Println(IsSquare(6, 36))
    fmt.Println(IsSquare(-10, 100))
    fmt.Println(IsSquare(100,10))
    fmt.Println(IsSquare(8, -64))
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

