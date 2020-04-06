## reachable_number

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(`func main()`) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Let us define a function f(x) by the following: first we add 1 to x, and then while the last digit of the number equals 0, we shall be deleting 0. Let us call 'y' reachable if we can apply **f** to **x** (zero or more times), and get **y**. 102 is reachable from 10098: f(f(f(10098))) = f(f(10099)) = f(101) = f(102). Any number is reachable from itself. You are given a positive number **n**, count how many integers are reachable from **n**.

### Expected function

```go
func Reachablenumber(n int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Reachablenumber(1))
	fmt.Println(Reachablenumber(10))
	fmt.Println(Reachablenumber(1001))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
9
19
36
student@ubuntu:~/[[ROOT]]/test$
```
