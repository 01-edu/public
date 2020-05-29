## printnbrinorder

### Instructions

Write a function which prints the digits of an `int` passed in parameter in ascending order.
All possible values of type `int` have to go through, besides negative numbers.
Conversion to `int64` is not allowed.

### Expected function

```go
func PrintNbrInOrder(n int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import piscine ".."

func main() {
	piscine.PrintNbrInOrder(321)
	piscine.PrintNbrInOrder(0)
	piscine.PrintNbrInOrder(321)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test | cat -e
1230123$
student@ubuntu:~/[[ROOT]]/test$
```
