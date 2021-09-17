## iterativepower

### Instructions

Write an **iterative** function that returns the value of `nb` to the power of `power`.

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
