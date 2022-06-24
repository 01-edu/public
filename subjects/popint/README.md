## popint

### Instructions

Write a function that receives a slice of int and returns a new slice without the last element.

### Expected function

```go
func PopInt(ints []int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
    "piscine"
    "fmt"
)

func main() {
    ints := {6,7,8,9,}
    l := piscine.PopInt(ints)
    fmt.Println(l)
}
```

And its output:

```console
$ go run .
[6 7 8]
```
