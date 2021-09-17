## collatzcountdown

### Instructions

Write a function, `CollatzCountdown`, that returns the number of steps necessary to reach 1 using the [collatz countdown](https://en.wikipedia.org/wiki/Collatz_conjecture).

- It must return `-1` if `start` is equal to `0` or negative.

### Expected function

```go
func CollatzCountdown(start int) int {

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
	steps := piscine.CollatzCountdown(12)
	fmt.Println(steps)
}
```

And its output :

```console
$ go run .
9
$
```
