## printnbrinorder

### Instructions

Write a function which prints the digits of an `int` passed in parameter in ascending order.
All possible values of type `int` have to go through, besides negative numbers.
Convertion to `int64` is not allowed.

### Expected function

```go
func PrintNbrInOrder(n int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

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
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1230123
student@ubuntu:~/piscine-go/test$
```
