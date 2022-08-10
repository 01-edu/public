## countalpha

### Instructions

Write a function `CounAlpha` that takes a string in parameter and returns the number of alphabetic characters in the string.

### Expected functions

```go
func CountAlpha(s string) int {
}
```

### Usage

Here is a program to test your function.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(CountAlpha("Hello world"))
    fmt.Println(CountAlpha("H e l l o"))
    fmt.Println(CountAlpha("H1e2l3l4o"))
}
```

And its output :

```console
$ go run .
10
5
5
```