## iterativepower

### Instructions

Write an **iterative** function that returns the power of the `int` passed as parameter.

Negative powers will return `0`. Overflows do **not** have to be dealt with.

### Expected function

```go
func IterativePower(nb int, power int) int {

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
	fmt.Println(piscine.IterativePower(4, 3))
}
```

And its output :

```console
$ go run .
64
$
```
