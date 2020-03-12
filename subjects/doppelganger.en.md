## doppelganger

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

You are given 2 strings. Find out if first string contains second string. If it does, return index of the first string where second string occurs last time. If it does not contain, return "-1"

### Expected function

```go
package main

func DoppelGanger(big, little string) int {

}
```

```go
package main

import (
	"fmt"
)

func main() {
	var result int

	result = DoppelGanger("aaaaaaa", "a")
	fmt.Println(result) // 6

	result = DoppelGanger("qwerty", "t")
	fmt.Println(result) // 4

	result = DoppelGanger("a", "b")
	fmt.Println(result) // -1
}
```
