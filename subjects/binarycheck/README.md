## binarycheck

### Instructions

Write a function that takes an int as an argument and returns 0 if the number is odd, and 1 if the number is even.

### Expected function

```go
func BinaryCheck(nbr int) int {

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
    fmt.Println(BinaryCheck(5))
    fmt.Println(BinaryCheck(0))
    fmt.Println(BinaryCheck(8))
    fmt.Println(BinaryCheck(-9))
    fmt.Println(BinaryCheck(-4))
}
```

And its output:

```console
$ go run .
0
1
1
0
1
```
