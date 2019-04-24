## max

### Instructions

Write a function, Max, that returns the maximum value in a slice of integers

The function must have the next signature.

### Expected function

```go
func Max(arr []int) int {

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
	arrInt := []int{23, 123, 1, 11, 55, 93}
	max := student.Max(arrInt)
	fmt.Println(max
}
```

And its output :

```console
student@ubuntu:~/student/max$ go build
student@ubuntu:~/student/max$ ./max
123
student@ubuntu:~/student/max$
```
