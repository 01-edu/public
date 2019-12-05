## nauuo

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

There was a vote. There are people who voted positively, negatively, and randomly. 
Figure out if the final answer depends on random people or not.
If it does, return '?', otherwise the result must be either '+', '-', or '0'
Previous characters stand for outcome of the vote: positive/negative/draw.
Input is always positive.

Write a function, `Nauuo`, that returns final result of voting.

### Expected function

```go
func Nauuo(plus, minus, rand int) string {

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
	fmt.Println(piscine.Nauuo(50, 43, 20))
	fmt.Println(piscine.Nauuo(13, 13, 0))
	fmt.Println(piscine.Nauuo(10, 9, 0))
	fmt.Println(piscine.Nauuo(5, 9, 2))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
?
0
+
-
student@ubuntu:~/[[ROOT]]/test$
```
