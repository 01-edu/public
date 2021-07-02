## sqrt

### Instructions

Write a function that returns the square root of the `int` passed as parameter, if that square root is a whole number. Otherwise it returns `0`.

### Expected function

```go
func Sqrt(nb int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Sqrt(4))
	fmt.Println(piscine.Sqrt(3))
}
```

And its output :

```console
$ go run .
2
0
$
```
