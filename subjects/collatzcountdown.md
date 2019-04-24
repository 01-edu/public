## collatzcountdown

### Instructions

Write a function, CollatzCountdown, that returns the number of steps to reach 1 using the collatz countdown.

The function must have the following signature.

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
  student ".."
)

func main() {
	steps := student.CollatzCountdown(12)
	fmt.Println(steps)
}
```

And its output :

```console
student@ubuntu:~/student/collatzcountdown$ go build
student@ubuntu:~/student/collatzcountdown$ ./collatzcountdown
10
student@ubuntu:~/student/collatzcountdown$
```
