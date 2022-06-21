## common-multiples

### Instructions

Write a function to check whether a given non-negative number is a multiple of 3 or 7.
- If the number is a multiple of 3 or 7, return `true`.
- If the number is not a multiple of 3 or 7, return `false`.
- If the number is less or equal to 0, return `false`.

### Expected function

```go
func IsMultiple(number int) bool {
    // Your code here
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
    fmt.Println(IsMultiple(3))
    fmt.Println(IsMultiple(7))
    fmt.Println(IsMultiple(8))
    fmt.Println(IsMultiple(9))
    fmt.Println(IsMultiple(-1))
}
```

and the output should be:

```console
$ go run .
true
true
false
true
false
```
