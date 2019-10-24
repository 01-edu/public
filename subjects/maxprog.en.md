## maxprog

##**WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a function, `Max`, that returns the maximum value in a slice of integers.

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
)

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
