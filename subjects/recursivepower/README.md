## recursivepower

### Instructions

Write a **recursive** function that returns the the value of `nb` to the power `power`.

Negative powers will return `0`. Overflows do **not** have to be dealt with.

`for` is **forbidden** for this exercise.

### Expected function

```go
func RecursivePower(nb int, power int) int {

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
	fmt.Println(piscine.RecursivePower(4, 3))
}
```

And its output :

```console
$ go run .
64
$
```
