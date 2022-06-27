## ascii

### Instructions

Write a function that receives a string and returns a slice with the ASCII values of its characters. If the string is empty it should return an empty slice.

### Expected function

```go
func Ascii(str string) []byte {

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
    l := piscine.Ascii("Hello")
    fmt.Println(l)
}
```

And its output:

```console
$ go run .
[104 101 108 108 111]
```
