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
    "piscine"
)

func main() {
    fmt.Println(piscine.BinaryCheck(5))
    fmt.Println(piscine.BinaryCheck(0))
    fmt.Println(piscine.BinaryCheck(8))
    fmt.Println(piscine.BinaryCheck(-9))
    fmt.Println(piscine.BinaryCheck(-4))
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
