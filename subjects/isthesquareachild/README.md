## is-the-square-a-child

### Instructions

Write a function `IsTheSquareAChild` that takes two numbers and returns `true` if the second number is the square of the first and `false` otherwise.
- The function should return `false` if the number is negative.
- Check only if the first parameter is the square of the second one.

### Expected Function
```go
func IsTheSquareAChild(number int, square int) bool {
    // your code here
}
```

### Usage

Here is a possible program to test your function:

``` go
package main

import "fmt"

func main(){
    fmt.Println(IsTheSquareAChild(4, 16))
    fmt.Println(IsTheSquareAChild(5, 23))
    fmt.Println(IsTheSquareAChild(5, 25))
    fmt.Println(IsTheSquareAChild(2, 27))
    fmt.Println(IsTheSquareAChild(6, 36))
    fmt.Println(IsTheSquareAChild(-10, 100))
    fmt.Println(IsTheSquareAChild(100,10))
    fmt.Println(IsTheSquareAChild(8, -64))
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

