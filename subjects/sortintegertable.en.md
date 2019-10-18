## sortintegertable

### Instructions

-   Write a function that reorder an array of `int` in ascending order.

### Expected function

```go
func SortIntegerTable(table []int) {

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
	s := []int{5,4,3,2,1,0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[0,1,2,3,4,5]
student@ubuntu:~/piscine-go/test$
```
