## reversebits

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a function that takes a `byte`, reverses it `bit` by `bit` (as shown in the example) and returns the result.

### Expected function

```go
func ReverseBits(by byte) byte {

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
	a := byte(0x26)
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", ReverseBits(a))
}
```

and its output :

```console
student@ubuntu:~/piscine-go/reversebits$ go build
student@ubuntu:~/piscine-go/reversebits$ ./reversebits
00100110
01100100
student@ubuntu:~/piscine-go/reversebits$
```
