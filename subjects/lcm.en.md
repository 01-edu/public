## lcm

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a function, `lcm`, that returns least common multiple.

All arguments tested will be positive `int` values.

### Expected function

```go
func Lcm(first, second int) int {

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
	fmt.Println(piscine.Lcm(2, 7))
}
```

### Output

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
14
student@ubuntu:~/[[ROOT]]/test$
```
