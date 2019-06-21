## collatzcountdown

### Instructions

Write a function, `CollatzCountdown`, that returns the number of steps necessary to reach 1 using the collatz countdown.

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
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
10
student@ubuntu:~/piscine/test$
```
