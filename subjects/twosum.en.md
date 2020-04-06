## twosum

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Given an array of integers, return indexes of the two numbers such that they add up to a specific target.

If there are more than one solution, return the first one.

If there are no solutions, return nil.

Expected function :

```go
func TwoSum(nums []int, target int) []int {
}
```

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
)

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[0 3]
student@ubuntu:~/[[ROOT]]/test$
```
