## priorprime

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

You are given an integer.
Your function must return sum of all prime numbers prior to the number exclusively. The number is not included.
### Expected function and struct

```go
package piscine

func Priorprime(x int) int {

}
```

### Usage
Here is a possible program to test your function:

```go
package main

import (
	piscine ".."
	"fmt"
)

func main() {
	fmt.Println(piscine.Priorprime(14))
}

```

Its output:

```console
$> go build
$> ./main
41
```
