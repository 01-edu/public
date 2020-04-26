## maxprog

### Instructions

Write a function, `Max`, that returns the maximum value in a slice of integers.

### Expected function

```go
func Max(a []int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	arrInt := []int{23, 123, 1, 11, 55, 93}
	max := Max(arrInt)
	fmt.Println(max
}
```

And its output :

```console
student@ubuntu:~/test$ go build
student@ubuntu:~/test$ ./test
123
student@ubuntu:~/test$
```
