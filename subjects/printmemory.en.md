## printmemory

##**WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a function that takes `(arr [10]int)`, and displays the memory as in the example.

### Expected function

```go
func PrintMemory(arr [10]int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	piscine ".."
)

func main() {
	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	piscine.Memory(arr)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test | cat -e
6800 0000 6500 0000 6c00 0000 6c00 0000 $
6f00 0000 1000 0000 1500 0000 2a00 0000 $
0000 0000 0000 0000 $
hello..*..$
student@ubuntu:~/piscine-go/test$
```
