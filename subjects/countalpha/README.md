## countalpha

### Instructions

Write a function `CountAlpha()` that takes a string as an argument and returns the number of alphabetic characters in the string.

### Expected functions

```go
func CountAlpha(s string) int {

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
	fmt.Println(piscine.CountAlpha("Hello world"))
	fmt.Println(piscine.CountAlpha("H e l l o"))
	fmt.Println(piscine.CountAlpha("H1e2l3l4o"))
}

```

And its output:

```console
$ go run .
10
5
5
```
