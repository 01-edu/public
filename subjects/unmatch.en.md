## join

### Instructions

Write a function, Unmatch, that returns the element of the slice (arr) that does not have a correspondent pair.

The function must have the next signature.

### Expected function

```go
func Unmatch(arr []int) int {

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
	arr := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := student.Unmatch(arr)
	fmt.Println(unmatch)
}
```

And its output :

```console
student@ubuntu:~/student/unmatch$ go build
student@ubuntu:~/student/unmatch$ ./unmatch
4
student@ubuntu:~/student/unmatch$
```
