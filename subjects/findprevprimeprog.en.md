## findprevprime

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(\`\`\`func main()\`\`\`) even if empty.
- The function main declared needs to **also pass** the \`Restrictions Checker\`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a \`program\`.

### Instructions

Write a function that returns the first prime number that is equal or inferior to the `int` passed as parameter.

If there are no primes inferior to the `int` passed as parameter the function should return 0.

### Expected function

```go
func FindPrevPrime(nb int) int {

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
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
5
3
student@ubuntu:~/piscine-go/test$
```
