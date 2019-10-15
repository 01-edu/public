## unmatch

### Instructions

Write a function, `Unmatch`, that returns the element of the slice (arr) that does not have a correspondent pair.

-   If all the number have a correspondent pair, it shoud return `-1`.

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
	piscine ".."
)

func main() {
	arr := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(arr)
	fmt.Println(unmatch)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
4
student@ubuntu:~/piscine-go/test$
```
