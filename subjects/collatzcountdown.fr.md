## collatzcountdown

### Instructions

Write a function, CollatzCountdown, that returns the number of steps to reach 1 using the collatz countdown.

### Expected function

```go
func CollatzCountdown(start int) int {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	steps := piscine.CollatzCountdown(12)
	fmt.Println(steps)
}
```

And its output :

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./test
10
student@ubuntu:~/student/test$
```
