## sortintegertable

### Instructions

- Write a function that reorder a slice of `int` in ascending order.

### Expected function

```go
func SortIntegerTable(table []int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	s := []int{5,4,3,2,1,0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[0 1 2 3 4 5]
student@ubuntu:~/[[ROOT]]/test$
```
