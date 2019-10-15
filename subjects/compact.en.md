## Compact

### Instructions

Write a function `Compact` that takes a pointer to a slice of strings as the argument.
This function must:

-   Return the number of elements with non-`nil`.

-   Compact, i.e., delete the elements with `nil` in the slice.

### Expected functions

```go
func Compact(ptr *[]string) int {

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

const N = 6

func main() {
	arr := make([]string, N)
	arr[0] = "a"
	arr[2] = "b"
	arr[4] = "c"

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", piscine.Compact(&arr))

	for _, v := range arr {
		fmt.Println(v)
	}
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
a

b

c

Size after compacting: 3
a
b
c
student@ubuntu:~/piscine-go/test$
```
